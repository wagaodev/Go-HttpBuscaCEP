package main

import "net/http"

func main() {
	http.HandleFunc("/", buscaCEP)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func buscaCEP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, World!\n"))
}
