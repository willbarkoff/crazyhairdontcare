package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/willbarkoff/crazyhairdontcare/errlog"
)

type contestantsResponse struct {
	Status      string       `json:"status"`
	Contestants []contestant `json:"contestants"`
}

type contestant struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photoURL"`
}

func contestantsRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		errlog.LogError("getting session", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{"error", "internal_server_error"})
		return
	}

	loggedIn, ok := session.Values["loggedIn"].(bool)

	if !ok || !loggedIn {
		writeJSON(w, http.StatusUnauthorized, errorResponse{"error", "unauthorized"})
		return
	}

	rows, err := db.Query("SELECT id, name, photoURL FROM contestants")
	if err != nil {
		errlog.LogError("Getting contestants", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{"error", "internal_server_error"})
	}
	contestants := []contestant{}

	for rows.Next() {
		var id int
		var name, photoURL string
		rows.Scan(&id, &name, &photoURL)
		contestants = append(contestants, contestant{
			ID:       id,
			Name:     name,
			PhotoURL: photoURL,
		})
	}

	rows.Close()

	writeJSON(w, http.StatusOK, contestantsResponse{
		Status:      "ok",
		Contestants: contestants,
	})
}
