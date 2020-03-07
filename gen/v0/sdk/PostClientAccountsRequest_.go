// generated by metactl sdk gen 
package sdk

const (
	PostClientAccountsRequestName = "PostClientAccountsRequest"
)

type PostClientAccountsRequest struct {
    Auth *Auth `json:"auth,omitempty",yaml:"auth,omitempty"`
    ClientAccounts []ClientAccount `json:"clientAccounts,omitempty",yaml:"clientAccounts,omitempty"`
    Meta *RequestMeta `json:"meta,omitempty",yaml:"meta,omitempty"`
    Mode *PostMode `json:"mode,omitempty",yaml:"mode,omitempty"`
    Select *PostClientAccountsResponseSelect `json:"select,omitempty",yaml:"select,omitempty"`
    ServiceFilter *ServiceFilter `json:"serviceFilter,omitempty",yaml:"serviceFilter,omitempty"`
}