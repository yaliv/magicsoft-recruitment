package queue

import (
	"testing"
)

var testValues = []interface{}{
	"lorem",
	"ipsum",
	1,
	2,
	3,
	"jack",
	"jill",
	"felix",
	"donking",
}

func TestAdd(t *testing.T) {
	size := 5
	lru := New(size)

	for i, v := range testValues {
		// validate
		// item existence
		if !lru.Contains(v) {
			t.Errorf("newly inserted %v must be exists", v)
		}

		if i < 5 && lru.Len() != (i+1) {
			t.Errorf("expected length: %d", i+1)
		} else {
			t.Errorf("expexted length: %d", size)
		}
	}

}
