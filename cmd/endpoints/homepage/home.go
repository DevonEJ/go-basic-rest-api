package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"html/template"

	"github.com/gorilla/mux"
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

// Mock up welcome data
type Welcome struct {
	Name string
	Time string
	Mssg string
}

//homepageData prints a message to given I/O using user's name if available in url
func homepageContent(res http.ResponseWriter, req *http.Request) {

	// Mock up the default welcome message
	welcome := Welcome{Name: "Anonymous", Time: time.Now().Format(time.Stamp), Mssg: "Got veg?"}

	// Set HTML template to be used
	templates := template.Must(template.ParseFiles("/Users/devonedwardsjoseph/Documents/dev/repos/go-basic-rest-api/static/home.html"))

	// Tell go to also serve CSS files in static/ dir
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// is user's name available in url - if no, use default
	if name := req.FormValue("name"); name != "" {
		welcome.Name = name
	}

	if err := templates.ExecuteTemplate(res, "home.html", welcome); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}


}

//requestHandler maps endpoints to functions
func requestHandler() {

	// Define gorilla mux router for requests
	Router := mux.NewRouter().StrictSlash(true)

	// Map requests for the root to the homepageData function
	Router.HandleFunc("/", homepageContent)

	// Map requests for the get all the veggies function - only accessible via GET request
	Router.HandleFunc("/vegetables", getAllVegetables).Methods("GET")

	fmt.Println("Server running on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", Router))
}

func main() {
	requestHandler()
}
