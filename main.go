package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Document struct
type Document struct {
	ID   string
	Name string
	Size int
}

func main() {
	readFiles()
	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocuments).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
	var docs []Document
	docs = append(docs,
		Document{ID: "doc-1", Name: "Report.docx", Size: 1500})
	docs = append(docs,
		Document{ID: "doc-2", Name: "Sheet.xlsx", Size: 5000})
	docs = append(docs,
		Document{ID: "doc-3", Name: "Container.tar", Size: 50000})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}

func readFiles() {
	path := "./files/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		b, err := ioutil.ReadFile(path + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		fileContent := string(b) // convert file content into a 'string'
		fmt.Println(f.Name())
		fmt.Println(getMD5Checksum(fileContent))
	}
}

func getMD5Checksum(content string) string {
	hasher := md5.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}
