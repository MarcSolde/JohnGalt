package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

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
	// For now, assume the repo is already downloaded
	var prevCommit *object.Commit;
	var prevTree *object.Tree;

	var currentCommit *object.Commit;
	var currentTree *object.Tree;
	// Pull changes
	r, err := git.PlainOpen("./.git");
	CheckIfError(err);
	w, err := r.Worktree();
	CheckIfError(err);
	err = w.Pull(&git.PullOptions{RemoteName: "origin"});
	CheckIfError(err);

	commits, err := repo.Log({}) //This returns a CommitIterator
	defer commits.Close();

	// Get the info from the yaml conf
	// Apply using docker sdk

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/webhook/git", gitWebhook).Methods("POST");
	log.Fatal(http.ListenAndServe(":8080", router))
}
