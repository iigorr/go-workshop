package main

import (
	ab "../3-data/addressbook"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type jsonEntry struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

var book ab.AddressBook

func main() {
	r := mux.NewRouter()

	book = ab.New()
	book.Add(ab.NewEntry("Gopher", "Google Street"))

	r.HandleFunc("/entries", listEntriesHandler).Methods("GET")
	r.HandleFunc("/entries", addEntryHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}

func listEntriesHandler(w http.ResponseWriter, r *http.Request) {
	var entrylist []ab.Entry
	if r.FormValue("filter") == "" {
		entrylist = book.List()
	} else {
		entrylist = book.Find(r.FormValue("filter"))
	}

	entries := getJsonEntries(entrylist)

	data, _ := json.MarshalIndent(entries, "", "  ")
	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func addEntryHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var entry jsonEntry
	if err = json.Unmarshal(body, &entry); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to decode data")
		return
	}

	book.Add(ab.NewEntry(entry.Name, entry.Address))
	w.WriteHeader(http.StatusCreated)
}

func getJsonEntries(bookEntries []ab.Entry) []jsonEntry {

	entries := make([]jsonEntry, len(bookEntries))

	for i, be := range bookEntries {
		entries[i] = jsonEntry{be.Name(), be.Address()}
	}

	return entries
}
