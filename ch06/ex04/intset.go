package main

import (
	"bytes"
	"fmt"
	"math/bits"
)

type IntSet struct {
	words []uint64
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // {1 9 144}
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // {9 42}
	x.UnionWith(&y)
	fmt.Println(x.String()) // {1 9 42 144}
	x.DifferenceWith(&y)
	fmt.Println(x.String()) // {1 144}
	x.Add(9)
	x.IntersectWith(&y)
	fmt.Println(x.String()) // {9}
	x.Add(1)
	x.Add(42)
	y.Add(3)
	x.SymmetricDifference(&y)
	fmt.Println(x.String()) // {1 3}
	fmt.Println(x.Has(1), x.Has(123)) // true false
	fmt.Println(x.String(), x.Len()) // {1 3} 2
	x.Remove(1)
	fmt.Println(x.String()) // {3}
	x.Clear()
	fmt.Println(x.String()) // {}
	c1 := x.Copy()
	c2 := x.Copy()
	fmt.Println(&c1) // 0x1400000e030
	fmt.Println(&c2) // 0x1400000e038
	x.AddAll(2, 3, 4)
	fmt.Println(x.String()) // {2 3 4}
	e := x.Elems()
	fmt.Println(e) // [2 3 4]
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		len += bits.OnesCount64(word)
	}
	return len
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word < len(s.words) {
		break
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	var c IntSet
	c.words = make([]uint64, len(s.words))
	copy(c.words, s.words)
	return &c
}

func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			break
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elems() []int {
	if s.Len() == 0 {
		return []int{}
	}
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for bit := 0; bit < 64; bit++ {
			if word&(1<<uint(bit)) != 0 {
				elems = append(elems, 64*i+bit)
			}
		}
	}
	return elems
}