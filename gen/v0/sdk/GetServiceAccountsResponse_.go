// generated by metactl sdk gen 
package sdk

const (
	GetServiceAccountsResponseName = "GetServiceAccountsResponse"
)

type GetServiceAccountsResponse struct {
    Meta *CollectionMeta `json:"meta,omitempty",yaml:"meta,omitempty"`
    ServiceAccounts []ServiceAccount `json:"serviceAccounts,omitempty",yaml:"serviceAccounts,omitempty"`
}