// generated by metactl sdk gen 
package mql

const (
	IndexPageFilterName = "IndexPageFilter"
)

type IndexPageFilter struct {
    And []IndexPageFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []IndexPageFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []IndexPageFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Value *Int32Filter `json:"value,omitempty" yaml:"value,omitempty"`
}