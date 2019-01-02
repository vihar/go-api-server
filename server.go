package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type Post struct {
    ID string  `json:"id,omitempty"`
    Author string `json:"author,omitempty"`
    Title string
    Description string
    Category *Category
}
type Category struct {
    Name  string
    Description string
}

var posts []Post

func GetPost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range posts {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Post{})
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var post Post
    _ = json.NewDecoder(r.Body).Decode(&post)
    post.ID = params["id"]
    posts = append(posts, post)
    json.NewEncoder(w).Encode(posts)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range posts {
        if item.ID == params["id"] {
            posts = append(posts[:index], posts[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(posts)
    }
}

func main() {
    // Used to add routes to the application.
    router := mux.NewRouter()
    posts = append(posts, Post{ID: "1", Author: "John", Title: "Sample Title", Category: &Category{Name: "Science", Description: "Sample D"}})
    posts = append(posts, Post{ID: "2", Author: "John", Title: "Sample Title", Category: &Category{Name: "Science", Description: "Sample D"}})

    router.HandleFunc("/post", GetPost).Methods("GET")
    router.HandleFunc("/post/{id}", GetPost).Methods("GET")
    router.HandleFunc("/post/{id}", CreatePost).Methods("POST")
    router.HandleFunc("/post/{id}", DeletePost).Methods("DELETE")

    // Report and Start Server on Localhost 8000.
    fmt.Println("Running on Port 8000")
    log.Fatal(http.ListenAndServe(":9000", router))
}