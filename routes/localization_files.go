package nexrev_routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	nexrev_utils "nexusreveries/cdn/utils"
)

func GetLocalizationsFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Add("Content-Disposition", "attachment; filename=localization-files.zip")

	if r.Body == nil {
		nexrev_utils.ErrorHTTP(w, http.StatusBadRequest, "Request body is missing.")
	}

	var missingFiles []string
	err := json.NewDecoder(r.Body).Decode(&missingFiles)
	if err != nil {
		nexrev_utils.ErrorHTTP(w, http.StatusBadRequest, "Request body is malformed.")
	}

	for i := range missingFiles {
		missingFiles[i] = loc_dir + "/" + missingFiles[i]
	}

	zipBuffer, err := nexrev_utils.ZipFiles(missingFiles)
	if err != nil {
		nexrev_utils.ErrorHTTP(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Set("Content-Length", fmt.Sprint(len(zipBuffer)))
	w.Write(zipBuffer)
}
