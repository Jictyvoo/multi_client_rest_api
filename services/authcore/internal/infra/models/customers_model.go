package models

type (
	CustomersModel struct {
		Name      string `json:"name"`
		Key       []byte `json:"key"`
		Namespace string `json:"namespace"`
	}
	CustomersList struct {
		Customers []*CustomersModel `json:"customers"`
	}
)
