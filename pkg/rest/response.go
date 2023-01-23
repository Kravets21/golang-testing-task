package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	type structResponse struct {
		data interface{}
	}

	response, err := json.Marshal(structResponse{
		data: data,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func InternalServerErrorResponse(w http.ResponseWriter, err error) {
	fmt.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func BadRequestResponse(w http.ResponseWriter, err error) {
	fmt.Println(err)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
