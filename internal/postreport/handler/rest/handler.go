package rest

import (
	"strconv"

	"github.com/MingPV/PostService/internal/postreport/dto"
	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/postreport/usecase"
	"github.com/MingPV/PostService/pkg/apperror"
	responses "github.com/MingPV/PostService/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpPostReportHandler struct {
	postReportUseCase usecase.PostReportUseCase
}

func NewHttpPostReportHandler(uc usecase.PostReportUseCase) *HttpPostReportHandler {
	return &HttpPostReportHandler{postReportUseCase: uc}
}

// CreatePostReport godoc
// @Summary Create a new post report
// @Tags post-reports
// @Accept json
// @Produce json
// @Param post_report body dto.CreatePostReportRequest true "Post report payload"
// @Success 201 {object} dto.PostReportResponse
// @Router /post_reports [post]
func (h *HttpPostReportHandler) CreatePostReport(c *fiber.Ctx) error {
	var req dto.CreatePostReportRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	report := &entities.PostReport{
		PostId:    req.PostId,
		Reporter:  req.Reporter,
		Report_to: req.ReportTo,
		Detail:    req.Detail,
		Status:    req.Status,
	}

	if msg, err := validateCreatePostReport(report); err != nil {
		return responses.ErrorWithMessage(c, err, msg)
	}

	if err := h.postReportUseCase.CreatePostReport(report); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToPostReportResponse(report))
}

func validateCreatePostReport(report *entities.PostReport) (string, error) {
	if report.PostId == 0 {
		return "PostId is required", apperror.ErrRequiredField
	}
	if report.Detail == "" {
		return "Detail is required", apperror.ErrRequiredField
	}
	return "", nil
}

// FindAllPostReports godoc
// @Summary Get all post reports
// @Tags post-reports
// @Produce json
// @Success 200 {array} dto.PostReportResponse
// @Router /post_reports [get]
func (h *HttpPostReportHandler) FindAllPostReports(c *fiber.Ctx) error {
	reports, err := h.postReportUseCase.FindAllPostReports()
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostReportResponseList(reports))
}

// FindPostReportByID godoc
// @Summary Get post report by ID
// @Tags post-reports
// @Produce json
// @Param id path int true "Post report ID"
// @Success 200 {object} dto.PostReportResponse
// @Router /post_reports/{id} [get]
func (h *HttpPostReportHandler) FindPostReportByID(c *fiber.Ctx) error {
	id := c.Params("id")
	reportId, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	report, err := h.postReportUseCase.FindPostReportByID(reportId)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToPostReportResponse(report))
}

// DeletePostReport godoc
// @Summary Delete a post report by ID
// @Tags post-reports
// @Produce json
// @Param id path int true "Post report ID"
// @Success 200 {object} response.MessageResponse
// @Router /post_reports/{id} [delete]
func (h *HttpPostReportHandler) DeletePostReport(c *fiber.Ctx) error {
	id := c.Params("id")
	reportId, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	if err := h.postReportUseCase.DeletePostReport(reportId); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "post report deleted")
}

// PatchPostReport godoc
// @Summary Update a post report partially
// @Tags post-reports
// @Accept json
// @Produce json
// @Param id path int true "Post report ID"
// @Param post_report body dto.CreatePostReportRequest true "Post report update payload"
// @Success 200 {object} dto.PostReportResponse
// @Router /post_reports/{id} [patch]
func (h *HttpPostReportHandler) PatchPostReport(c *fiber.Ctx) error {
	id := c.Params("id")
	reportId, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	var req dto.CreatePostReportRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	report := &entities.PostReport{
		PostId:    req.PostId,
		Reporter:  req.Reporter,
		Report_to: req.ReportTo,
		Detail:    req.Detail,
		Status:    req.Status,
	}

	msg, err := validateCreatePostReport(report)
	if err != nil {
		return responses.ErrorWithMessage(c, err, msg)
	}

	updatedReport, err := h.postReportUseCase.PatchPostReport(reportId, report)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToPostReportResponse(updatedReport))
}