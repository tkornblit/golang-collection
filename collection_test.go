package collection_test

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"github.com/tkornblit/collection"
	"testing"
)

func createSliceOfNElements(n int) (slice []interface{}) {

	for {
		n--
		u, _ := uuid.NewV4()
		slice = append(slice, u)
		if n == 0 {
			break
		}
	}
	return
}

func Test_Collection_Count(t *testing.T) {
	c := collection.NewCollectionFromSlice(createSliceOfNElements(2310))
	total := len(c)
	n := 0

	chunks, _ := c.Chunks(100)

	for chunk := range chunks {
		fmt.Println(chunk)
		for _, el := range chunk {
			n++
			fmt.Println(n, el)

		}
	}

	assert.Equal(t, total, n, "Should be equal")

}
