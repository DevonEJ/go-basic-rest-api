package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Veggie is a struct holding details for one vegetable
type Veggie struct {
	Name     string `json:"name"`
	Colour   string `json:"colour"`
	Calories int    `json:"calories"`
}

//Vegetables is an array used to hold individual Veggies
type Vegetables []Veggie

func getAllVegetables(res http.ResponseWriter, req *http.Request) {
	// Mock up an array of veggies to return
	veg := Vegetables{
		Veggie{Name: "carrot", Colour: "orange", Calories: 41},
		Veggie{Name: "broccoli", Colour: "green", Calories: 34},
		Veggie{Name: "edamame bean", Colour: "green", Calories: 1},
	}

	fmt.Println("You hit the endpoint: get all veggies!")

	// Send back the veggies array in JSON format
	json.NewEncoder(res).Encode(veg)

}

//homepageData prints a message to given I/O
func homepageContent(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "Hello - you have hit the homepage endpoint. Got veg?")
}

//requestHandler maps endpoints to functions
func requestHandler() {

	// Map requests for the root to the homepageData function
	http.HandleFunc("/", homepageContent)

	// Map requests for the get all the veggies function
	http.HandleFunc("/vegetables", getAllVegetables)

	// Log any errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	requestHandler()
}
