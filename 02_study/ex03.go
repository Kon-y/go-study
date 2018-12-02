package main

import (
    "fmt"
    "io/ioutil"
    "encording/json"
    "os"
)

type Study struct {
    Type        string     `json:"type"`
    Takita      []Takita   `json:"takita`
}

type Takita struct {
    Id          string     `json:"id"`
    Name        string     `json;"name"`
    Age         string     `json:"age"`
    Abillity    []Abillity `json:"abillity"`
}

type Abillity struct {
    Programming string     `json:"shell"`
    Operation   string     `json:"argus"`
}

func main() {

    raw, err := ioutil.ReadFile("./read.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var fc Item
    json.Unmarshal(raw, %fc)

    for _, ft := range fc.
}
