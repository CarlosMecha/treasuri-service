package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
)

type Currency struct {
    name, price, symbol, ts, res_type, utctime, volume string
}

type Resource struct {
    classname string
    fields Currency
}

type Metadata struct {
    res_type string
    start, count int
}

type ResourceList struct {
    meta Metadata
    resources []Resource
}

type CurrencyData struct {
    list ResourceList
}

func Request() (*http.Response, error) {
    resp, err := http.Get("http://finance.yahoo.com/webservice/v1/symbols/allcurrencies/quote?format=json")
    if err != nil {
        fmt.Printf("Error %s", err)
        return nil, err
    }
    return resp, nil
}

func Transform(content io.ReadCloser, data *CurrencyData) (error) {
    err := json.NewDecoder(content).Decode(data)
    if err != nil {
        fmt.Printf("Error decoding %s", err)
        return err
    }
    return nil
}

func main() {
    resp, err := Request()
    if err != nil {
        os.Exit(2)
    }

    defer resp.Body.Close()

    var data CurrencyData
    err = Transform(resp.Body, &data)
    if err != nil {
        os.Exit(2)
    }

    fmt.Printf("Data %q", data.list.resources)

}
