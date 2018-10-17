package lists

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// ListInterface ...
type ListInterface interface {
	append(element interface{})
	remove(element interface{}) error
	// find(element interface{}) (ListInterface, error)
	// contains(element interface{}) bool
	// insert(element interface{}, after int) error
	// end() ListInterface
	// length() int
	// clear() ListInterface
	// next() ListInterface
	// moveTo(pos int) (ListInterface, error)
	// getElement() interface{}
	// isEmpty() bool
	display()
}

// Node ...
type Node struct {
	next  *Node
	value interface{}
}

// List ...
type List struct {
	root *Node
	len  int
}

// append ...
func (l *List) append(element interface{}) {
	node := &Node{value: element, next: nil}
	if l.root.value == nil {
		l.root = node
	} else {
		last := l.root
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = node
	}
	l.len++
}

func (l *List) remove(element interface{}) error {
	last := l.root
	prev := l.root
	for {
		if validateElement(last.value, element) {
			if prev == last {
				l.root = last.next
			} else {
				prev.next = last.next
			}
			l.len--
			return nil
		}
		if last.next == nil {
			break
		}
		prev = last
		last = last.next
	}
	return errors.New("No element in the list")
}

func (l *List) display() {
	d := ""
	if l.root.value != nil {
		last := l.root
		for {
			if last.next != nil {
				d += fmt.Sprintf("%v, ", last.value)
			}
			if last.next == nil {
				d += fmt.Sprintf("%v", last.value)
				break
			}
			last = last.next
		}
	}
	log.Println("\n", d)
}

// New ...
func New() *List {
	return &List{
		root: &Node{
			value: nil,
			next:  nil,
		},
		len: 0,
	}
}

func validateElement(node, element interface{}) bool {
	val := reflect.TypeOf(node)
	valE := reflect.TypeOf(element)
	if val != valE {
		return false
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if node.(int) == element.(int) {
			return true
		}
	case reflect.String:
		if node.(string) == element.(string) {
			return true
		}
		// etc...
	}
	return false
}
