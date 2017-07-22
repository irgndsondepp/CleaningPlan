package interfaces

type Converter interface {
	ConvertTo(Plan) ([]byte, error)
	ReadFrom([]byte, Plan) error
}
