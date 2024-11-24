package store

import (
	"errors"
	"fmt"
	"gocache/pkg/model"
)

// KVStore is a simple in-memory key-value store
type KVStore struct {
	data []model.Person
	// Index singular fields
	idIndex    map[int]*model.Person
	nameIndex  map[string][]*model.Person
	emailIndex map[string][]*model.Person
}

func NewKVStore() PersonStore {
	return &KVStore{
		data:       make([]model.Person, 0),
		idIndex:    make(map[int]*model.Person),
		nameIndex:  make(map[string][]*model.Person),
		emailIndex: make(map[string][]*model.Person),
	}
}

func (k *KVStore) InsertPerson(p model.Person) {
	k.data = append(k.data, p)
	k.idIndex[p.ID] = &p
	k.nameIndex[p.Name] = append(k.nameIndex[p.Name], &p)
	k.emailIndex[p.Email] = append(k.emailIndex[p.Email], &p)
}

func (k *KVStore) InsertPersons(p []model.Person) {
	for _, person := range p {
		k.InsertPerson(person)
	}
}

func (k *KVStore) GetPerson(id int) (model.Person, bool) {
	p, ok := k.idIndex[id]
	if !ok {
		return model.Person{}, false
	}

	return *p, ok
}

func (k *KVStore) GetAllPersons() []model.Person {
	return k.data
}

// Delete a person by ID
func (k *KVStore) DeletePerson(id int) error {
	person, ok := k.idIndex[id]
	if !ok {
		return errors.New("person not found")
	}

	var indexToDelete int
	for i, p := range k.data {
		if p.ID == id {
			indexToDelete = i
			break
		}
	}
	k.data = append(k.data[:indexToDelete], k.data[indexToDelete+1:]...)

	delete(k.idIndex, id)

	for i, p := range k.nameIndex[person.Name] {
		if p.ID == id {
			k.nameIndex[person.Name] = append(k.nameIndex[person.Name][:i], k.nameIndex[person.Name][i+1:]...)
			break
		}
	}

	for i, p := range k.emailIndex[person.Email] {
		if p.ID == id {
			k.emailIndex[person.Email] = append(k.emailIndex[person.Email][:i], k.emailIndex[person.Email][i+1:]...)
			break
		}
	}

	return nil
}

// Update a person by ID
func (k *KVStore) UpdatePerson(updatedPerson model.Person) error {
	id := updatedPerson.ID
	existingPerson, ok := k.idIndex[id]
	if !ok {
		return errors.New("person not found")
	}

	for i, p := range k.data {
		if p.ID == id {
			k.data[i] = updatedPerson
			break
		}
	}

	delete(k.nameIndex, existingPerson.Name)
	delete(k.emailIndex, existingPerson.Email)

	// Insert the updated person
	k.InsertPerson(updatedPerson)

	return nil
}

// Query KV store
func (k *KVStore) Query(name, email string, age []int) []model.Person {
	// BASE CASE: If all fields are empty, return all persons
	if email == "" && name == "" && len(age) == 0 {
		return k.GetAllPersons()
	}

	set := k.querySetBuilder(name, email, age)
	return buildSlice(set)
}

// for singular fields apply intersection
func (k *KVStore) querySetBuilder(name, email string, age []int) map[*model.Person]bool {
	// build intersection sets first
	nameSet := buildSet(name, k.nameIndex)
	emailSet := buildSet(email, k.emailIndex)

	// intersection of email and name
	intersection := make(map[*model.Person]bool, len(k.data))

	if email == "" && name == "" {
		for _, p := range k.data {
			intersection[&p] = true
		}
	} else {
		intersection = setIntersection(emailSet, nameSet)
	}

	result := filterByAge(intersection, age)
	return result
}

func (k *KVStore) String() string {
	ret := "KVStore\n"
	for _, p := range k.data {
		ret += fmt.Sprintf("%+v\n", p)
	}
	for id, p := range k.idIndex {
		ret += fmt.Sprintf("ID: %d, Person: %+v\n", id, *p)
	}

	ret += "\nName Index\n"
	for name, persons := range k.nameIndex {
		ret += fmt.Sprintf("Name: %s\n", name)
		for _, p := range persons {
			ret += fmt.Sprintf("\tPerson: %+v\n", *p)
		}
	}

	ret += "\nEmail Index\n"
	for email, persons := range k.emailIndex {
		ret += fmt.Sprintf("Email: %s\n", email)
		for _, p := range persons {
			ret += fmt.Sprintf("\tPerson: %+v\n", *p)
		}
	}

	return ret
}

func filterByAge(set map[*model.Person]bool, age []int) map[*model.Person]bool {
	if len(age) == 0 {
		return set
	}

	ageIndex := make(map[int]bool)
	for _, a := range age {
		ageIndex[a] = true
	}

	result := make(map[*model.Person]bool, len(set))
	for p := range set {
		if ageIndex[p.Age] {
			result[p] = true
		}
	}

	return result
}

func buildSlice(set map[*model.Person]bool) []model.Person {
	result := make([]model.Person, 0, len(set))

	for p := range set {
		result = append(result, *p)
	}

	return result
}

func buildSet(key string, index map[string][]*model.Person) map[*model.Person]bool {
	set := make(map[*model.Person]bool)

	if key == "" {
		return set
	}

	for _, p := range index[key] {
		set[p] = true
	}

	return set
}

// Helper function to find the intersection of two sets
func setIntersection(set1, set2 map[*model.Person]bool) map[*model.Person]bool {
	if len(set1) == 0 {
		return set2
	}

	if len(set2) == 0 {
		return set1
	}

	result := make(map[*model.Person]bool)
	for p := range set1 {
		if set2[p] {
			result[p] = true
		}
	}
	return result
}
