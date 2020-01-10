// GCI 2019 MiskaKyto
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Foods API!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/foods", returnAllFoods)
	myRouter.HandleFunc("/food", createNewFood).Methods("POST")
	// add our new DELETE endpoint here
	myRouter.HandleFunc("/food/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/food/{id}", returnSingleFood)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllFoods(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Foods)
}

func returnSingleFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Foods {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewFood(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var food Food
	json.Unmarshal(reqBody, &food)
	// update our global Articles array to include
	// our new Article
	Foods = append(Foods, food)

	json.NewEncoder(w).Encode(food)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range Foods {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			Foods = append(Foods[:index], Foods[index+1:]...)
		}
	}

}

type Food struct {
	Id      string `json:"Id`
	Title   string `json:"Title`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Foods []Food

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Foods = []Food{
		Food{
			Id:      "1",
			Title:   "Banana",
			Desc:    "🍌",
			Content: "Bananas are yellow",
		},
		Food{
			Id:      "2",
			Title:   "Apple",
			Desc:    "🍏",
			Content: "Apples are green",
		},
		Food{
			Id:      "3",
			Title:   "Cherry",
			Desc:    "🍒",
			Content: "Cherries are red",
		},
	}
	handleRequests()
}
