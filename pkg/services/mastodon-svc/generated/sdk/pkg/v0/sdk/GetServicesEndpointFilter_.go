// generated by metactl sdk gen 
package sdk

const (
	GetServicesEndpointFilterName = "GetServicesEndpointFilter"
)

type GetServicesEndpointFilter struct {
    And []GetServicesEndpointFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Not []GetServicesEndpointFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []GetServicesEndpointFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}