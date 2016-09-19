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

// TestPush validate evict old item policy
func TestEvictPolicy(t *testing.T) {
	size := 5
	q := New(size)

	for i, v := range testValues {
		q.Push(v, i)

		// validate
		// item existence
		if !q.Contains(v) {
			t.Errorf("policy: newly inserted %v must be exists", v)
		}

		if i < 5 && q.Len() != (i+1) {
			t.Errorf("expected length: %d", i+1)
		} else {
			t.Errorf("expexted length: %d", size)
		}
	}
}

// TestPop validate pop item policy
func TestPop(t *testing.T) {
	size := 5
	q := New(size)

	for i, v := range testValues {
		q.Push(v, i)
	}

	for q.Len() > 0 {
		v := q.Pop()

		// validate
		expect := testValues[len(testValues)-q.Len()-1]
		if v != expect {
			t.Error("expected %v but recevied %v", expect, v)
		}
	}

}
