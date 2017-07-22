package impl

type Flatmate struct {
	Name string `json:"Name"`
}

func NewFlatmate(name string) *Flatmate {
	fm := Flatmate{
		Name: name,
	}
	return &fm
}

func (f *Flatmate) GetName() string {
	return f.Name
}
