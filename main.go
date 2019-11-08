package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"
	"gopkg.in/src-d/go-git.v4"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Who is John Galt?")
}

func gitWebhook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body);
	log.Println(reqBody);
	var reqBodyUnmarshalled map[string]interface{};
	json.Unmarshall([]byte(reqBody), &reqBodyUnmarshalled);
	// Print json to play with it
	// use go git to download the repo to /tmp/<repo_name>
	fmt.Fprintf(w, "AAAAA: %+v", string(reqBody));
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/webhook/git", gitWebhook).Methods("POST");
	log.Fatal(http.ListenAndServe(":8080", router))
}
