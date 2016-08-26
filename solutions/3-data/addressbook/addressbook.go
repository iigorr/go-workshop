package addressbook

import (
	"strings"
)

// represents an address book entry
type Entry interface {
	Name() string
	Address() string
}

type entry struct {
	name    string
	address string
}

func (e *entry) Name() string {
	return e.name
}

func (e *entry) Address() string {
	return e.address
}

func NewEntry(name, address string) Entry {
	return &entry{name, address}
}

// The address book interface
type AddressBook interface {
	Add(entry Entry)
	Delete(name string)
	List() []Entry
	Find(needle string) []Entry
}

type addressBook struct {
	data map[string]Entry
}

func New() AddressBook {
	return &addressBook{
		data: make(map[string]Entry),
	}
}

func (ab *addressBook) Add(entry Entry) {
	key := strings.ToLower(entry.Name())
	ab.data[key] = entry
}

func (ab *addressBook) Delete(name string) {
	delete(ab.data, strings.ToLower(name))
}

func (ab *addressBook) List() []Entry {
	list := make([]Entry, len(ab.data))
	i := 0
	for _, entry := range ab.data {
		list[i] = entry
		i++
	}

	return list

}

func (ab *addressBook) Find(needle string) []Entry {
	needle = strings.ToLower(needle)

	found := make([]Entry, 0)

	for key, entry := range ab.data {
		if strings.Contains(key, needle) {
			found = append(found, entry)
		}
	}

	return found
}
