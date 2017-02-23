package src

import (
	_ "net/http"
	"handlers"
)

type Route struct {
	method string `json:"method"`
	pattern string `json:"pattern"`
	//handler handlers.handler
}

type Router struct {


}


