package main

import (
	"fmt"

	"github.com/DCH81/azuretls-client"
)

func main() {

	// Second way
	session := azuretls.NewSession()

	session.GetClientHelloSpec = azuretls.GetLastChromeVersion

	resp, err := session.Get("https://cfdata.lol/tools/connection/")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp.Body))
}
