package dto

import (
	"strconv"
	"time"

	"github.com/altsaqif/go-graphql/cmd/entity"
	"github.com/altsaqif/go-graphql/graph/model"
)

func ConvertToUserResponse(user *entity.User) *model.UserResponse {
	resp := &model.UserResponse{
		ID:        strconv.FormatUint(uint64(user.ID), 10),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	// Check if UpdatedAt is not NULL before formatting
	if !user.UpdatedAt.IsZero() {
		resp.UpdatedAt = user.UpdatedAt.Format(time.RFC3339)
	}

	// Check if DeletedAt is not NULL before formatting
	if !user.DeletedAt.Valid {
		resp.DeletedAt = user.DeletedAt.Time.Format(time.RFC3339)
	}

	return resp
}
