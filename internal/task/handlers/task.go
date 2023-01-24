package handlers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	baseErrors "task_api/pkg/errors"
	"task_api/pkg/rest"
)

var (
	ErrServiceTask = func(err error) error {
		return errors.Wrapf(err, "function Task failed in api service")
	}
)

type ApiServiceInterface interface {
	Task([]int64) ([]bool, error)
}

type ApiHandler struct {
	service ApiServiceInterface
}

func NewApiHandler(service ApiServiceInterface) *ApiHandler {
	return &ApiHandler{
		service: service,
	}
}

func (handler *ApiHandler) Task(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		rest.InternalServerErrorResponse(w, baseErrors.ErrReadAll(err))
		return
	}

	var response []int64

	err = json.Unmarshal(body, &response)
	if err != nil {
		rest.BadRequestResponse(w, baseErrors.ErrJSONUnmarshal(err))
		return
	}

	result, err := handler.service.Task(response)
	if err != nil {
		rest.BadRequestResponse(w, ErrServiceTask(err))
		return
	}

	rest.SliceSuccessResponse(w, result)
}
