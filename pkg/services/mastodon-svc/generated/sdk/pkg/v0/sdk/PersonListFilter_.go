// generated by metactl sdk gen 
package sdk

const (
	PersonListFilterName = "PersonListFilter"
)

type PersonListFilter struct {
    Every *PersonFilter `json:"every,omitempty",yaml:"every,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    None *PersonFilter `json:"none,omitempty",yaml:"none,omitempty"`
    Some *PersonFilter `json:"some,omitempty",yaml:"some,omitempty"`
}