package api

import (
	"embark/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	defaultFormat = `.pdf`
)

func EmbarkationCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	format, ok := vars["format"]
	if !ok {
		format = defaultFormat
	}
	format = strings.TrimLeft(format, ".")

	request := &dto.EmbarkationCardRequest{Format: format}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(request)
	if err != nil {
		fmt.Println("Error parsing request. Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := request.CheckFormat(); err != nil {
		fmt.Printf("JSON request is invalid. Error: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	fmt.Printf("Received request: %+v\n", request)
	w.WriteHeader(http.StatusOK)
}
