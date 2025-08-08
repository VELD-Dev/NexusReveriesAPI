package nexrev_routes

import (
	"encoding/json"
	"net/http"
	nexrev_utils "nexusreveries/cdn/utils"
)

const loc_dir string = "./localizations/"

func GetLocalizationsHashes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hashes, err := nexrev_utils.GetHashesForFiles(loc_dir, ".json")
	if err != nil {
		nexrev_utils.ErrorHTTP(w, http.StatusInternalServerError, err.Error())
	}

	json.NewEncoder(w).Encode(hashes)
}
