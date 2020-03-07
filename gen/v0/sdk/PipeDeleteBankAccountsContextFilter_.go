// generated by metactl sdk gen 
package sdk

const (
	PipeDeleteBankAccountsContextFilterName = "PipeDeleteBankAccountsContextFilter"
)

type PipeDeleteBankAccountsContextFilter struct {
    And []PipeDeleteBankAccountsContextFilter `json:"and,omitempty",yaml:"and,omitempty"`
    ClientRequest *DeleteBankAccountsRequestFilter `json:"clientRequest,omitempty",yaml:"clientRequest,omitempty"`
    ClientResponse *DeleteBankAccountsResponseFilter `json:"clientResponse,omitempty",yaml:"clientResponse,omitempty"`
    Not []PipeDeleteBankAccountsContextFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PipeDeleteBankAccountsContextFilter `json:"or,omitempty",yaml:"or,omitempty"`
    ServiceRequest *DeleteBankAccountsRequestFilter `json:"serviceRequest,omitempty",yaml:"serviceRequest,omitempty"`
    ServiceResponse *DeleteBankAccountsResponseFilter `json:"serviceResponse,omitempty",yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}