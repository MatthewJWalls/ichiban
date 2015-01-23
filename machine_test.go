package main

import (
	"testing"
)

func TestMachine(t *testing.T) {

    m := NewMachine("192.168.1.1")

	if m.IP != "192.168.1.1" {
		t.Error("IP was not correct.")
	}

	if m.Accounts["root"].Name != "root" {
		t.Error("Machine did not have a root account.")
	}

	if m.Root.Name != "/" {
		t.Error("Machine's root directory was not /")
	}

}

func TestNetworkOfMachines(t *testing.T) {

	m1 := NewMachine("1.1.1.1")
	m2 := NewMachine("1.1.1.2")
	m3 := NewMachine("1.1.1.3")
	m4 := NewMachine("1.1.1.4")
	m5 := NewMachine("1.1.1.5")

	m1.Attach(m3)
	m2.Attach(m3)
	m3.Attach(m4)

	if m1.Network.Network != m4 {
		t.Error("Incorrect network route")
	}

	if ! AreNetworked(m1, m1) {
		t.Error("m1 cannot see itself")
	}

	if ! AreNetworked(m1, m2) {
		t.Error("m1 cannot see m2")
	}

	if ! AreNetworked(m1, m3) {
		t.Error("m1 cannot see m3")
	}

	if ! AreNetworked(m2, m4) {
		t.Error("m2 cannot see m4")
	}

	if AreNetworked(m1, m5) {
		t.Error("m1 can see m5")
	}

}
