package utils

type ansiTagsIterator struct {
	tagIndex [][]int
	tagData  [][]string
	pos      int
	valid    bool
}

// Create stateful generator
func NewANSITagsIterator(index [][]int, tags [][]string) *ansiTagsIterator {
	ref := new(ansiTagsIterator)
	ref.tagIndex = index
	ref.tagData = tags
	ref.pos = -1
	ref.valid = len(ref.tagIndex) > 0

	if len(ref.tagIndex) != len(ref.tagData) {
		panic("Length of tags data and its index mismatch")
	}

	return ref
}

// Previous is returning either first value or previous before Next()
func (ref *ansiTagsIterator) Previous() (int, string) {
	var idx int
	if ref.pos == 0 {
		idx = 0
	} else {
		idx = ref.pos - 1
	}

	return ref.tagIndex[idx][0], ref.tagData[idx][0]
}

// Value is returning the value of the element in the iterator
func (ref *ansiTagsIterator) Value() (int, string) {
	return ref.tagIndex[ref.pos][0], ref.tagData[ref.pos][0]
}

// Resets the iterator
func (ref *ansiTagsIterator) Reset() *ansiTagsIterator {
	ref.pos = -1
	return ref
}

// HasNext returns a flag if the iterator has next element
func (ref *ansiTagsIterator) HasNext() bool {
	if !ref.valid {
		return false
	}

	return !(len(ref.tagIndex) < ref.pos)
}

// Next returns next element in the iterator
func (ref *ansiTagsIterator) Next() bool {
	if !ref.valid {
		return false
	}
	ref.pos++
	return len(ref.tagIndex) > ref.pos
}
