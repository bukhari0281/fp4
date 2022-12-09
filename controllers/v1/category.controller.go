package v1

import (
	"fmt"
	"fp4/common/obj"
	"fp4/common/response"
	"fp4/dto"
	"fp4/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	// All(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
	FindOneCategoryByID(ctx *gin.Context)
}

type categoryHandler struct {
	categoryService service.CategoryService
	jwtService      service.JWTService
}

func NewCategoryHandler(categoryService service.CategoryService, jwtService service.JWTService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
		jwtService:      jwtService,
	}
}

// CreateCategory implements CategoryHandler
func (c *categoryHandler) CreateCategory(ctx *gin.Context) {
	var createCategoryReq dto.CreateCategoryRequest
	err := ctx.ShouldBind(&createCategoryReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.categoryService.CreateCategory(createCategoryReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)
}

func (c *categoryHandler) FindOneCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.categoryService.FindOneCategoryByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}
