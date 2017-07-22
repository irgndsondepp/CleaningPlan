package impl

import (
	"encoding/json"

	"github.com/irgndsondepp/cleaningplan/interfaces"
)

type JSONConverter struct {
}

func NewJSONConverter() *JSONConverter {
	return &JSONConverter{}
}

func (*JSONConverter) ConvertTo(plan interfaces.Plan) ([]byte, error) {
	return json.MarshalIndent(plan, "", "\t")
}

func (*JSONConverter) ReadFrom(bytes []byte, cp interfaces.Plan) error {
	err := json.Unmarshal(bytes, &cp)
	return err
}
