package routes

import (
	userHandler "github.com/MingPV/PostService/internal/user/handler/rest"
	userRepository "github.com/MingPV/PostService/internal/user/repository"
	userUseCase "github.com/MingPV/PostService/internal/user/usecase"
	middleware "github.com/MingPV/PostService/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app fiber.Router, db *gorm.DB) {

	route := app.Group("/api/v1", middleware.JWTMiddleware())

	userRepo := userRepository.NewGormUserRepository(db)
	PostService := userUseCase.NewPostService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(PostService)

	route.Get("/me", userHandler.GetUser)

}
