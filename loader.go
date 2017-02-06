package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("http://finance.yahoo.com/webservice/v1/symbols/allcurrencies/quote?format=json")
    if err != nil {
        fmt.Printf("Error %s", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error %s", err)
    }

    fmt.Printf("Body %s", body)
}
