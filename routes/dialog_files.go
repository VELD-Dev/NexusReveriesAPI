package nexrev_routes

import (
	"archive/zip"
	"encoding/json"
	"io/fs"
	"net/http"
	nexrev_utils "nexusreveries/cdn/utils"
	"os"
	"slices"
)

func DialogFilesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/zip")
	w.Header().Add("Content-Disposition", "attachment; filename=dialog-files.zip")

	if r.Header.Get("Content-Type") != "application/json" {
		nexrev_utils.ErrorHTTP(&w, http.StatusBadRequest, "Invalid request type. Expected application/json")
		return
	}

	if r.Body == nil {
		nexrev_utils.ErrorHTTP(&w, http.StatusBadRequest, "Request body is empty")
		return
	}

	missingFiles := []string{}
	err := json.NewDecoder(r.Body).Decode(&missingFiles)
	if err != nil {
		nexrev_utils.ErrorHTTP(&w, http.StatusBadRequest, "Invalid request body")
		return
	}

	filesFs, _ := os.ReadDir(dialogs_dir)

	zipWriter := zip.NewWriter(w)

	for i := range missingFiles {
		missingFile := missingFiles[i]
		if !slices.ContainsFunc(filesFs, func(e fs.DirEntry) bool { return missingFile == e.Name() }) {
			continue
		}

		originalFile, _ := os.ReadFile(dialogs_dir + "/" + missingFile)
		zippedFile, _ := zipWriter.Create(missingFile)
		_, err = zippedFile.Write(originalFile)

		if err != nil {
			println(err.Error())
		}
	}

	defer zipWriter.Close()
}
