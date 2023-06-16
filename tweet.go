package go_functions

import (
	"io"
	"net/http"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World")
}
