// generated by metactl sdk gen 
package sdk

const (
	GetServiceAccountsRequestFilterName = "GetServiceAccountsRequestFilter"
)

type GetServiceAccountsRequestFilter struct {
    And []GetServiceAccountsRequestFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Meta *RequestMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Mode *GetModeFilter `json:"mode,omitempty",yaml:"mode,omitempty"`
    Not []GetServiceAccountsRequestFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []GetServiceAccountsRequestFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Pages *ServicePageListFilter `json:"pages,omitempty",yaml:"pages,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}