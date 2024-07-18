package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/altsaqif/go-graphql/cmd/config"
	"github.com/altsaqif/go-graphql/cmd/delivery/middleware"
	"github.com/altsaqif/go-graphql/cmd/entity"
	"github.com/altsaqif/go-graphql/cmd/shared/service"
	_ "github.com/altsaqif/go-graphql/docs"
	"github.com/altsaqif/go-graphql/graph"
	"github.com/altsaqif/go-graphql/graph/generated"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// @title Go GraphQL API
// @version 1.0
// @description This is a sample server for a Go GraphQL API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	defaultPort := cfg.ApiPort
	if port := os.Getenv("PORT"); port != "" {
		defaultPort = port
	}

	// Initialize Gin
	r := gin.Default()
	r.Use(middleware.HandleGraphQLErrors(), middleware.CORSMiddleware())

	// Database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DbConfig.User, cfg.DbConfig.Password, cfg.DbConfig.Host, cfg.DbConfig.Port, cfg.DbConfig.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Drop existing tables
	db.Migrator().DropTable(&entity.User{}, &entity.Product{}, &entity.Enrollment{})

	err = db.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Enrollment{})
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// JWT service
	tokenConfig := service.TokenConfig{
		IssuerName:       cfg.TokenConfig.IssuerName,
		JwtSignatureKey:  cfg.TokenConfig.JwtSignatureKey,
		JwtSigningMethod: cfg.TokenConfig.JwtSigningMethod,
		JwtExpiresTime:   cfg.TokenConfig.JwtExpiresTime,
	}
	jwtService := service.NewJwtService(tokenConfig)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtService, db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			DB:         db,
			JwtService: jwtService,
		},
		Directives: generated.DirectiveRoot{
			Auth: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
				return authMiddleware.AuthDirective(ctx, obj, next, []string{
					"admin",
					"customer",
					"reseller",
				})
			},
		},
	}))

	// API v1 group
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/query", func(c *gin.Context) {
			ctx := context.WithValue(c.Request.Context(), "ginContext", c)
			srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
		})
		apiV1.GET("/", func(c *gin.Context) {
			playground.Handler("GraphQL", "/api/v1/query").ServeHTTP(c.Writer, c.Request)
		})
		apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Printf("connect to http://localhost:%s/api/v1/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
