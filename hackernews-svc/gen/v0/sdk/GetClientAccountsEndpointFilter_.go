// generated by metactl sdk gen 
package sdk

const (
	GetClientAccountsEndpointFilterName = "GetClientAccountsEndpointFilter"
)

type GetClientAccountsEndpointFilter struct {
    And []GetClientAccountsEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []GetClientAccountsEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetClientAccountsEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}