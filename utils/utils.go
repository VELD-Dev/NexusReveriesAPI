package nexrev_utils

import (
	"archive/zip"
	"bytes"
	"crypto"
	"encoding/hex"
	"net/http"
	"os"
	"path"
	"strings"
)

// ErrorHTTP writes an error message to the response writer.
// It's better to avoid using panic for control flow in HTTP handlers.
func ErrorHTTP(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write([]byte(message))
	println(message)
}

func GetHashesForFiles(directory string, fileType string) (map[string]string, error) {
	hashset := make(map[string]string)
	files, err := os.ReadDir(directory)

	if err != nil {
		return nil, err
	}

	filteredFiles := []string{}
	for i := range files {
		file := files[i]
		if strings.HasSuffix(file.Name(), fileType) {
			filteredFiles = append(filteredFiles, file.Name())
		}
	}

	for i := range filteredFiles {
		file := filteredFiles[i]
		fileBytes, err := os.ReadFile(directory + "/" + file)
		if err != nil {
			return nil, err
		}

		hash := crypto.MD5.New()
		_, err = hash.Write(fileBytes)
		if err != nil {
			return nil, err
		}

		hashset[file] = hex.EncodeToString(hash.Sum(nil))
	}

	return hashset, nil
}

func ZipFiles(filepaths []string) ([]byte, error) {
	zipBuffer := bytes.Buffer{}
	zipFile := zip.NewWriter(&zipBuffer)

	for i := range filepaths {
		filebuf, err := os.ReadFile(filepaths[i])
		if err != nil {
			return nil, err
		}

		_, filename := path.Split(filepaths[i])
		zippedFileBuf, err := zipFile.Create(filename)
		if err != nil {
			return nil, err
		}

		_, err = zippedFileBuf.Write(filebuf)
		if err != nil {
			return nil, err
		}
	}

	err := zipFile.Close()
	if err != nil {
		return nil, err
	}

	return zipBuffer.Bytes(), nil
}
