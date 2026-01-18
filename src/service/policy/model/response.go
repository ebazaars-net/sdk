package model

type CheckResponse struct {
	Allowed    bool                  `json:"allowed"`
	Decision   string                `json:"decision"`
	ReasonCode string                `json:"reason_code"`
	Matched    *CheckMatchedResponse `json:"matched,omitempty"`
}

type CheckMatchedResponse struct {
	Role         CheckMatchedRoleResponse  `json:"role"`
	Scope        CheckMatchedScopeResponse `json:"scope"`
	AssignmentId interface{}               `json:"assignment_id"`
}

type CheckMatchedRoleResponse struct {
	Code string      `json:"code"`
	Id   interface{} `json:"id"`
}

type CheckMatchedScopeResponse struct {
	Key string `json:"key"`
}
