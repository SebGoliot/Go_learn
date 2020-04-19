package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

type Article struct {
	Id string `json:"Id"`
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}
var Articles []Article

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World !")
    fmt.Println("Endpoint Hit: helloWorld")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnLastArticle(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnLastArticle")
	json.NewEncoder(w).Encode(Articles[len(Articles)-1])
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}


func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloWorld)
	router.HandleFunc("/articles", returnAllArticles)
	router.HandleFunc("/articles/last", returnLastArticle)
	router.HandleFunc("/article/{id}", returnSingleArticle)
	router.HandleFunc("/article", createNewArticle).Methods("POST")
    log.Fatal(http.ListenAndServe(":5050", router))
}

func main() {
    Articles = []Article{
        Article{Id: "0", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "1", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}
