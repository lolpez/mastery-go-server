package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Document struct
type Document struct {
	ID   string
	Name string
	Size int64
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocuments).Methods("GET")
	router.HandleFunc("/documents/{id}", getDocumentById).Methods("GET")
	router.HandleFunc("/documents", setDocument).Methods("POST")
	router.HandleFunc("/documents/{id}", deleteDocumentById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
	var docs []Document
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
		docs = append(docs, Document{ID: getMD5Checksum(fileContent), Name: f.Name(), Size: f.Size()})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}

func getMD5Checksum(content string) string {
	hasher := md5.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getDocumentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id := vars["id"]
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
		fileContent := string(b)
		if getMD5Checksum(fileContent) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Document{ID: getMD5Checksum(fileContent), Name: f.Name(), Size: f.Size()})
		}
	}
}

func setDocument(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func deleteDocumentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id := vars["id"]
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
		fileContent := string(b)
		fmt.Println(path + f.Name())
		if getMD5Checksum(fileContent) == id {
			fullpath := path + f.Name()
			err := os.Remove(fullpath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
