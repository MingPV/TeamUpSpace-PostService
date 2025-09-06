package rest

import (
	"strconv"

	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/post/dto"
	"github.com/MingPV/PostService/internal/post/usecase"
	"github.com/MingPV/PostService/pkg/apperror"
	responses "github.com/MingPV/PostService/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpPostHandler struct {
	postUseCase usecase.PostUseCase
}

func NewHttpPostHandler(useCase usecase.PostUseCase) *HttpPostHandler {
	return &HttpPostHandler{postUseCase: useCase}
}

// CreatePost godoc
// @Summary Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body entities.Post true "Post payload"
// @Success 201 {object} entities.Post
// @Router /post [post]
func (h *HttpPostHandler) CreatePost(c *fiber.Ctx) error {
	var req dto.CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	post := &entities.Post{
		PostBy:		req.PostBy,
		Title:		req.Title,
		Detail:		req.Detail,
		ImageURL:	req.ImageURL,
		EventID:	req.EventID,
		Status:		req.Status,
	}

	//check validate
	if msg, err := validateCreateOrPatchPost(post); err != nil {
		return responses.ErrorWithMessage(c, err, msg);
	}

	if err := h.postUseCase.CreatePost(post); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToPostResponse(post))
}

func validateCreateOrPatchPost(post *entities.Post) (string, error) {
	if(post.Title == ""){
		return "Title is required", apperror.ErrRequiredField
	}

	return "", nil
}

// FindAllPosts godoc
// @Summary Get all posts
// @Tags posts
// @Produce json
// @Success 200 {array} entities.Order
// @Router /posts [get]
func (h *HttpPostHandler) FindAllPosts(c *fiber.Ctx) error {
	posts, err := h.postUseCase.FindAllPosts()
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToPostResponseList(posts))
}

// FindPostByID godoc
// @Summary Get post by ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} entities.Post
// @Router /posts/{id} [get]
func (h *HttpPostHandler) FindPostByID(c *fiber.Ctx) error {
	id := c.Params("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	post, err := h.postUseCase.FindPostByID(postID)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToPostResponse(post))
}

// DeletePost godoc
// @Summary Delete an post by ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} response.MessageResponse
// @Router /posts/{id} [delete]
func (h *HttpPostHandler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	if err := h.postUseCase.DeletePost(postID); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "post deleted")
}

// PatchPost godoc
// @Summary Update an post partially
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body entities.Post true "Post update payload"
// @Success 200 {object} entities.Post
// @Router /posts/{id} [patch]
func (h *HttpPostHandler) PatchPost(c *fiber.Ctx) error {
	id := c.Params("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	var req dto.CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	post := &entities.Post{
		PostBy:		req.PostBy,
		Title:		req.Title,
		Detail:		req.Detail,
		ImageURL:	req.ImageURL,
		EventID:	req.EventID,
		Status:		req.Status,
	}

	msg, err := validateCreateOrPatchPost(post)
	if err != nil {
		return responses.ErrorWithMessage(c, err, msg)
	}

	updatedPost, err := h.postUseCase.PatchPost(postID, post)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToPostResponse(updatedPost))
}