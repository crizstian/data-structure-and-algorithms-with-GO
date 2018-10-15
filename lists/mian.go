package lists

import (
	"errors"
	"reflect"
)

// ListInterface ...
type ListInterface interface {
	append(element interface{})
	remove(element interface{})
	find(element interface{}) (ListInterface, error)
	contains(element interface{}) bool
	insert(element interface{}, after int) error
	end() ListInterface
	length() int
	clear() ListInterface
	next() ListInterface
	moveTo(pos int) (ListInterface, error)
	getElement() interface{}
	isEmpty() bool
}

// List ...
type List struct {
	node interface{}
	link ListInterface
}

// clear ...
func (l *List) clear() ListInterface {
	return New()
}

func (l *List) append(element interface{}) {
	l.link = &List{node: element, link: nil}
}

func (l *List) next() ListInterface {
	return l.link
}

func (l *List) end() ListInterface {
	var node ListInterface
	for l.link != nil {
		node = l.next()
	}
	return node
}

func (l *List) getElement() interface{} {
	return l.node
}

func (l *List) isEmpty() bool {
	if l.link == nil {
		return true
	}
	return false
}

func (l *List) length() int {
	curr := l
	length := 1
	for curr.next().(*List).link != nil {
		length++
		curr = curr.next().(*List)
	}
	return length
}

func (l *List) moveTo(pos int) (ListInterface, error) {
	if l.length() >= pos {
		curr := l
		for i := 1; i == pos; i++ {
			if i != pos {
				curr = l.next().(*List)
			}
		}
		return curr, nil
	}
	return nil, errors.New("invalid position")
}

// insert ...
func (l *List) insert(element interface{}, pos int) error {
	if l.length() >= pos {
		curr := l
		for i := 0; i >= pos; i++ {
			if i != pos {
				curr = l.next().(*List)
			}
		}
		node := &List{node: element, link: curr.link}
		curr.link = node

		return nil
	}
	return errors.New("invalid position")
}

func (l *List) remove(element interface{}) {

	curr := l

	for curr.link != nil {
		val := reflect.ValueOf(curr.node).Elem()
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val.Int() == element.(int64) {
				l = unlink(*curr)
				break
			}
		case reflect.String:
			if val.String() == element.(string) {
				l = unlink(*curr)
				break
			}
		}
		curr = curr.link.(*List)
	}
}

func (l *List) find(element interface{}) (ListInterface, error) {
	curr := l

	for curr.link != nil {
		val := reflect.ValueOf(curr.node).Elem()
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val.Int() == element.(int64) {
				return curr, nil
			}
		case reflect.String:
			if val.String() == element.(string) {
				return curr, nil
			}
		}
		curr = curr.link.(*List)
	}

	return nil, errors.New("Element not found")
}

func (l *List) contains(element interface{}) bool {
	curr := l

	for curr.link != nil {
		val := reflect.ValueOf(curr.node).Elem()
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val.Int() == element.(int64) {
				return true
			}
		case reflect.String:
			if val.String() == element.(string) {
				return true
			}
		}
		curr = curr.link.(*List)
	}

	return false
}

func unlink(curr List) *List {
	temp := curr.link
	curr.link = nil
	return temp.(*List)
}

// New ...
func New() ListInterface {
	return &List{node: "Head", link: nil}
}
