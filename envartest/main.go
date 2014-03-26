package main

import (
	"log"

	"github.com/partkyle/envar"
)

func main() {
	port := envar.Int("PORT", -1)
	var kport int
	envar.IntVar(&kport, "KPORT", -1)

	host := envar.String("HOST")
	var khost string
	envar.StringVar(&khost, "KHOST")

	debug := envar.Bool("DEBUG", false)
	var kdebug bool
	envar.BoolVar(&kdebug, "KDEBUG", true)

	stddev := envar.Float("STDDEV", -1)
	var kstddev float64
	envar.FloatVar(&kstddev, "KSTDDEV", -1)

	err := envar.Parse()
	if err != nil {
		log.Printf("error parsing err=%s", err)
	}

	log.Printf("got port=%d", *port)
	log.Printf("got kport=%d", kport)
	log.Printf("got host=%s", *host)
	log.Printf("got khost=%s", khost)
	log.Printf("got debug=%v", *debug)
	log.Printf("got kdebug=%v", kdebug)
	log.Printf("got stddev=%f", *stddev)
	log.Printf("got kstddev=%f", kstddev)
}
