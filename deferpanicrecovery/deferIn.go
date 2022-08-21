package deferpanicrecovery

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DeferInGo() {
	fmt.Println("Start")
	fmt.Println("Mid 1")
	defer fmt.Println("Mid 2")
	defer fmt.Println("Mid 3")
	fmt.Println("End")
	a := "Star 1"
	a = "Star 2"
	defer fmt.Println(a)
	fmt.Println(a)
}

func DeferResourceManager() {
	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
