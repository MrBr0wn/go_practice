package person

type Person struct {
	name     string
	homemork string
	children []*Person
}

func (p Person) DoHomework() string {
	return p.homemork
}

func (p Person) Children() []*Person {
	return p.children
}

func (p Person) Work(tasks []string) string {
	s := p.name + " work:"
	for _, task := range tasks {
		s += "\n I do " + task
	}
	return s
}

func (p Person) String() string {
	return p.name
}
