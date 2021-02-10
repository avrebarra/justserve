package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/leaanthony/clir"
)

var cmd *clir.Cli

func main() {
	quiet := false
	portint := 5000
	location := "."

	// setup commands
	cmd = clir.NewCli("justserve", "just serve a static file", "v1")
	cmd.BoolFlag("quiet", "perform quiet operation", &quiet)
	cmd.StringFlag("location", "static files location", &location)
	cmd.IntFlag("port", "port to bind", &portint)

	// default function
	cmd.Action(func() (err error) {
		// set output
		setuplog(quiet)

		// start static server
		fs := http.FileServer(http.Dir(location))
		http.Handle("/", fs)

		port := fmt.Sprintf(":%d", portint)
		log.Println(fmt.Sprintf("Listening on http://localhost%s...", port))
		if err = http.ListenAndServe(port, nil); err != nil {
			return
		}
		return
	})

	// run server
	err := cmd.Run()
	if err != nil {
		fmt.Println("unexpected error:", err.Error())
		return
	}
}

func setuplog(quiet bool) {
	if quiet {
		log.SetOutput(ioutil.Discard)
	}
	log.SetFlags(0)
}
