package handlers

import (
	"net/http"
)

import (
	g "grouter/src"
	v "github.com/gima/govalid/v1"
	structs "github.com/fatih/structs"
	"github.com/stretchr/testify/http"
)

func SecurityDefinitionHandler(securityDefinition *map[interface{}]interface{}) handler {
	return nil
}


//func(http.ResponseWriter, *http.Request) {
//	var isPresent = func (name string, context interface{}) bool {
//		var _, ok = context["foo"];
//
//		return ok
//	}
//	var handler = func (w http.ResponseWriter, r *http.Request) {
//		for _, def := range *securityDefinition {
//			var isGoodReq = false
//
//			switch def["in"] {
//			case "header":
//				isGoodReq = isPresent(def["name"], structs.Map(r.Header))
//			case "query":
//				isGoodReq = isPresent(def["name"], structs.Map(r.RequestURI))
//			}
//
//			if (!isGoodReq) {
//				// TODO : 400
//				return;
//			}
//
//			// TODO : 200
//			return;
//		}
//	}
//
//	return handler
//}