// generated by metactl sdk gen 
package sdk

const (
	ClientAccountSelectName = "ClientAccountSelect"
)

type ClientAccountSelect struct {
    All *bool `json:"all,omitempty",yaml:"all,omitempty"`
    AlternativeIds *IdSelect `json:"alternativeIds,omitempty",yaml:"alternativeIds,omitempty"`
    Id *ServiceIdSelect `json:"id,omitempty",yaml:"id,omitempty"`
    Meta *TypeMetaSelect `json:"meta,omitempty",yaml:"meta,omitempty"`
    Password *PasswordSelect `json:"password,omitempty",yaml:"password,omitempty"`
    Relations *ClientAccountRelationsSelect `json:"relations,omitempty",yaml:"relations,omitempty"`
}