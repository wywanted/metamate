// generated by metactl sdk gen 
package sdk

const (
	PutBankAccountsRequestFilterName = "PutBankAccountsRequestFilter"
)

type PutBankAccountsRequestFilter struct {
    And []PutBankAccountsRequestFilter `json:"and,omitempty" yaml:"and,omitempty"`
    BankAccounts *BankAccountListFilter `json:"bankAccounts,omitempty" yaml:"bankAccounts,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *PutModeFilter `json:"mode,omitempty" yaml:"mode,omitempty"`
    Not []PutBankAccountsRequestFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PutBankAccountsRequestFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}