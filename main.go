package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/globaspect/goresttrace/rest"
)

func main() {

	fmt.Println("gatewayComTracer started ...")
	port := ""
	if os.Args != nil && len(os.Args) >= 1 {
		port = os.Args[1]
		if port != "" {
			fmt.Println("Port: " + port)
		} else {
			port = "14001"
			fmt.Println("default port 14001")
		}
	} else {
		port = "14001"
		fmt.Println("default port 14001")
	}

	router := rest.NewRouter()
	routerPort := ":" + port
	log.Fatal(http.ListenAndServe(routerPort, router))
	fmt.Printf("... started!")

}
