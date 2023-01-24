package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SliceSuccessResponse(w http.ResponseWriter, data []bool) {
	response, err := json.Marshal(data)

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
