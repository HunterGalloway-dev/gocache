package store

import (
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

func NewKVStore() *KVStore {
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
}

func (k *KVStore) InsertPersons(p []model.Person) {
	for _, person := range p {
		k.InsertPerson(person)
	}
}

func (k *KVStore) GetPerson(id int) (model.Person, bool) {
	p, ok := k.idIndex[id]
	return *p, ok
}

func (k *KVStore) GetAllPersons() []model.Person {
	return k.data
}

// Query KV store
func (k *KVStore) Query(email, name string, age []int) []model.Person {
	// BASE CASE: If all fields are empty, return all persons
	if email == "" && name == "" && len(age) == 0 {
		return k.GetAllPersons()
	}

	set := k.QuerySetBuilder(email, name, age)
	return buildSlice(set)
}

// for singular fields apply intersection
func (k *KVStore) QuerySetBuilder(email string, name string, age []int) map[*model.Person]bool {
	// build intersection sets first
	emailSet := buildSet(email, k.emailIndex)
	nameSet := buildSet(name, k.nameIndex)

	// intersection of email and name
	intersection := setIntersection(emailSet, nameSet)

	result := filterByAge(intersection, age)
	return result
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
