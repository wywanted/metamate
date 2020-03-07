// generated by metactl sdk gen 
package sdk

const (
	PipeStatusesRequestFilterName = "PipeStatusesRequestFilter"
)

type PipeStatusesRequestFilter struct {
    And []PipeStatusesRequestFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Context *PipeStatusesContextFilter `json:"context,omitempty",yaml:"context,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Mode *PipeModeFilter `json:"mode,omitempty",yaml:"mode,omitempty"`
    Not []PipeStatusesRequestFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PipeStatusesRequestFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}