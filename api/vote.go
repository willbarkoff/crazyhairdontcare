package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/willbarkoff/crazyhairdontcare/errlog"
)

func voteRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		errlog.LogError("getting session", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{"error", "internal_server_error"})
		return
	}

	loggedIn, ok := session.Values["loggedIn"].(bool)
	if !ok {
		errlog.LogError("type assertion not ok when logging in", nil)
		writeJSON(w, http.StatusInternalServerError, errorResponse{"error", "internal_server_error"})
		return
	}

	if !loggedIn {
		writeJSON(w, http.StatusUnauthorized, errorResponse{"error", "unauthorized"})
		return
	}

	contestantNum, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{"error", "invalid_params"})
		return
	}

	var throwaway int

	contestantErr := db.QueryRow("SELECT id FROM contestants WHERE id = ?", contestantNum).Scan(&throwaway)
	if contestantErr != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{"error", "invalid_contestant"})
		return
	}

	_, err = db.Exec("UPDATE contestants SET votes = votes + 1 WHERE id = ?", contestantNum)
	if err != nil {
		errlog.LogError("Incrementing votes", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{"error", "internal_server_error"})
		return
	}

	writeJSON(w, http.StatusOK, statusResponse{"ok"})
}
