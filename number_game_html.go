package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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
	http.HandleFunc("/", index)
	http.HandleFunc("/submit.html", submitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
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
	fmt.Println(i)
	randomNumber := randomInt(0, 11)
	fmt.Println("The random number is: ", randomNumber)
	value1 := "Next try!"

	if randomNumber == i {
		fmt.Println(value1)
		value1 = "Got it"
	}
	calcIndexes := ResultData{i, randomNumber, value1}
	t, _ := template.ParseFiles("submit.html")
	t.Execute(w, calcIndexes)
}

//Int >= min oder <max
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)

}
