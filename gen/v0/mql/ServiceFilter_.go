// generated by metactl sdk gen 
package mql

const (
	ServiceFilterName = "ServiceFilter"
)

type ServiceFilter struct {
    AlternativeIds *IdListFilter `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    And []ServiceFilter `json:"and,omitempty" yaml:"and,omitempty"`
    CreatedAt *TimestampFilter `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    Endpoints *EndpointsFilter `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
    Id *ServiceIdFilter `json:"id,omitempty" yaml:"id,omitempty"`
    IsVirtual *BoolFilter `json:"isVirtual,omitempty" yaml:"isVirtual,omitempty"`
    Name *StringFilter `json:"name,omitempty" yaml:"name,omitempty"`
    Not []ServiceFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []ServiceFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Port *Int32Filter `json:"port,omitempty" yaml:"port,omitempty"`
    SdkVersion *StringFilter `json:"sdkVersion,omitempty" yaml:"sdkVersion,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Url *UrlFilter `json:"url,omitempty" yaml:"url,omitempty"`
}