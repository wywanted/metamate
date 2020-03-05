// generated by metactl sdk gen 
package sdk

const (
	PostStatusesRequestFilterName = "PostStatusesRequestFilter"
)

type PostStatusesRequestFilter struct {
    And []PostStatusesRequestFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Meta *RequestMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Not []PostStatusesRequestFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PostStatusesRequestFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
    Statuses *StatusListFilter `json:"statuses,omitempty",yaml:"statuses,omitempty"`
}