// generated by metactl sdk gen 
package sdk

const (
	GetStatusesEndpointName = "GetStatusesEndpoint"
)

type GetStatusesEndpoint struct {
    Filter *GetStatusesRequestFilter `json:"filter,omitempty",yaml:"filter,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
}