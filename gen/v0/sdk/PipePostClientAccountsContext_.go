// generated by metactl sdk gen 
package sdk

const (
	PipePostClientAccountsContextName = "PipePostClientAccountsContext"
)

type PipePostClientAccountsContext struct {
    ClientRequest *PostClientAccountsRequest `json:"clientRequest,omitempty",yaml:"clientRequest,omitempty"`
    ClientResponse *PostClientAccountsResponse `json:"clientResponse,omitempty",yaml:"clientResponse,omitempty"`
    ServiceRequest *PostClientAccountsRequest `json:"serviceRequest,omitempty",yaml:"serviceRequest,omitempty"`
    ServiceResponse *PostClientAccountsResponse `json:"serviceResponse,omitempty",yaml:"serviceResponse,omitempty"`
}