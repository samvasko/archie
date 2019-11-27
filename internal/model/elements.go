package model

import (
	"fmt"
)

const (
	actor = iota
	item
)

type element struct {
	name string
	tags []string
	kind uint
}

type Element interface {
	Name() string
	Tags() []string
	ID() string
	IsActor() bool
}

// Create a new item
func NewItem(name string, tags []string) Element {
	el := newElement(item)
	el.name = name
	el.tags = tags
	return &el
}

// Create a new actor
func NewActor(name string) Element {
	el := newElement(actor)
	el.name = name
	return &el
}

// Create an element
func newElement(kind uint) element {
	return element{
		kind: kind,
	}
}

func (e *element) ID() string {
	return fmt.Sprintf("%p", e)
}

func (e *element) IsActor() bool {
	return e.kind == actor
}

func (e *element) Name() string {
	return e.name
}

func (e *element) Tags() []string {
	return e.tags
}
