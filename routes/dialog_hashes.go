package nexrev_routes

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const dialogs_dir string = "./dialogs/"

func DialogHashesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//dirfs := os.DirFS(dialogs_dir)
	files, _ := os.ReadDir(dialogs_dir)
	hashes := make(map[string]string)

	for i := range files {
		file := files[i]
		println(file.Name())
		if file.IsDir() {
			continue
		}

		buffer, err := os.ReadFile(dialogs_dir + "/" + file.Name())

		if err != nil {
			fmt.Println(fmt.Errorf("an error occurred while reading file %s: %w", file.Name(), err))
			continue
		}

		fileHash := crypto.MD5.New()
		fileHash.Write(buffer)

		hashes[file.Name()] = hex.EncodeToString(fileHash.Sum(nil))
	}

	json.NewEncoder(w).Encode(hashes)
}
