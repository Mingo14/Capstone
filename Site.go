package main

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    "net/url"
    // "bytes"
)

func main() {
    MakeRequest() 
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8080", nil)

}

func MakeRequest() {

	formData := url.Values{
		"name": {"masnun"},
	}

	resp, err := http.PostForm("https://api.themoviedb.org/3/movie/now_playing?api_key=8899bc310c15c2755e2703aed0345bc5", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
}
// "https://api.themoviedb.org/3/movie/now_playing?api_key=8899bc310c15c2755e2703aed0345bc5"