// generated by metactl sdk gen 
package sdk

const (
	PutClientAccountsEndpointName = "PutClientAccountsEndpoint"
)

type PutClientAccountsEndpoint struct {
    Filter *PutClientAccountsRequestFilter `json:"filter,omitempty",yaml:"filter,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
}