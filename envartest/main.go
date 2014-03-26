package main

import (
	"log"

	"github.com/partkyle/envar"
)

func main() {
	port := envar.Int("PORT", -1, "simple port")
	var kport int
	envar.IntVar(&kport, "KPORT", -1, "simple port reference")

	host := envar.String("HOST", "simple host")
	var khost string
	envar.StringVar(&khost, "KHOST", "simple host reference")

	debug := envar.Bool("DEBUG", false, "simple debug")
	var kdebug bool
	envar.BoolVar(&kdebug, "KDEBUG", true, "simple debug reference")

	stddev := envar.Float("STDDEV", -1.2, "simple float")
	var kstddev float64
	envar.FloatVar(&kstddev, "KSTDDEV", -1.9, "simple float reference")

	err := envar.Parse()
	if err != nil {
		log.Printf("error parsing err=%s", err)
	}

	envar.Usage()

	log.Printf("got port=%d", *port)
	log.Printf("got kport=%d", kport)
	log.Printf("got host=%s", *host)
	log.Printf("got khost=%s", khost)
	log.Printf("got debug=%v", *debug)
	log.Printf("got kdebug=%v", kdebug)
	log.Printf("got stddev=%f", *stddev)
	log.Printf("got kstddev=%f", kstddev)
}
