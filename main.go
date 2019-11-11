package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/libgit2/git2go"
	// "gopkg.in/src-d/go-git.v4"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Who is John Galt?")
}

func gitWebhook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body);
	var reqBodyUnmarshalled map[string]interface{};
	json.Unmarshal([]byte(reqBody), &reqBodyUnmarshalled);
	// Print json to play with it
	// use go git to download the repo to /tmp/<repo_name>
	fmt.Fprintf(w, "AAAAA: %+v", string(reqBody));
	fmt.Fprintf(w, "BBBBB: %+v", reqBodyUnmarshalled["ref"]); // Check how to 
	gitUrl:= "view/comments" // View how to get a nested value a[b][c]??
	pullUpdates(&gitUrl);
}

func pullUpdates(gitUrl *string) {
	// Check if folder exists in file system
	// if it does not, clone it
	// if it does, pull the new version

	// Get the info from the yaml conf
	// Apply using docker sdk

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/webhook/git", gitWebhook).Methods("POST");
	log.Fatal(http.ListenAndServe(":8080", router))
}
