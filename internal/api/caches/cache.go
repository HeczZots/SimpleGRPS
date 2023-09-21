package caches

import "time"

type Buffer struct {
	Arr      Array
	Capacity int
}
type Item struct {
	Value int64
	TS    time.Time
}
type Array []Item

func NewBuffer(c int) *Buffer {
	arr := make(Array, 0, c)
	return &Buffer{Capacity: c, Arr: arr}
}
func (b *Buffer) Insert(data int64, t time.Time) {
	b.Arr = append(b.Arr, Item{Value: data, TS: t})
}
