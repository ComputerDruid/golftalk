package main

import "errors"

type SexpPair struct {
	val interface{}
	next interface{}
}

var EmptyList *SexpPair = nil

func (pair *SexpPair) Len() (length int, err error) {
	if pair == EmptyList {
		return
	}
	if next, ok := pair.next.(*SexpPair); ok {
		length, err = next.Len()
		length++
		return
	}
	return 1, errors.New("pair does not represent a list")
}

func toList(items... interface{}) (head *SexpPair) {
	head = EmptyList
	for i := len(items) - 1; i >= 0; i-- {
		head = &SexpPair{items[i], head}
	}
	return
}

// Get is a simple utility function to Get the nth item from a linked list.
func Get(lst *SexpPair, n int) interface{} {
	obj := lst

	for i := 0; i < n; i++ {
		obj, _ = obj.next.(*SexpPair)
	}

	return obj.val
}

// ToSlice converts a linked list into a slice.
func ToSlice(lst *SexpPair) (result []interface{}) {
	ok := true
	for e := lst ; e != EmptyList && ok; e, ok = e.next.(*SexpPair) {
		result = append(result, e.val)
	}
	return
}
