// generated by metactl sdk gen 
package sdk

const (
	PipePostServiceAccountsContextFilterName = "PipePostServiceAccountsContextFilter"
)

type PipePostServiceAccountsContextFilter struct {
    And []PipePostServiceAccountsContextFilter `json:"and,omitempty",yaml:"and,omitempty"`
    ClientRequest *PostServiceAccountsRequestFilter `json:"clientRequest,omitempty",yaml:"clientRequest,omitempty"`
    ClientResponse *PostServiceAccountsResponseFilter `json:"clientResponse,omitempty",yaml:"clientResponse,omitempty"`
    Not []PipePostServiceAccountsContextFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PipePostServiceAccountsContextFilter `json:"or,omitempty",yaml:"or,omitempty"`
    ServiceRequest *PostServiceAccountsRequestFilter `json:"serviceRequest,omitempty",yaml:"serviceRequest,omitempty"`
    ServiceResponse *PostServiceAccountsResponseFilter `json:"serviceResponse,omitempty",yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}