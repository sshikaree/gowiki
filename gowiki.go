package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Pageid  int64  `json:"pageid"`
	Ns      int    `json:"ns"`
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

type Pages map[string]Page

type Query struct {
	Pages Pages `json:"pages"`
}

type Response struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         Query  `json:"query"`
}

func main() {
	log.SetFlags(log.Lshortfile)
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Wrong request. Exiting")
		return
	}
	t := args[0]
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://en.wikipedia.org/w/api.php", nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("action", "query")
	q.Add("prop", "extracts")
	q.Add("exintro", "1")
	q.Add("explaintext", "1")
	q.Add("redirects", "1")
	q.Add("titles", t)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jr := new(Response)

	err = json.Unmarshal(data, jr)
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range jr.Query.Pages {
		fmt.Println(val.Extract)
	}

}
