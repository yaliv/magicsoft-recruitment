package queue

// The Queue interface, defining the required methods.
type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

// Create new Queue object.
func New(size int) Queue {
	var q Queue = QueueImpl{make([]interface{}, size)}
	return q
}

// The implementation of Queue.
// A struct with Key field, it holds items in the queue.
type QueueImpl struct {
	Key []interface{}
}

// Add new item to the queue and evict oldest item if the slots are full.
func (q QueueImpl) Push(key interface{}) {
	// If there is nil slot, simply put the item into the slot.
	for i, k := range q.Key {
		if k == nil {
			q.Key[i] = key
			return
		}
	}

	// No nil slot! Evict oldest item then put the item in the last slot.
	lastSlot := q.Len() - 1

	for i := range q.Key {
		if i < lastSlot {
			q.Key[i] = q.Key[i+1]
		} else {
			q.Key[i] = key
		}
	}
}

// Get the oldest item to be popped and shift the queue.
func (q QueueImpl) Pop() interface{} {
	lastSlot := q.Len() - 1

	for _, k := range q.Key {
		if k != nil {
			// Get the oldest item to be popped.
			poppedItem := k

			// Shift the queue.
			for i := range q.Key {
				if i == lastSlot || q.Key[i+1] == nil {
					q.Key[i] = nil
					return poppedItem
				}

				q.Key[i] = q.Key[i+1]
			}
		}
	}

	return nil
}

// Check if the queue contains the item.
func (q QueueImpl) Contains(key interface{}) bool {
	for _, k := range q.Key {
		if key == k {
			return true
		}
	}

	return false
}

// Get the number of items in the queue.
func (q QueueImpl) Len() int {
	var filledSlots int

	for _, k := range q.Key {
		if k == nil {
			// If there is a nil slot,
			// then the remaining slots must be nil too.
			return filledSlots
		} else {
			filledSlots++
		}
	}

	return filledSlots
}

// Get all the items in the queue.
func (q QueueImpl) Keys() []interface{} {
	return q.Key
}
