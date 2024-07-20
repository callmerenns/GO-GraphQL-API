package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/altsaqif/go-graphql/cmd/entity"
	"github.com/altsaqif/go-graphql/cmd/shared/common"
	"github.com/altsaqif/go-graphql/cmd/shared/service"
	"github.com/altsaqif/go-graphql/cmd/utils"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	jwtService *service.JwtService
	db         *gorm.DB
	Blacklist  utils.Blacklist
}

func NewAuthMiddleware(jwtService *service.JwtService, db *gorm.DB, blacklist utils.Blacklist) *AuthMiddleware {
	return &AuthMiddleware{jwtService: jwtService, db: db, Blacklist: blacklist}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (m *AuthMiddleware) AuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver, roles []string) (interface{}, error) {
	c := ctx.Value("ginContext").(*gin.Context)
	reqCtx := graphql.GetOperationContext(ctx)
	headers := reqCtx.Headers

	token := headers.Get("Authorization")
	if token == "" {
		common.SendErrorResponse(c, http.StatusUnauthorized, "No token provided")
		return nil, fmt.Errorf("no token provided")
	}

	token = strings.TrimPrefix(token, "Bearer ")

	// Check if token is blacklisted
	if m.Blacklist.IsBlacklisted(token) {
		common.SendErrorResponse(c, http.StatusUnauthorized, "Token is blacklisted")
		return nil, fmt.Errorf("token is blacklisted")
	}

	claims, err := m.jwtService.ValidateToken(token)
	if err != nil {
		common.SendErrorResponse(c, http.StatusUnauthorized, "Permission denied!")
		return nil, fmt.Errorf("permission denied")
	}

	if !contains(roles, claims.Role) {
		common.SendErrorResponse(c, http.StatusForbidden, "Only admin, reseller & customer role is authorized to access this resource")
		return nil, fmt.Errorf("only admin, reseller & customer role is authorized to access this resource")
	}

	userID := claims.ID
	var user entity.User
	if err := m.db.Where("id = ?", userID).First(&user).Error; err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return nil, fmt.Errorf("user not found")
	}

	log.Printf("User fetched: %+v\n", user)

	ctx = context.WithValue(ctx, "user", &user)
	return next(ctx)
}

// HandleGraphQLErrors handles GraphQL errors
func HandleGraphQLErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				if gqlErr, ok := e.Err.(*gqlerror.Error); ok {
					httpCode, ok := gqlErr.Extensions["httpCode"].(int)
					if !ok {
						httpCode = http.StatusInternalServerError
					}
					c.JSON(httpCode, gin.H{"error": gqlErr.Message})
					return
				}
			}
		}
	}
}

func GetUserFromContext(ctx context.Context) (*entity.User, error) {
	c := ctx.Value("ginContext").(*gin.Context)
	user, ok := ctx.Value("user").(*entity.User)
	if !ok || user == nil {
		common.SendErrorResponse(c, http.StatusNotFound, "User not found in context")
		return nil, fmt.Errorf("user not found in context")
	}
	return user, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
