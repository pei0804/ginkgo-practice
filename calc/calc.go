package calc

import "fmt"

type Hoge struct {
	Atai1 int
	Atai2 int
}

func (h Hoge) Add() int {
	return h.Atai1 + h.Atai2
}

func (h Hoge) Sub() int {
	return h.Atai1 - h.Atai2
}

func NewHoge(atai1 int, atai2 int) (Hoge, error) {
	if atai1 == 0 || atai2 == 0 {
		return Hoge{}, fmt.Errorf("%s", "0はあかんよ")
	}
	return Hoge{Atai1: atai1, Atai2: atai2}, nil
}
