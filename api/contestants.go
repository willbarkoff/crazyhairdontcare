package api

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/willbarkoff/crazyhairdontcare/errlog"
)

type contestantsResponse struct {
	Status      string       `json:"status"`
	Contestants []contestant `json:"contestants"`
	Message     []string     `json:"message"`
}

type contestant struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PhotoURL          string `json:"photoURL"`
	OriginialPhotoURL string `json:"originalPhotoURL"`
	CutName           string `json:"cutName"`
}

func contestantsRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	rows, err := db.Query("SELECT id, name, photoURL, originalPhotoURL, cutName FROM contestants")
	if err != nil {
		errlog.LogError("Getting contestants", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{"error", "internal_server_error"})
	}
	contestants := []contestant{}

	for rows.Next() {
		var id int
		var name, photoURL, originalPhotoURL, cutName string
		rows.Scan(&id, &name, &photoURL, &originalPhotoURL, &cutName)
		fmt.Println(cutName)
		contestants = append(contestants, contestant{
			ID:                id,
			Name:              name,
			PhotoURL:          photoURL,
			OriginialPhotoURL: originalPhotoURL,
			CutName:           cutName,
		})
	}

	rows.Close()

	for i := range contestants {
		j := rand.Intn(i + 1)
		contestants[i], contestants[j] = contestants[j], contestants[i]
	}

	writeJSON(w, http.StatusOK, contestantsResponse{
		Status:      "ok",
		Contestants: contestants,
		Message:     message,
	})
}
