// generated by metactl sdk gen 
package mql

const (
	SearchGetModeName = "SearchGetMode"
)

type SearchGetMode struct {
    Location *LocationQuery `json:"location,omitempty" yaml:"location,omitempty"`
    Term *string `json:"term,omitempty" yaml:"term,omitempty"`
}