// generated by metactl sdk gen 
package sdk

const (
	GetBankAccountsEndpointFilterName = "GetBankAccountsEndpointFilter"
)

type GetBankAccountsEndpointFilter struct {
    And []GetBankAccountsEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []GetBankAccountsEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetBankAccountsEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}