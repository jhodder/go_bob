package main

import (
	"fmt"
	"github.com/go-statuspage-api"
	// "net/http"
	"encoding/json"
)

const (
	apiKey = "d6864cac-95e1-4830-866d-4ee5b81ccebe"
	// pageID = "nkpbndpckgcd" // test
	pageID = "704b94kplf0h" // prod
)

func main() {
	c, err := statuspage.NewClient(apiKey, pageID)
	if err != nil {
		// ...
	}

	// OpenIncidents
	i, err0 := c.GetOpenIncidents()
	if err0 != nil {
		// ...
	}

	// Print field names
	//fmt.Printf("%+v\n", i)

	// Parse i struct as json and print to screen
	// Using jsonpp from command line but how to do it in GO?
	i2J, _ := json.Marshal(i)
	fmt.Println(string(i2J))
}
