package api

import (
	"fmt"
	"net/http"
)

type Api struct {
	Router *http.ServeMux
}

//decorator pattern
func (a *Api) GET(route string, handler func(w http.ResponseWriter, r *http.Request) error) {
	a.Router.HandleFunc(fmt.Sprintf("GET %v", route), func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("Error occured"))
			return
		}
	})
}

func (a *Api) Start(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router)
}

func NewApi() *Api {
	return &Api{Router: http.NewServeMux()}
}
