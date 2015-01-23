package main

type Machine struct {
	IP string
	Hostname string
	Root *File
	Processes []Process
	Accounts map[string]Account
	Network *Machine
}

type Account struct {
	Name string
	Pass string
}

type Process struct {
	Name string
}

type File struct {
	Name string
	Directory bool
	Owner string
	Children []File
}

func (this *Machine) AddAccount(a Account) {
	this.Accounts[a.Name] = a
}

func (this *Machine) Attach(remote *Machine) {
	this.Network = remote
}

func NewMachine(ip string) (*Machine) {
	machine := Machine{}
	machine.IP = ip
	machine.Hostname = "salmon.fillet.net"
	machine.Accounts = make(map[string]Account)
	machine.AddAccount(Account{Name:"root", Pass:"password"})
	machine.Root = &File{"/", true, "root", make([]File, 0, 0)}
	return &machine
}

func AreNetworked(m1 *Machine, m2 *Machine) bool {

	if m1 == m2 {
		return true
	}

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

	return ancestor != nil

}


