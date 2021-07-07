package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Defer, res.Body.Close() get executed when exiting function
	// Useful to group open and close functions together
	// Run in LIFO (last-in, first-out) order
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(robots))

	// Panic
	// Use for unrecoverable events-cannot obtain TCP port for web server
	// If you don't handle panic application will shutdown
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		panic(err.Error())
	}

	// Recover
	// Used to recover from panics
	// Only useful in defered functions
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()

	panic("something bad happend")

}
