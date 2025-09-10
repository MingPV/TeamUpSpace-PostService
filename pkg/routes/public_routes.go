package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Order
	orderHandler "github.com/MingPV/PostService/internal/order/handler/rest"
	orderRepository "github.com/MingPV/PostService/internal/order/repository"
	orderUseCase "github.com/MingPV/PostService/internal/order/usecase"

	// Post
	postHandler "github.com/MingPV/PostService/internal/post/handler/rest"
	postRepository "github.com/MingPV/PostService/internal/post/repository"
	postUseCase "github.com/MingPV/PostService/internal/post/usecase"

	// PostReport
	postReportHandler "github.com/MingPV/PostService/internal/postreport/handler/rest"
	postReportRepository "github.com/MingPV/PostService/internal/postreport/repository"
	postReportUseCase "github.com/MingPV/PostService/internal/postreport/usecase"
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {

	api := app.Group("/api/v1")

	// === Dependency Wiring ===

	// Order
	orderRepo := orderRepository.NewGormOrderRepository(db)
	orderService := orderUseCase.NewOrderService(orderRepo)
	orderHandler := orderHandler.NewHttpOrderHandler(orderService)

	// Post
	postRepo := postRepository.NewGormPostRepository(db)
	postService := postUseCase.NewPostService(postRepo)
	postHandler := postHandler.NewHttpPostHandler(postService)

	// PostReport
	postReportRepo := postReportRepository.NewGormPostReportRepository(db)
	postReportService := postReportUseCase.NewPostReportService(postReportRepo)
	postReportHandler := postReportHandler.NewHttpPostReportHandler(postReportService)

	// === Public Routes ===

	// Order routes
	orderGroup := api.Group("/orders")
	orderGroup.Get("/", orderHandler.FindAllOrders)
	orderGroup.Get("/:id", orderHandler.FindOrderByID)
	orderGroup.Post("/", orderHandler.CreateOrder)
	orderGroup.Patch("/:id", orderHandler.PatchOrder)
	orderGroup.Delete("/:id", orderHandler.DeleteOrder)

	// Post routes
	postGroup := api.Group("/posts")
	postGroup.Get("/", postHandler.FindAllPosts)
	postGroup.Get("/:id", postHandler.FindPostByID)
	postGroup.Post("/", postHandler.CreatePost)
	postGroup.Patch("/:id", postHandler.PatchPost)
	postGroup.Delete("/:id", postHandler.DeletePost)

	// PostReport routes
	postReportGroup := api.Group("/post_reports")
	postReportGroup.Get("/", postReportHandler.FindAllPostReports)
	postReportGroup.Get("/:id", postReportHandler.FindPostReportByID)
	postReportGroup.Post("/", postReportHandler.CreatePostReport)
	postReportGroup.Patch("/:id", postReportHandler.PatchPostReport)
	postReportGroup.Delete("/:id", postReportHandler.DeletePostReport)
}
