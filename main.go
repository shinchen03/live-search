package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
	"net/http"
	"text/template"
	"io/ioutil"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func test(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello World !!", rand.Intn(100))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(sqrt(-2))
}

type Page struct {
	Title string
	Body string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World.", getBody()}
	tmpl, err := template.ParseFiles("test.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func getBody() string{
	url := "http://hall.zepp.co.jp/tokyo/schedule.html"
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}

func main() {
	http.HandleFunc("/", viewHandler) // hello
	http.ListenAndServe(":8080", nil)
}
