package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method:", r.Method)
		fmt.Println("URL:", r.URL)
		fmt.Println("Headers:", r.Header)

		fmt.Fprintln(w, `{"stats": "ok"}`)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("This is bad: " + err.Error())
	}
}
