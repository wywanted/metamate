// generated by metactl sdk gen 
package sdk

const (
	ClientAccountsCollectionName = "ClientAccountsCollection"
)

type ClientAccountsCollection struct {
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Meta *CollectionMeta `json:"meta,omitempty",yaml:"meta,omitempty"`
    ClientAccounts []ClientAccount `json:"clientAccounts,omitempty",yaml:"clientAccounts,omitempty"`
}