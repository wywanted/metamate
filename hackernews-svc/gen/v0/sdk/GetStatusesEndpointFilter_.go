// generated by metactl sdk gen 
package sdk

const (
	GetStatusesEndpointFilterName = "GetStatusesEndpointFilter"
)

type GetStatusesEndpointFilter struct {
    And []GetStatusesEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []GetStatusesEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetStatusesEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}