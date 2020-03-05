// generated by metactl sdk gen 
package sdk

const (
	ServiceFilterName = "ServiceFilter"
)

type ServiceFilter struct {
    AlternativeIds *IdListFilter `json:"alternativeIds,omitempty",yaml:"alternativeIds,omitempty"`
    And []ServiceFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Endpoints *EndpointsFilter `json:"endpoints,omitempty",yaml:"endpoints,omitempty"`
    Id *ServiceIdFilter `json:"id,omitempty",yaml:"id,omitempty"`
    IsVirtual *BoolFilter `json:"isVirtual,omitempty",yaml:"isVirtual,omitempty"`
    Meta *TypeMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Name *StringFilter `json:"name,omitempty",yaml:"name,omitempty"`
    Not []ServiceFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []ServiceFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Port *Int32Filter `json:"port,omitempty",yaml:"port,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
    Transport *EnumFilter `json:"transport,omitempty",yaml:"transport,omitempty"`
    Url *UrlFilter `json:"url,omitempty",yaml:"url,omitempty"`
}