package main

import (
	"encoding/json"
	"fmt"
	"github.com/iotdog/json2table/j2t"
	"io/ioutil"
	"log"
	"net/http"
	)



func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := "https://api.punkapi.com/v2/beers?beer_name=india"
	response, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	responseString := string(responseData)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", responseString)

	ok, responseString := j2t.JSON2HtmlTable(responseString, []string{"name", "tagline","first_brewed","description","twist","food_pairing","brewers_tips"}, []string{"image_url"})
	if ok {
		fmt.Fprint(w, responseString)
	} else {
		fmt.Fprint(w, "failed")
	}


}


func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/beers", ServeHTTP)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func getJson(url string, target interface{}) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		//return err.Error()
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

