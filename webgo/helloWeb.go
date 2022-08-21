package webgo

import (
	"fmt"
	"net/http"
)

func HelloWebGO() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go! This is Asutosh"))
	})
	fmt.Println("Spinning Up Server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
