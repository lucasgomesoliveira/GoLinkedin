package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"

	jsonReturn := contentFromServer(url)

	fmt.Printf("%v\n", jsonReturn)
}

func contentFromServer(url string) string {

	resp, err := http.Get(url)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	bytes := resp.Body

	bytesReaded, err := ioutil.ReadAll(bytes)

	return string(bytesReaded)

}
