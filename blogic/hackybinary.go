package blogic

import (
	"math/rand"
	"time"
)

//Since it's a real pain to work with bits on a granular level,
//we're just going to use a slice of ints and cap them to 1 and
//zero
type Bnum struct {
	Number []int
}

//NewBnum creates a new Bnum struct and randomly fills it
func NewBnum() *Bnum {
	rand.Seed(time.Now().UnixNano())
	l := 8
	n := make([]int, l)
	for i := range n {
		n[i] = rand.Intn(2)
	}
	return &Bnum{
		Number: n,
	}
}

//NewBnumValue creates a new Bnum based on a given value, rather
//than filling in with random values
func NewBnumValue(n []int) *Bnum {
	return &Bnum{
		Number: n,
	}
}

//Add takes 2 Bnums and sums them into a new slice. Er, number
//Let's be honest, I *could* do the work to use byte slices or
//integers instead of slices of integers capped to 0 or 1, but
//the point is not working with binary in go, it's simulating
//a simple logic gate adder
func Add(a, b *Bnum) *Bnum {
	//Now the fun part!

	res := make([]int, 9)
	var num int
	var carry int

	for i := 7; i >= 0; i-- {
		num = Xor(a.Number[i], b.Number[i])
		//fmt.Printf("i: %v, Num: %v, Carry: %v, Xor: %v\n", i, num, carry, Xor(num, carry))
		res[i+1] = Xor(num, carry)
		//fmt.Println(res)
		carry = Or(And(a.Number[i], b.Number[i]), And(Or(a.Number[i], b.Number[i]), carry))
	}

	res[0] = carry

	return NewBnumValue(res)

}
