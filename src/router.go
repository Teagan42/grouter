package src

import (
	"errors"
	"net/http"
	"handlers"
	"fmt"
	"github.com/fatih/structs"
)


type Route struct {
	method string `json:"method"`
	pattern string `json:"pattern"`
	handler handlers
	//handler handlers.handler
}

type RouteKey string

//func keyFromRoute(r *Route) string {
//	if r == nil || &r == nil {
//		return nil
//	}
//
//	return fmt.Sprintf("%s-%s", r.method, r.pattern)
//}

//var mux map[RouteKey]Route

func loadSchema(
	schema *interface{},
	handlers *map[RouteKey]func(w http.ResponseWriter, r *http.Request),
	permissionProvider *func(r *http.Request) map[string][]string) ([]Route, error) {

	if schema == nil {
		return []Route{}, errors.New("Schema must be defined")
	}

	if handlers == nil {
		return []Route{}, errors.New("Route handler map must be defined")
	}

	//var routes map[RouteKey]Route = make(map[RouteKey]Route)
	//
	//for k,  := range structs.Map(schema["paths"]) {
	//	registerRoute()
	//}

	return nil, nil
}

func registerRoute(schema *interface{}, handler *func(w http.ResponseWriter, r *http.Request)) ([]Route, error) {
	if schema == nil {
		return nil, errors.New("Route chema must be defined")
	}

	if handler == nil {
		return nil, errors.New("Route handler must be defined")
	}

	return nil, nil
}


