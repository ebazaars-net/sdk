package policy

import (
	"os"
	m2 "sdk/src/service/policy/model"

	"github.com/godispatcher/dispatcher/coordinator"
	"github.com/godispatcher/dispatcher/model"
)

func New(licence string) Policy {
	return Policy{Licence: licence}
}

type Policy struct {
	Address string
	Licence string
}

func (s Policy) GetAddress() string {
	if s.Address == "" && os.Getenv("POLICY_SERVICE_ADDRESS") != "" {
		return os.Getenv("POLICY_SERVICE_ADDRESS")
	}

	return s.Address
}

func (s Policy) Check(req m2.CheckRequest) (*m2.CheckResponse, error) {
	form := model.DocumentForm{}
	err := form.FromInterface(req)
	if err != nil {
		return nil, err
	}
	requester := coordinator.ServiceRequest[m2.CheckRequest, m2.CheckResponse]{
		Address: s.GetAddress(),
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

func (s Policy) CheckChain(req m2.CheckChainRequest) (*m2.CheckChainResponse, error) {
	form := model.DocumentForm{}
	err := form.FromInterface(req)
	if err != nil {
		return nil, err
	}
	requester := coordinator.ServiceRequest[m2.CheckChainRequest, m2.CheckChainResponse]{
		Address: s.GetAddress(),
		Document: model.Document{
			Department:  DEPARTMENT,
			Transaction: "checkChain",
			Security:    &model.Security{Licence: s.Licence},
		},
		Request: req,
	}

	response, err := requester.CallTransaction()

	return &response, err
}

func (s Policy) CheckBulk(req m2.CheckBulkRequest) (*m2.CheckBulkResponse, error) {
	form := model.DocumentForm{}
	err := form.FromInterface(req)
	if err != nil {
		return nil, err
	}
	requester := coordinator.ServiceRequest[m2.CheckBulkRequest, m2.CheckBulkResponse]{
		Address: s.GetAddress(),
		Document: model.Document{
			Department:  DEPARTMENT,
			Transaction: "checkBulk",
			Security:    &model.Security{Licence: s.Licence},
		},
		Request: req,
	}

	response, err := requester.CallTransaction()

	return &response, err
}
