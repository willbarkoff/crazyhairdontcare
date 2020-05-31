package api

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func photosRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// session, err := store.Get(r, "session-name")
	// if err != nil {
	// 	errlog.LogError("getting session", err)
	// 	writeJSON(w, http.StatusInternalServerError, errorResponse{"error", "internal_server_error"})
	// 	return
	// }

	// loggedIn, ok := session.Values["loggedIn"].(bool)

	// if !ok || !loggedIn {
	// 	writeJSON(w, http.StatusUnauthorized, errorResponse{"error", "unauthorized"})
	// 	return
	// }

	photo := p.ByName("name")

	image, openErr := ioutil.ReadFile("./photos/" + photo)
	if openErr != nil {
		fmt.Println(openErr)
		writeJSON(w, http.StatusNotFound, errorResponse{"error", "not_found"})
		return
	}

	w.Write(image)

	mimeType := mime.TypeByExtension("." + strings.Split(photo, ".")[len(strings.Split(photo, "."))-1])
	w.Header().Set("Content-Type", mimeType)
}
