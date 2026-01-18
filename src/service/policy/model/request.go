package model

type Client struct {
	Code string `json:"code"`
}

type Principal struct {
	ExternalID string `json:"external_id"`
	Type       string `json:"type"`
}

type Product struct {
	Code string `json:"code"`
}

type Scope struct {
	Key string `json:"key"`
}

type CheckChainScope struct {
	Chain     interface{} `json:"chain"`
	Requested Scope       `json:"requested"`
}

type CheckRequest struct {
	Client     Client    `json:"client"`
	Permission string    `json:"permission"`
	Principal  Principal `json:"principal"`
	Product    Product   `json:"product"`
	Scopes     Scope     `json:"scopes"`
}

type CheckChainRequest struct {
	Client     Client          `json:"client"`
	Permission string          `json:"permission"`
	Principal  Principal       `json:"principal"`
	Product    Product         `json:"product"`
	Scope      CheckChainScope `json:"scope"`
}

type CheckBulkItem struct {
	Permission string `json:"permission"`
	Scope      Scope  `json:"scope"`
}

type CheckBulkRequest struct {
	Checks    []CheckBulkItem `json:"checks"`
	Client    Client          `json:"client"`
	Principal Principal       `json:"principal"`
	Product   Product         `json:"product"`
}

type Role struct {
	Code string `json:"code"`
	ID   int    `json:"id"`
}

type Matched struct {
	Role         Role  `json:"role"`
	Scope        Scope `json:"scope"`
	AssignmentID int   `json:"assignment_id"`
}
