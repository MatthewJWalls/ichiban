package main

import (
	"testing"
)

func TestPathing(t *testing.T){

	m1 := NewMachine("1.1.1.1")
	m2 := NewMachine("1.1.1.2")
	m3 := NewMachine("1.1.1.3")
	m4 := NewMachine("1.1.1.4")

	m1.Attach(m3)
	m2.Attach(m3)
	m3.Attach(m4)

	rootA := PathToRoot(m1)
	rootB := PathToRoot(m2)
	pathA := PathBetween(m1, m2)
	pathB := PathBetween(m2, m1)

	if len(rootA) != 3 {
		t.Errorf("Path was incorrect.")
	}

	if len(rootB) != 3 {
		t.Errorf("Path was incorrect.")
	}

	if len(rootA) != 3 {
		t.Errorf("Path was incorrect.")
	}

	if len(pathA) != 3 {
		t.Errorf("Path was incorrect.")
	}

	if len(pathB) != 3 {
		t.Errorf("Path was incorrect.")
	}

}


