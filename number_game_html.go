package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	//"github.com/etsy/statsd/examples/go"
)

type IndexVariables struct {
	Time string
}
type ResultData struct {
	Number       int
	RandomNumber int
	Value1       string
}

func main() {
	// t1 := time.Now() // Record a start time
	// // Create a new StatsD connection
	// host := "localhost"
	// port := 8080

	// client := statsd.New(host, port)

	http.HandleFunc("/", index)
	http.HandleFunc("/submit.html", submitHandler)
	// Increment a stat counter
	//client.Increment("stat.metric1")

	log.Fatal(http.ListenAndServe(":8080", nil))
	// Decrement a stat counter
	// client.Decrement("stat.metric1")

	// // Record an end time
	// t2 := time.Now()

	// // Submit timing information
	// duration := int64(t2.Sub(t1) / time.Millisecond)
	// client.Timing("stat.timer", duration)
}
func index(w http.ResponseWriter, r *http.Request) {
	//init the location
	loc, _ := time.LoadLocation("Europe/Berlin")

	//set timezone with location
	now := time.Now().In(loc)
	//now := time.Now()
	IndexVars := IndexVariables{ //store the data
		Time: now.Format("15:04:05"),
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, IndexVars)
	if err != nil { // error case
		log.Print("template executing error: ", err) //log it
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	number := r.FormValue("number")
	i, _ := strconv.Atoi(number)
	fmt.Println("Your number is:", i)
	randomNumber := randomInt(0, 11)
	fmt.Println("The random number is: ", randomNumber)
	value1 := "Next try!"

	//Error Handling
	if i >= 11 {
		fmt.Println("Wrong input")
		value1 = "Wrong input!\n Please enter a number between 0 and 10"
	}
	if randomNumber == i {
		fmt.Println(value1)
		value1 = "Got it"
	}
	calcIndexes := ResultData{i, randomNumber, value1}
	t, _ := template.ParseFiles("submit.html")
	t.Execute(w, calcIndexes)
}

//Int >= min oder <max handler func draus machen mit Query Parameter eingabe Zahl
//als result eine json mit dem submit inhalt
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)

}
