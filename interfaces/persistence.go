package interfaces

type Persistence interface {
	Save(Plan)
	Load(Plan)
}
