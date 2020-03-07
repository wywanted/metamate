// generated by metactl sdk gen 
package sdk

const (
	GetTransactionsResponseFilterName = "GetTransactionsResponseFilter"
)

type GetTransactionsResponseFilter struct {
    And []GetTransactionsResponseFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Meta *CollectionMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Not []GetTransactionsResponseFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []GetTransactionsResponseFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
    Transactions *TransactionListFilter `json:"transactions,omitempty",yaml:"transactions,omitempty"`
}