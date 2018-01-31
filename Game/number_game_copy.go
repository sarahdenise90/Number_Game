package main

import (
	"strconv"
	//"cmd/go/internal/web"
	"fmt"

	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type IndexVariables struct {
	Time   string
	Number int64
	input  int
}
type Page struct {
	Title  string
	Body   []byte
	number int
	input  int
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

//Int >= min oder <max
//func randomInt(min, max int) int {
//	return min + rand.Intn(max-min)
//}

//Speicherung der html Seite
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//Laden der HTML Seite
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}
func submitHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/submit"):]
	// p, err := loadPage(title)

	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	p = &Page{Title: title}
	// }

	r.ParseForm()
	number := r.FormValue("number")
	fmt.Printf("number: %s", number)
	i, _ := strconv.ParseInt(number, 10, 0)
	fmt.Printf("integer")
	fmt.Println(i)
	calcIndexes := IndexVariables{"Time", i, 500}
	t, _ := template.ParseFiles("submit.html")
	t.Execute(w, calcIndexes)
}

//Variable "number" aus html
/*func input(c web.C, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
	}
	number := r.PostFormValue("number")
	fmt.Fprintf(w, "Variable: %s!", number)
}*/

/*var input int
	fmt.Print("Please enter your number: ")
	fmt.Scan(&input)
	fmt.Println("Your number", input)

	//Returns the now random number
	rand.Seed(time.Now().UnixNano())
	var randomNumber = randomInt(0, 11)
	fmt.Println("The random number is: ", randomNumber)

	//Compare the random number with your input
	if randomNumber == input { //equal
		fmt.Println("You got it!")
	} else { //not equal
		fmt.Println("The right number is: ", randomNumber, ", Next try!")
	}
}*/
