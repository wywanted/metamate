// generated by metactl sdk gen 
package sdk

const (
	DeleteStatusesResponseFilterName = "DeleteStatusesResponseFilter"
)

type DeleteStatusesResponseFilter struct {
    And []DeleteStatusesResponseFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Not []DeleteStatusesResponseFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []DeleteStatusesResponseFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}