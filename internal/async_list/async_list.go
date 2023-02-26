package asynclist

import (
	"container/list"
	"sync"
)

type AsyncList interface {
	PushBack(in interface{})
	Front() *list.Element
	Back() *list.Element
	Next(in *list.Element) *list.Element
	Prev(in *list.Element) *list.Element
	GetValue(in *list.Element) interface{}
}

type asyncList struct {
	list *list.List
	mu   sync.Mutex
}

func New() *asyncList {
	return &asyncList{
		list: list.New(),
	}
}

func (l *asyncList) PushBack(in interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.PushBack(in)
}

func (l *asyncList) Front() *list.Element {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Front()
}

func (l *asyncList) Back() *list.Element {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Back()
}

func (l *asyncList) Next(in *list.Element) *list.Element {
	l.mu.Lock()
	defer l.mu.Unlock()
	return in.Next()
}

func (l *asyncList) Prev(in *list.Element) *list.Element {
	l.mu.Lock()
	defer l.mu.Unlock()
	return in.Prev()
}

func (l *asyncList) GetValue(in *list.Element) interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()
	return in.Value
}
