// generated by metactl sdk gen 
package sdk

const (
	PostTransactionsEndpointFilterName = "PostTransactionsEndpointFilter"
)

type PostTransactionsEndpointFilter struct {
    And []PostTransactionsEndpointFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Not []PostTransactionsEndpointFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PostTransactionsEndpointFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}