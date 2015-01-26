package main

import (
	"fmt"
	"math/rand"
)

const (
	r rand.Rand = rand.New(rand.NewSource(999))
)

func randIP() string {
	return fmt.Sprintf(
		"%d.%d.%d.%d", 
		r.Intn(255), 
		r.Intn(255), 
		r.Intn(255), 
		r.Intn(255),
	)
}

func buildInternet() {

	// builds up a fake internet of machines.

	toExpand := make([]*Machine, 0, 0)
	root := NewMachine(randIP())
	toExpand = append(toExpand, root)

	for _ := range(50) {

		// pick a random node to expand

		i := r.Intn(len(toExpand)-1)

		if i < 0 {
			i = 0
		}

		m := toExpand(i)
		toExpand = append(toExpand[:i], toExpand[i+1:]...)

		for _ := range(10){
			n := NewMachine(randInt())
			toExpand = append(toExpand, n)
		}

	}



}

func main() {

}
