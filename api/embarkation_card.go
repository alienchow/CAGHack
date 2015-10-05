package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/alienchow/CAGHack/dto"
	"github.com/alienchow/CAGHack/overlay"

	"github.com/gorilla/mux"
)

const (
	defaultFormat = `.png`
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
	request.Parse()

	fmt.Printf("Received request: %+v\n", request)

	buf := overlay.Process(request)
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	_, err = w.Write(buf.Bytes())
}
