// generated by metactl sdk gen 
package sdk

const (
	PutServicesResponseFilterName = "PutServicesResponseFilter"
)

type PutServicesResponseFilter struct {
    And []PutServicesResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []PutServicesResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PutServicesResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}