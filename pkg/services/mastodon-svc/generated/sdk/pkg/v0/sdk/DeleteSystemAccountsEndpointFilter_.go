// generated by metactl sdk gen 
package sdk

const (
	DeleteClientAccountsEndpointFilterName = "DeleteClientAccountsEndpointFilter"
)

type DeleteClientAccountsEndpointFilter struct {
    And []DeleteClientAccountsEndpointFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Not []DeleteClientAccountsEndpointFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []DeleteClientAccountsEndpointFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}