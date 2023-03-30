package api

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	"project/internal/common"
	"project/internal/common/errorx"
	"project/internal/common/messages"
	"project/internal/services/task"
)

func (s *createTaskRequest) Validate() *errorx.Result {
	out := new(errorx.Result)
	if !common.StringInArray(
		s.Method,
		[]string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodHead,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodPut,
			http.MethodTrace,
		}) {
		out.AddFieldError(messages.FieldMethod, errorx.ErrorDetailsFn(errorx.ErrInvalidMethod))
	}

	if s.URL == "" {
		out.AddFieldError(messages.FieldURL, errorx.ErrorDetailsFn(errorx.ErrInvalidURL))
	} else {
		_, err := url.ParseRequestURI(s.URL)
		if err != nil {
			out.AddFieldError(messages.FieldURL, errorx.ErrorDetailsFn(errorx.ErrInvalidURL))
		}
	}

	return out
}

// createTask godoc
// @Tags task
// @Summary Create Task
// @Description Create task.
// @Accept   json
// @Produce  json
// @Param params body createTaskRequest true "method, url, headers, query_params, timeout_in_milliseconds, body"
// @Success 200 {object} createTaskResponse
// @Failure 400 {object} errorx.Result
// @Failure 500 {object} errorx.Result
// @Router /task [post]
func createTask(taskSvc taskService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var req createTaskRequest
		if err = c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, errorx.CodeError(errorx.Code(errorx.ErrInvalidRequestBody), err))
		}

		if v := errorx.Validate(&req); !v.IsValid() {
			return c.JSON(http.StatusBadRequest, v)
		}

		id, err := taskSvc.AddTask(task.CreateTaskInput{
			Method:      req.Method,
			URL:         req.URL,
			Headers:     req.Headers,
			Body:        req.Body,
			QueryParams: req.QueryParams,
		})

		if err != nil {
			return c.JSON(http.StatusBadRequest, errorx.CodeError(errorx.Code(err), err))
		}

		return c.JSON(http.StatusOK, createTaskResponse{ID: id})
	}
}

// GetTask godoc
// @Tags task
// @Summary Get Task
// @Description Get task.
// @Produce  json
// @Param id path string true "task ID"
// @Success 200 {object} task.JobResult
// @Failure 400 {object} errorx.Result
// @Failure 500 {object} errorx.Result
// @Router /task/{id} [get]
func GetTask(taskSvc taskService) echo.HandlerFunc {
	return func(c echo.Context) error {
		jobID := c.Param("id")
		if jobID == "" {
			return c.JSON(http.StatusBadRequest, errorx.CodeError(errorx.Code(errorx.ErrInvalidJobID), errorx.ErrInvalidJobID))
		}
		res, err := taskSvc.GetTask(jobID)
		if err != nil {
			switch {
			case errors.Is(err, errorx.ErrJobNotFound):
				return c.JSON(http.StatusBadRequest, errorx.CodeError(errorx.Code(errorx.ErrJobNotFound), errorx.ErrJobNotFound))
			default:
				return c.JSON(http.StatusInternalServerError, errorx.CodeError(errorx.Code(err), err))
			}
		}
		return c.JSON(http.StatusOK, res)
	}
}
