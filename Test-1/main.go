// filename: main.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//handler functions
//home endpoint
func home(w http.ResponseWriter, r *http.Request) {
	//we use the servefile function to serve our html file to the client
	http.ServeFile(w,r,"index.html")
}

//second handler function
//greeting endpoint
func greetings(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now() //time function to present our time
	//create another variable that when used will give us our date example "wednesday"
	date := time.Now()
	//format our date and time as request by instructor 
	formattedtime := currentTime.Format("3:04 PM")
	formatTime := date.Format("Monday")
	fmt.Fprintf(w,"Right now is %s\n", formattedtime)
	fmt.Fprintf(w, "Enjoy the rest of your %s\n", formatTime)
}

//third handler function
//random endpoint
func random(w http.ResponseWriter, r *http.Request) {
	//create a map that contains quotes to randomize and presnet a 
	//different qoute everytime a request is sent
	qoute := map[int]string {
		1: "Many Of Life's Failures Are People Who Did Not Realize How Close They Were To Success When They Gave Up. - (Thomas A. Edison)",
		2: "Don't Be The Person Who Says Yes With Their Mouth But No With Their Actions. - (Ryan Holiday)",
		3: "Life Is A Long Lesson In Humility. - (James M. Barrie)",
		4: "In Three Words I Can Sum Up Everything I've Learned About Life: It Goes On. - (Robert Frost)",
		5: "We Have Two Lives, And The Second Begins When We Realize We Only Have One. - (Confucius)",
	}

	//we use to basically give us a unique random number sequence everytime we run the program 
	rand.Seed(time.Now().UnixNano())

	//generate a random integer used to match up with the qoute in the map qoutes 
	response := qoute[rand.Intn(len(qoute))+1]

	//write a response to the client when a request is made
	fmt.Fprintf(w,"%s",response)


}
func main() {
	//our multiplexer/router
	//we create a newserver mux instance to handle our http requests
	mux := http.NewServeMux()
	mux.HandleFunc("/home",home)
	mux.HandleFunc("/greetings",greetings)
	mux.HandleFunc("/random",random)

	//we start a server on port 8080
	log.Print("staring server on: 8080")
	err := http.ListenAndServe(":8080",mux)
	log.Fatal(err)


}