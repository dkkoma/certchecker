package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
)

func main() {
	var site string

	flag.StringVar(&site, "site", "", "check site FQDN")
	flag.Parse()

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	if site == "" {
		log.Fatal("site is missing")
	}

	conn, err := tls.Dial("tcp", site+":443", conf)
	if err != nil {
		log.Println("Error in Dial", err)
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		fmt.Printf("Subject CN: %s \n", cert.Subject.CommonName)
		fmt.Printf("Issuer: %s\n", cert.Issuer)
		fmt.Printf("Expiry: %s \n", cert.NotAfter.Format("2006-January-02"))
		fmt.Printf("\n")

	}
}
