package model

type CheckResponse struct {
	Allowed    bool                  `json:"allowed"`
	Decision   string                `json:"decision"`
	ReasonCode string                `json:"reason_code"`
	Matched    *CheckMatchedResponse `json:"matched,omitempty"`
}

type CheckChainEvaluationResponse struct {
	MatchedViaScope string `json:"matched_via_scope"`
	RequestedScope  string `json:"requested_scope"`
}

type CheckChainResponse struct {
	Allowed    bool                         `json:"allowed"`
	Decision   string                       `json:"decision"`
	Evaluation CheckChainEvaluationResponse `json:"evaluation"`
	Matched    *CheckMatchedResponse        `json:"matched,omitempty"`
	ReasonCode string                       `json:"reason_code"`
}

type CheckBulkItemResponse struct {
	Allowed    bool                  `json:"allowed"`
	Decision   string                `json:"decision"`
	Matched    *CheckMatchedResponse `json:"matched,omitempty"`
	ReasonCode string                `json:"reason_code"`
}

type CheckBulkResponse struct {
	Checks []CheckBulkItemResponse `json:"checks"`
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
