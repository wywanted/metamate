// generated by metactl sdk gen 
package sdk

const (
	PutStatusesRequestFilterName = "PutStatusesRequestFilter"
)

type PutStatusesRequestFilter struct {
    And []PutStatusesRequestFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *PutModeFilter `json:"mode,omitempty" yaml:"mode,omitempty"`
    Not []PutStatusesRequestFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PutStatusesRequestFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Statuses *StatusListFilter `json:"statuses,omitempty" yaml:"statuses,omitempty"`
}