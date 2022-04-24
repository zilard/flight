package main

import (
    "encoding/json"
    "net/http"
    "log"
    "strconv"
    "fmt"
    "os"

    "github.com/gorilla/mux"
)

const PORT int = 8080

type airportStat struct {
    occurrence int
    isSrc bool
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/track", TrackFlight).Methods("POST")
    fmt.Printf("Server listening on :%d\n", PORT)
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(PORT), r))
}

func TrackFlight(w http.ResponseWriter, r *http.Request) {
    var flightList [][]string
    result := json.NewDecoder(r.Body).Decode(&flightList)
    if result != nil {
        fmt.Fprintf(os.Stderr, "error=%v\n", result)
        return
    }
    srcDestAirport := findSrcDest(flightList)
    json.NewEncoder(w).Encode(srcDestAirport)
}

func findSrcDest(flightList [][]string) [2]string {
    airportMap := make(map[string]*airportStat)
    for _, flight := range flightList {
        for i, airport := range flight {
            if _, ok := airportMap[airport]; ok {
                airportMap[airport].occurrence++
                airportMap[airport].isSrc = i == 0
            } else {
                airportMap[airport] = &airportStat{occurrence: 1, isSrc: i == 0}
            }
        }
    }
    var srcDestAirport [2]string
    for k, v := range airportMap {
        if v.occurrence == 1 {
            if v.isSrc {
                srcDestAirport[0] = k
            } else {
                srcDestAirport[1] = k
            }
        }
    }
    return srcDestAirport
}
