package nexrev_routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	nexrev_utils "nexusreveries/cdn/utils"
)

func DialogFilesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/zip")
	w.Header().Add("Content-Disposition", "attachment; filename=dialog-files.zip")

	if r.Header.Get("Content-Type") != "application/json" {
		nexrev_utils.ErrorHTTP(w, http.StatusBadRequest, "Invalid request type. Expected application/json")
		return
	}

	if r.Body == nil {
		nexrev_utils.ErrorHTTP(w, http.StatusBadRequest, "Request body is empty")
		return
	}

	missingFiles := []string{}
	err := json.NewDecoder(r.Body).Decode(&missingFiles)
	if err != nil {
		nexrev_utils.ErrorHTTP(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	for i := range missingFiles {
		missingFiles[i] = dialogs_dir + "/" + missingFiles[i]
	}

	zipBuf, err := nexrev_utils.ZipFiles(missingFiles)
	if err != nil {
		nexrev_utils.ErrorHTTP(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Add("Content-Length", fmt.Sprint(len(zipBuf)))
	w.Write(zipBuf)
}
