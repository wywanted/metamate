// generated by metactl sdk gen 
package sdk

const (
	WhateverUnionListFilterName = "WhateverUnionListFilter"
)

type WhateverUnionListFilter struct {
    Every *WhateverUnionFilter `json:"every,omitempty" yaml:"every,omitempty"`
    None *WhateverUnionFilter `json:"none,omitempty" yaml:"none,omitempty"`
    Some *WhateverUnionFilter `json:"some,omitempty" yaml:"some,omitempty"`
}