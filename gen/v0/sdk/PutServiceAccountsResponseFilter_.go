// generated by metactl sdk gen 
package sdk

const (
	PutServiceAccountsResponseFilterName = "PutServiceAccountsResponseFilter"
)

type PutServiceAccountsResponseFilter struct {
    And []PutServiceAccountsResponseFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Not []PutServiceAccountsResponseFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PutServiceAccountsResponseFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}