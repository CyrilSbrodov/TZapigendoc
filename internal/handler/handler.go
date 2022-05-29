package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"TZapigendoc/internal"
	"github.com/gorilla/mux"
)

type Handlers interface {
	Register(router *mux.Router)
}

type handler struct {
	Service
}

func NewHandler(service *Service) Handlers {
	return &handler{
		*service,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/", Handler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var result internal.Result
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var rep internal.Replace
	defer r.Body.Close()

	if err := json.Unmarshal(content, &rep); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result.URLWord, result.URLPDF = internal.ReplaceWord(&rep)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result.URLWord))
	w.Write([]byte("\n"))
	w.Write([]byte(result.URLPDF))
	return
}
