package mock

import (
	"github.com/irgndsondepp/cleaningplan/interfaces"
)

type Persistence struct {
	plan interfaces.Plan
}

func NewMockPersistence(p interfaces.Plan) *Persistence {
	return &Persistence{
		plan: p,
	}
}

func (p *Persistence) Save(plan interfaces.Plan) {
	p.plan = plan
}

func (p *Persistence) Load(plan interfaces.Plan) {
	plan = p.plan
}
