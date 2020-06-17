package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/marcobellaccini/catchcert/certfun"
)

// program version
const version string = "v0.0.1"

// usage examples
const usageEx = `Connect to a server using TLS and get PEM-encoded server certificate.
Usage:
catchcert <SERVER> [<PORT>]

Examples:
catchcert wikipedia.org
catchcert server.contoso.com 8443
catchcert github.com > github.pem
catchcert dc.contoso.com 636 > contoso_dc.crt

Check https://github.com/marcobellaccini/catchcert for more information.`

func main() {
	verFlag := flag.Bool("version", false, "print version")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageEx)
		fmt.Fprintf(os.Stderr, "\n\nOptions:\n")

		flag.PrintDefaults()
	}

	flag.Parse()

	if *verFlag {
		fmt.Println("catchcert", version)
		os.Exit(0)
	}

	args := flag.Args()

	var port string = "443" // default port

	if len(args) == 2 {
		port = args[1]
		if _, err := strconv.Atoi(port); err != nil {
			fmt.Println(usageEx)
			os.Exit(101)
		}
	} else if len(args) > 2 || len(args) == 0 {
		fmt.Println(usageEx)
		os.Exit(102)
	}

	tgt := args[0] // target server

	cert, err := certfun.GetCertsPEM(tgt + ":" + port)
	if err != nil {
		fmt.Println("Unable to get certificate for", tgt, "port", port)
		os.Exit(201)
	} else {
		fmt.Println(cert)
	}
}
