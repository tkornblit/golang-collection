package collection

import (
	set "github.com/deckarep/golang-set"
)

type Collection []interface{}

func NewCollection() Collection {
	return make(Collection, 0)
}

// Creates and returns a reference to a collection from an existing slice
func NewCollectionFromSlice(s []interface{}) Collection {
	a := NewCollection()
	for _, item := range s {
		a = append(a, item)
	}
	return a
}

func NewCollectionFromSet(set set.Set) (collection Collection) {
	collection = NewCollection()
	var i int64

	for item := range set {
		collection = append(collection, item)
		i++
	}

	return

}

func (collection Collection) Chunks(count int64) (c chan interface{}) {
	c = make(chan interface{})

	go func() {

		batch := []interface{}{}

		var i int64

		for _, item := range collection {
			i++
			if i%count == 0 {

				c <- batch
				batch = nil

			}

			batch = append(batch, item)

		}

		close(c)

	}()

	return
}
