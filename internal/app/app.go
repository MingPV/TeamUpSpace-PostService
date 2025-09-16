package app

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/MingPV/PostService/internal/entities"
	GrpcOrderHandler "github.com/MingPV/PostService/internal/order/handler/grpc"
	orderRepository "github.com/MingPV/PostService/internal/order/repository"
	orderUseCase "github.com/MingPV/PostService/internal/order/usecase"
	orderpb "github.com/MingPV/PostService/proto/order"

	GrpcPostHandler "github.com/MingPV/PostService/internal/post/handler/grpc"
	postRepository "github.com/MingPV/PostService/internal/post/repository"
	postUseCase "github.com/MingPV/PostService/internal/post/usecase"
	postpb "github.com/MingPV/PostService/proto/post"

	GrpcQuestionHandler "github.com/MingPV/PostService/internal/question/handler/grpc"
	questionRepository "github.com/MingPV/PostService/internal/question/repository"
	questionUseCase "github.com/MingPV/PostService/internal/question/usecase"
	questionpb "github.com/MingPV/PostService/proto/question"

	GrpcAnswerHandler "github.com/MingPV/PostService/internal/answer/handler/grpc"
	answerRepository "github.com/MingPV/PostService/internal/answer/repository"
	answerUseCase "github.com/MingPV/PostService/internal/answer/usecase"
	answerpb "github.com/MingPV/PostService/proto/answer"

	GrpcPostLikeHandler "github.com/MingPV/PostService/internal/postlike/handler/grpc"
	postlikeRepository "github.com/MingPV/PostService/internal/postlike/repository"
	postlikeUseCase "github.com/MingPV/PostService/internal/postlike/usecase"
	postlikepb "github.com/MingPV/PostService/proto/postlike"

	GrpcPostReportHandler "github.com/MingPV/PostService/internal/postreport/handler/grpc"
	postreportRepository "github.com/MingPV/PostService/internal/postreport/repository"
	postreportUseCase "github.com/MingPV/PostService/internal/postreport/usecase"
	postreportpb "github.com/MingPV/PostService/proto/postreport"

	GrpcTeamRequestHandler "github.com/MingPV/PostService/internal/teamrequest/handler/grpc"
	teamrequestRepository "github.com/MingPV/PostService/internal/teamrequest/repository"
	teamrequestUseCase "github.com/MingPV/PostService/internal/teamrequest/usecase"
	teamrequestpb "github.com/MingPV/PostService/proto/teamrequest"

	"github.com/MingPV/PostService/pkg/config"
	"github.com/MingPV/PostService/pkg/database"
	"github.com/MingPV/PostService/pkg/middleware"
	"github.com/MingPV/PostService/pkg/routes"

	"google.golang.org/grpc/reflection"
)

// rest
func SetupRestServer(db *gorm.DB, cfg *config.Config) (*fiber.App, error) {
	app := fiber.New()
	middleware.FiberMiddleware(app)
	// comment out Swagger when testing
	// routes.SwaggerRoute(app)
	routes.RegisterPublicRoutes(app, db)
	routes.RegisterPrivateRoutes(app, db)
	routes.RegisterNotFoundRoute(app)
	return app, nil
}

// grpc
func SetupGrpcServer(db *gorm.DB, cfg *config.Config) (*grpc.Server, error) {
	s := grpc.NewServer()
	reflection.Register(s)
	//order
	orderRepo := orderRepository.NewGormOrderRepository(db)
	orderService := orderUseCase.NewOrderService(orderRepo)

	orderHandler := GrpcOrderHandler.NewGrpcOrderHandler(orderService)
	orderpb.RegisterOrderServiceServer(s, orderHandler)

	//post
	postRepo := postRepository.NewGormPostRepository(db)
	postService := postUseCase.NewPostService(postRepo)

	postHandler := GrpcPostHandler.NewGrpcPostHandler(postService)
	postpb.RegisterPostServiceServer(s, postHandler)

	//question
	questionRepo := questionRepository.NewGormQuestionRepository(db)
	questionService := questionUseCase.NewQuestionService(questionRepo)

	questionHandler := GrpcQuestionHandler.NewGrpcQuestionHandler(questionService)
	questionpb.RegisterQuestionServiceServer(s, questionHandler)

	//answer
	answerRepo := answerRepository.NewGormAnswerRepository(db)
	answerService := answerUseCase.NewAnswerService(answerRepo)

	answerHandler := GrpcAnswerHandler.NewGrpcAnswerHandler(answerService)
	answerpb.RegisterAnswerServiceServer(s, answerHandler)

	//postlike
	postlikeRepo := postlikeRepository.NewGormPostLikeRepository(db)
	postlikeService := postlikeUseCase.NewPostLikeService(postlikeRepo)

	postlikeHandler := GrpcPostLikeHandler.NewGprcPostLikeHandler(postlikeService)
	postlikepb.RegisterPostLikeServiceServer(s, postlikeHandler)

	//postreport
	postreportRepo := postreportRepository.NewGormPostReportRepository(db)
	postreportService := postreportUseCase.NewPostReportService(postreportRepo)

	postreportHandler := GrpcPostReportHandler.NewGrpcPostReportHandler(postreportService)
	postreportpb.RegisterPostReportServiceServer(s, postreportHandler)

	//teamrequest
	teamrequestRepo := teamrequestRepository.NewGormTeamRequestRepository(db)
	teamrequestService := teamrequestUseCase.NewTeamRequestService(teamrequestRepo)

	teamrequestHandler := GrpcTeamRequestHandler.NewGrpcTeamRequestHandler(teamrequestService)
	teamrequestpb.RegisterTeamRequestServiceServer(s, teamrequestHandler)

	return s, nil
}

// dependencies
func SetupDependencies(env string) (*gorm.DB, *config.Config, error) {
	cfg := config.LoadConfig(env)

	db, err := database.Connect(cfg.DatabaseDSN)
	if err != nil {
		return nil, nil, err
	}

	if env == "test" {
		db.Migrator().DropTable(&entities.Order{}, &entities.Post{}, &entities.Answer{}, &entities.Comment{}, &entities.PostLike{}, &entities.PostReport{}, &entities.Question{}, &entities.Subcomment{}, &entities.TeamRequest{})
	}

	// db.Migrator().DropTable(&entities.Order{}, &entities.Post{}, &entities.Answer{}, &entities.Comment{}, &entities.PostLike{}, &entities.PostReport{}, &entities.Question{}, &entities.Subcomment{}, &entities.TeamRequest{})

	if err := db.AutoMigrate(&entities.Order{}, &entities.Post{}, &entities.Answer{}, &entities.Comment{}, &entities.PostLike{}, &entities.PostReport{}, &entities.Question{}, &entities.Subcomment{}, &entities.TeamRequest{}); err != nil {
		return nil, nil, err
	}

	return db, cfg, nil
}
