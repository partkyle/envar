package envar

import "fmt"

func ExamplePointers() {
	host := String("INTERFACE", "localhost", "interface to listen on")
	port := Int("PORT", 8080, "port to listen on")
	verbose := Bool("VERBOSE", true, "print verbose logs")

	if *verbose {
		fmt.Printf("http://%s:%d", *host, *port)
	}
}

func ExampleReferences() {
	var host string
	var port int
	var verbose bool

	StringVar(&host, "INTERFACE", "localhost", "interface to listen on")
	IntVar(&port, "PORT", 8080, "port to listen on")
	BoolVar(&verbose, "VERBOSE", true, "print verbose logs")

	if verbose {
		fmt.Printf("http://%s:%d", host, port)
	}
}
