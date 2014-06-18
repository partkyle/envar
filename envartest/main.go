package main

import (
	"log"
	"os"

	"github.com/partkyle/envar"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	port := envar.Int("PORT", -1, "simple port")
	var kport int
	envar.IntVar(&kport, "KPORT", -1, "simple port reference")

	host := envar.String("HOST", "localhost", "simple host")
	var khost string
	envar.StringVar(&khost, "KHOST", "localhost", "simple host reference")

	debug := envar.Bool("DEBUG", false, "simple debug")
	var kdebug bool
	envar.BoolVar(&kdebug, "KDEBUG", true, "simple debug reference")

	stddev := envar.Float("STDDEV", -1.2, "simple float")
	var kstddev float64
	envar.FloatVar(&kstddev, "KSTDDEV", -1.9, "simple float reference")

	err := envar.Parse()
	if err != nil {
		logger.Printf("error parsing err=%s", err)
	}

	logger.Printf("PORT=%d", *port)
	logger.Printf("KPORT=%d", kport)
	logger.Printf("HOST=%s", *host)
	logger.Printf("KHOST=%s", khost)
	logger.Printf("DEBUG=%v", *debug)
	logger.Printf("KDEBUG=%v", kdebug)
	logger.Printf("STDDEV=%f", *stddev)
	logger.Printf("KSTDDEV=%f", kstddev)
}
