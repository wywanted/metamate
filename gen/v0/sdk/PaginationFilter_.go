// generated by metactl sdk gen 
package sdk

const (
	PaginationFilterName = "PaginationFilter"
)

type PaginationFilter struct {
    And []PaginationFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Current *PageFilter `json:"current,omitempty",yaml:"current,omitempty"`
    Next *PageFilter `json:"next,omitempty",yaml:"next,omitempty"`
    Not []PaginationFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PaginationFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Previous *PageFilter `json:"previous,omitempty",yaml:"previous,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}