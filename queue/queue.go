package queue

type Queue interface {
	Push(key interface{}, item interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

func New(size int) Queue {
	// change constructor
	return nil
}
