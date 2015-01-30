package main

import (
	"fmt"
	"strings"
	"math/rand"
	"bufio"
	"os"
)

func randIP(r *rand.Rand) string {
	return fmt.Sprintf(
		"%d.%d.%d.%d", 
		r.Intn(255), 
		r.Intn(255), 
		r.Intn(255), 
		r.Intn(255),
	)
}

// returns the tree root + an index of all machines
func buildInternet() (*Machine, []*Machine) {
	
	r := rand.New(rand.NewSource(999))
	
	// builds up a fake internet of machines.

	masterList := make([]*Machine, 0, 0)
	toExpand := make([]*Machine, 0, 0)
	root := NewMachine(randIP(r))

	toExpand = append(toExpand, root)
	masterList = append(masterList, root)

	for n := 0; n < 50; n++  {

		// pick a random node, m, to expand

		if len(toExpand) == 0 {
			break
		}

		i := r.Intn(len(toExpand))-1

		if i < 0 {
			i = 0
		}

		m := toExpand[i]
		toExpand = append(toExpand[:i], toExpand[i+1:]...)

		// generate 10 machines, n, to attach to m

		for q := 0; q < 10; q++ {
			n := NewMachine(randIP(r))
			n.Attach(m)
			toExpand = append(toExpand, n)
			masterList = append(masterList, n)
		}

	}

	return root, masterList

}

func pathBetween(m1 *Machine, m2 *Machine) []*Machine {

	pathToRoot := func(m *Machine) []*Machine {

		out := make([]*Machine, 0, 0)
		out = append(out, m)
		t := m
		for t.Network != nil {
			t = t.Network
			out = append(out, t)
		}
		return out
	}

	m1path := pathToRoot(m1)
	m2path := pathToRoot(m2)

	var ancestor *Machine
	for _, m := range(m1path) {
		for _, n := range(m2path) {
			if m == n {
				ancestor = m
				break
			}
		}
	}

	outPath := make([]*Machine, 0, 0)

	for _, m := range(m1path){
		outPath = append(outPath, m)
		if m == ancestor {
			break
		}
	}

	// bug: this needs to be reversed

	for i := len(m2path)-1; i >= 0; i-- {
		m := m2path[i]
		if m == ancestor {
			break
		} else {
			outPath = append(outPath, m)
		}
	}

	return outPath

}

func main() {

	// build an internet

	root, index := buildInternet()
	_ = root
	_ = index

	find := func(idx []*Machine, ip string) *Machine {
		for _, m := range(idx) {
			if m.IP == ip {
				return m
			}
		}
		return nil
	}

	// drop into an input loop

	var input string
	var current string

	for input != "exit" {

		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		tokens := strings.Split(input, " ")

		fmt.Printf("\n")

		if tokens[0] == "exit" {

			break

		} else if tokens[0] == "ssh" {

			fmt.Printf("Now connected to %s\n", tokens[1])
			current = tokens[1]

		} else if tokens[0] == "scan" {

			fmt.Printf("Scanning %s\n", tokens[1])

			for _, m := range(index){
				if strings.HasPrefix(m.IP, tokens[1]) {
					fmt.Printf("  %s\n", m.IP)
				}
			}

		} else if tokens[0] == "trace" {

			fmt.Printf("Path between %s and %s\n", current, tokens[1])
			for _, m := range(pathBetween(find(index, current), find(index, tokens[1]))){
				fmt.Println(m.IP)
			}

		} else if tokens[0] == "ping" {

			if find(index, tokens[1]) != nil {
				fmt.Printf("%s exists\n", tokens[1])
			} else {
				fmt.Printf("%s does not\n", tokens[1])
			}

		} else if tokens[0] == "debug" {

			m := find(index, tokens[1])

			fmt.Printf("  ip: %s", m.IP)
			fmt.Printf("  network attachments: %s", m.Network)

		} else {


			fmt.Printf("%s: command not found\n", tokens[0])

		}

	}

}
