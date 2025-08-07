package nexrev_utils

import (
	"archive/zip"
	"bytes"
	"crypto"
	"net/http"
	"os"
	"path"
	"strings"
)

func ErrorHTTP(w *http.ResponseWriter, code int, message string) {
	(*w).WriteHeader(code)
	(*w).Header().Set("Content-Type", "text/plain")
	(*w).Header().Set("X-Content-Type-Options", "nosniff")
	(*w).Write([]byte(message))
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

		hashset[file] = string(hash.Sum(nil))
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

	defer zipFile.Close()

	return zipBuffer.Bytes(), nil
}
