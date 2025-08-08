package nexrev_routes

import (
	"encoding/json"
	"net/http"
	nexrev_utils "nexusreveries/cdn/utils"
)

const dialogs_dir string = "./dialogs/"

func DialogHashesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	hashes, err := nexrev_utils.GetHashesForFiles(dialogs_dir, ".json")

	if err != nil {
		nexrev_utils.ErrorHTTP(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(hashes)
}
