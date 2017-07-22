package impl

type SimplePerson struct {
	Name string
}

func NewSimplePerson(name string) *SimplePerson {
	return &SimplePerson{Name: name}
}

func (s *SimplePerson) GetName() string {
	return s.Name
}
