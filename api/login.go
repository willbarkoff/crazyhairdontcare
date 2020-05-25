package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/willbarkoff/crazyhairdontcare/errlog"
)

func loginRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		errlog.LogError("getting session", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{"error", "internal_server_error"})
		return
	}

	if r.FormValue("password") != password {
		writeJSON(w, http.StatusUnauthorized, errorResponse{"error", "invalid_password"})
		return
	}

	session.Values["loggedIn"] = true

	err = session.Save(r, w)
	if err != nil {
		errlog.LogError("saving session", err)
		writeJSON(w, http.StatusUnauthorized, errorResponse{"error", "internal_server_error"})
		return
	}

	writeJSON(w, http.StatusOK, statusResponse{"ok"})
}
