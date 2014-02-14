package collection

import (
	"bytes"
	// "fmt"
	set "github.com/deckarep/golang-set"
	"math"
	"strconv"
)

type Collection []interface{}

type Chunk struct {
	Collection Collection
	Done       bool
}

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

func (collection Collection) Chunks(count int64) (c chan Collection, numChunks int) {
	c = make(chan Collection)

	total := int64(len(collection))

	numChunks = int(math.Ceil(float64(total) / float64(count)))

	go func() {

		batch := NewCollection()

		var i int64

		var remaining int64 = total

		for _, item := range collection {
			remaining--

			i++
			if i%count == 0 || remaining == 0 {

				c <- batch
				batch = nil

			}

			batch = append(batch, item)

		}

		close(c)

	}()

	return
}

func (collection Collection) ToCsv() string {
	var buffer bytes.Buffer

	for _, i := range collection {
		buffer.WriteString(strconv.FormatInt(i.(int64), 10))
		buffer.WriteString(",")
	}

	return buffer.String()
}

func (collection Collection) ToInt64() (array []int64) {
	for _, i := range collection {

		array = append(array, i.(int64))
	}

	return
}
