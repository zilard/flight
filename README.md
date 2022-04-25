# Problem statement

## Story

There are over 100,000 flights a day, with millions of people and cargo being transferred
around the world. With so many people and different carrier/agency groups, it can be hard to
track where a person might be. In order to determine the flight path of a person, we must sort
through all of their flight records.

## Goal

To create a simple microservice API that can help us understand and track how a
particular person's flight path may be queried. The API should accept a request that includes a
list of flights, which are defined by a source and destination airport code. These flights may not
be listed in order and will need to be sorted to find the total flight paths starting and ending
airports.

## Required Structure

    [["SFO", "EWR"]] => ["SFO", "EWR"]
    
    [["ATL", "EWR"], ["SFO", "ATL"]] => ["SFO", "EWR"]
    
    [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]
    

# Format of API Endpoint

## Request

`POST /track <LIST OF FLIGHTS>`

    curl -i -X POST -H "Content-Type:application/json" \
    http://localhost:8080/track \
    -d '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'

## Response
    
    HTTP/1.1 200 OK
    Date: Sun, 24 Apr 2022 11:42:19 GMT
    Content-Length: 14
    Content-Type: text/plain; charset=utf-8

    ["SFO","EWR"]


# Lifecycle

## Build

    make image


## Run 

    make run


## Stop

    make stop


## Cleanup

    make clean


