# 5 web

## Exercise 5.1

Implement a simple web interface (json) for the address book from exercise 4.1. 

Make the following work:
curl localhost:8080/entries (list all entries)
curl localhost:8080/entries?filter={filter} (finds all entries with defined substring)
curl -x POST localhost:8080/entries -d '{ name: "Another Gopher", address: "Google Street 1" }' (add entry)


Hints:
* use github.com/gorilla/mux
* you may have to convert address book entries (interface type), so you can marshal them via json
* is your service thread-safe? try to break your service by many request. e.g.
  `for i in $(seq 1 100); do curl -X POST localhost:8080/entries -d '{ name: "Gopher", address: "Google Street 1" }'; done`

## Variant 1
Serve static data to deploy a simple JS frontend

## Variant 2
embed data using
see https://github.com/jteeuwen/go-bindata