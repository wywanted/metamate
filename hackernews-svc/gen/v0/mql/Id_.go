// generated by metactl sdk gen 
package mql

const (
	IdName = "Id"
)

type Id struct {
    Ean *string `json:"ean,omitempty" yaml:"ean,omitempty"`
    Email *Email `json:"email,omitempty" yaml:"email,omitempty"`
    Kind *string `json:"kind,omitempty" yaml:"kind,omitempty"`
    Local *string `json:"local,omitempty" yaml:"local,omitempty"`
    Me *bool `json:"me,omitempty" yaml:"me,omitempty"`
    Name *string `json:"name,omitempty" yaml:"name,omitempty"`
    ServiceId *ServiceId `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`
    Url *Url `json:"url,omitempty" yaml:"url,omitempty"`
    Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}