package main

import (
	"fp4/config"
	v1 "fp4/controllers/v1"
	middleware "fp4/middlewares"
	"fp4/repo"
	"fp4/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                = config.SetupDatabaseConnection()
	userRepo        repo.UserRepository     = repo.NewUserRepo(db)
	productRepo     repo.ProductRepository  = repo.NewProductRepo(db)
	categoryRepo    repo.CategoryRepository = repo.NewCategoryRepo(db)
	authService     service.AuthService     = service.NewAuthService(userRepo)
	jwtService      service.JWTService      = service.NewJWTService()
	userService     service.UserService     = service.NewUserService(userRepo)
	productService  service.ProductService  = service.NewProductService(productRepo)
	categoryService service.CategoryService = service.NewCategoryService(categoryRepo)
	authHandler     v1.AuthHandler          = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler     v1.UserHandler          = v1.NewUserHandler(userService, jwtService)
	productHandler  v1.ProductHandler       = v1.NewProductHandler(productService, jwtService)
	categoryHandler v1.CategoryHandler      = v1.NewCategoryHandler(categoryService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	productRoutes := server.Group("api/product", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.GET("/", productHandler.All)
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.FindOneProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	categoryRoutes := server.Group("api/category", middleware.AuthorizeJWT(jwtService))
	{
		categoryRoutes.POST("/", categoryHandler.CreateCategory)
		categoryRoutes.GET("/:id", categoryHandler.FindOneCategoryByID)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
