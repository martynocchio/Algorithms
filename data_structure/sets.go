package main

type Set struct {
	Collection []interface{}
}

func (s *Set) Contains(val interface{}) (int, bool) {
	for k, v := range s.Collection {
		if v == val {
			return k, true
		}
	}
	return 0, false
}

func NewSet() *Set {
	return &Set{
		Collection: []interface{}{},
	}
}

func (s *Set) Add(val interface{}) bool {
	if _, exists := s.Contains(val); exists == false {
		s.Collection = append(s.Collection, val)
		return true
	}
	return false
}

func (s *Set) Delete(val interface{}) bool {
	if i, exists := s.Contains(val); exists == true {
		s.Collection[i] = s.Collection[len(s.Collection)-1]
		s.Collection = s.Collection[:len(s.Collection)-1]
		return true
	}
	return false
}
