package utils

//IDStore is a struct for keeping track of IDs delegated
type IDStore struct {
	upperBound uint64
	freeSpaces []uint64
}

//Fetch returns a new ID from the store
func (s *IDStore) Fetch() uint64 {
	var index uint64
	if len(s.freeSpaces) > 0 {
		index = s.freeSpaces[0]
		s.freeSpaces = s.freeSpaces[1:]
	} else {
		s.upperBound += 1
		index = s.upperBound
	}
	return index
}

//Release frees up an ID for the store
func (s *IDStore) Release(id uint64) {
	s.freeSpaces = append(s.freeSpaces, id)
}
