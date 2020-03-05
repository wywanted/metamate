// generated by metactl sdk gen 
package sdk

const (
	ClientAccountName = "ClientAccount"
)

type ClientAccount struct {
    AlternativeIds []Id `json:"alternativeIds,omitempty",yaml:"alternativeIds,omitempty"`
    Id *ServiceId `json:"id,omitempty",yaml:"id,omitempty"`
    Meta *TypeMeta `json:"meta,omitempty",yaml:"meta,omitempty"`
    Password *Password `json:"password,omitempty",yaml:"password,omitempty"`
    Relations *ClientAccountRelations `json:"relations,omitempty",yaml:"relations,omitempty"`
}