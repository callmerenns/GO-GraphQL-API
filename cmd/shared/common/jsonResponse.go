// shared/common/jsonResponse.go

package common

import (
	"net/http"

	common "github.com/altsaqif/go-graphql/cmd/shared/model"
	"github.com/altsaqif/go-graphql/graph/model"
	"github.com/gin-gonic/gin"
)

// SendCreateResponse defines the standard create response structure
func SendCreateRegisterResponse(ctx *gin.Context, message string, data *model.RegisterResponse) {
	ctx.JSON(http.StatusCreated, &model.SingleRegisterResponse{
		Status: &model.Status{

			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}

// SendCreateResponse defines the standard create response structure
func SendCreateProductResponse(ctx *gin.Context, message string, data *model.ProductResponse) {
	ctx.JSON(http.StatusCreated, &model.SingleProductResponse{
		Status: &model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}

// SendSuccessResponse defines the standard success response structure
func SendSuccessLogoutResponse(ctx *gin.Context, data *model.LogoutResponse) {
	ctx.JSON(http.StatusOK, &model.SingleLogoutResponse{
		Status: &model.Status{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data: data,
	})
}

// SendSuccessResponse defines the standard success response structure
func SendSuccessProductResponse(ctx *gin.Context, data *model.ProductResponse) {
	ctx.JSON(http.StatusOK, &model.SingleProductResponse{
		Status: &model.Status{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data: data,
	})
}

// SendSuccessResponse defines the standard success response structure
func SendSuccessProductByStockResponse(ctx *gin.Context, data []*model.ProductResponse) {
	ctx.JSON(http.StatusOK, &model.AnyProductResponse{
		Status: &model.Status{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data: data,
	})
}

// SendSuccessResponse defines the standard success response structure
func SendSuccessUserResponse(ctx *gin.Context, data *model.UserResponse) {
	ctx.JSON(http.StatusOK, &model.SingleUserResponse{
		Status: &model.Status{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data: data,
	})
}

// SendSuccessResponse defines the standard success response structure
func SendSuccessDeleteProductResponse(ctx *gin.Context, data *model.Status) {
	ctx.JSON(http.StatusOK, &model.Status{
		Code:    data.Code,
		Message: data.Message,
	})
}

// SendSingleLoginResponse defines the standard single response structure
func SendSingleLoginResponse(ctx *gin.Context, message string, data *model.LoginResponse) {
	ctx.JSON(http.StatusOK, &model.SingleLoginResponse{
		Status: &model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

// SendPagedResponse defines the standard paged response structure
func SendPagedUserResponse(ctx *gin.Context, total int, limit int, offset int, users []*model.UserResponse) {
	ctx.JSON(http.StatusOK, &model.UserListResponse{
		Total:  total,
		Limit:  limit,
		Offset: offset,
		Users:  users,
	})
}

// SendPagedResponse defines the standard paged response structure
func SendPagedProductResponse(ctx *gin.Context, total int, limit int, offset int, products []*model.ProductResponse) {
	ctx.JSON(http.StatusOK, &model.ProductListResponse{
		Total:    total,
		Limit:    limit,
		Offset:   offset,
		Products: products,
	})
}

// SendErrorResponse400 defines the standard error response structure
func SendErrorResponse400(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, &model.Status{
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

// SendErrorResponse401 defines the standard error response structure
func SendErrorResponse401(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, &model.Status{
		Code:    http.StatusUnauthorized,
		Message: message,
	})
}

// SendErrorResponse404 defines the standard error response structure
func SendErrorResponse404(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, &model.Status{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

// SendErrorResponse500 defines the standard error response structure
func SendErrorResponse500(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, &model.Status{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

// SendErrorResponse defines the standard error response structure
func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &common.Status{
		Code:    code,
		Message: message,
	})
}
