package policy

import (
	"github.com/godispatcher/dispatcher/coordinator"
	"github.com/godispatcher/dispatcher/model"
	m2 "sdk/src/service/policy/model"
)

func New(licence string) Policy {
	return Policy{Licence: licence}
}

type Policy struct {
	Address string
	Licence string
}

func (s Policy) Check(req m2.CheckRequest) (*m2.CheckResponse, error) {
	form := model.DocumentForm{}
	err := form.FromInterface(req)
	if err != nil {
		return nil, err
	}
	requester := coordinator.ServiceRequest[m2.CheckRequest, m2.CheckResponse]{
		Address: s.Address,
		Document: model.Document{
			Department:  DEPARTMENT,
			Transaction: "check",
			Security:    &model.Security{Licence: s.Licence},
		},
		Request: req,
	}

	response, err := requester.CallTransaction()

	return &response, err
}
