// generated by metactl sdk gen 
package mql

const (
	GetBlueDummiesResponseName = "GetBlueDummiesResponse"
)

type GetBlueDummiesResponse struct {
    BlueDummies []BlueDummy `json:"blueDummies,omitempty" yaml:"blueDummies,omitempty"`
    Count *int32 `json:"count,omitempty" yaml:"count,omitempty"`
    Errors []Error `json:"errors,omitempty" yaml:"errors,omitempty"`
    Pagination *Pagination `json:"pagination,omitempty" yaml:"pagination,omitempty"`
    Warnings []Warning `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}