package graph

import (
	"github.com/altsaqif/go-graphql/cmd/shared/service"
	"github.com/altsaqif/go-graphql/cmd/utils"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB         *gorm.DB
	JwtService *service.JwtService
	Blacklist  *utils.Blacklist
}
