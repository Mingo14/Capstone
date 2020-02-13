package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/thedevsaddam/gojsonq"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "main", &Page{Title: "Welcome to TL;DR"})

	fmt.Println("Endpoint Hit: homePage")
	if fileExists("test.json") {
		log.Print("it exisits")
	} else {
		d := apiRequest()
		file, _ := json.MarshalIndent(d, "", " ")

		_ = ioutil.WriteFile("test.json", file, 0644)
	}
	// movieMap := make(map[string]string)
	temp := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < 10; i++ {
		jq := gojsonq.New().File("./test.json")
		title := jq.Find("results.[" + temp[i] + "].original_title")
		jq2 := gojsonq.New().File("./test.json")
		des := jq2.Find("results.[" + temp[i] + "].overview")

		if jq.Error() != nil {
			log.Fatal(jq.Errors())
		}
		if jq2.Error() != nil {
			log.Fatal(jq.Errors())
		}
		fmt.Fprintln(w, title)
		fmt.Fprintln(w, des)
		// titleStr := fmt.Sprintf("%v", title)
		// desStr := fmt.Sprintf("%v", des)
		// movieMap[titleStr] = desStr
	}

	// test := dom.GetWindow().Document()
	// e1 := test.GetElementByID("movieTitle")
	// e2 := test.GetElementByID("movieDes")

}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About")
	templ.ExecuteTemplate(w, "about", &Page{Title: "About TL;DR"})
}

func apiRequest() map[string]interface{} {
	log.Println("pulling api data")
	url := "https://api.themoviedb.org/3/movie/now_playing?api_key=8899bc310c15c2755e2703aed0345bc5"
	timeStart := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(time.Since(timeStart), url)
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	// log.Println(result)
	enc := json.NewEncoder(os.Stdout)
	d := result
	enc.Encode(d)
	fmt.Println("The JSON data is:")
	fmt.Println(d)

	return d
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	log.Println("hlllo")
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

var templ = func() *template.Template {
	t := template.New("")
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			fmt.Println(path)
			_, err = t.ParseFiles(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err)
	}
	return t
}()

// Page ... structure for web page, in this case just a title
type Page struct {
	Title string //title of the page
}

// "https://api.themoviedb.org/3/movie/now_playing?api_key=8899bc310c15c2755e2703aed0345bc5"
