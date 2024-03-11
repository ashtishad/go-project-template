package transport

import (
	"fmt"
	"net/http"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name") // go 1.22 introduced path param from routes.
	if name == "" {
		name = "World"
	}

	resp := fmt.Sprintf("Hello %s", name)

	writeResponse(w, http.StatusOK, map[string]string{"response": resp})
}
