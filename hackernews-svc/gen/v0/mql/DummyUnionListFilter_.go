// generated by metactl sdk gen 
package mql

const (
	DummyUnionListFilterName = "DummyUnionListFilter"
)

type DummyUnionListFilter struct {
    Every *DummyUnionFilter `json:"every,omitempty" yaml:"every,omitempty"`
    None *DummyUnionFilter `json:"none,omitempty" yaml:"none,omitempty"`
    Some *DummyUnionFilter `json:"some,omitempty" yaml:"some,omitempty"`
}