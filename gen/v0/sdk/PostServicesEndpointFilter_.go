// generated by metactl sdk gen 
package sdk

const (
	PostServicesEndpointFilterName = "PostServicesEndpointFilter"
)

type PostServicesEndpointFilter struct {
    And []PostServicesEndpointFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Not []PostServicesEndpointFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PostServicesEndpointFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}