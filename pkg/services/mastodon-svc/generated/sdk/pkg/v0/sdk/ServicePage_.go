// generated by metactl sdk gen 
package sdk

const (
	ServicePageName = "ServicePage"
)

type ServicePage struct {
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Page *Page `json:"page,omitempty",yaml:"page,omitempty"`
    Service *Service `json:"service,omitempty",yaml:"service,omitempty"`
}