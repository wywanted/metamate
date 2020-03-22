// generated by metactl sdk gen 
package sdk

const (
	PipeGetTransactionsContextFilterName = "PipeGetTransactionsContextFilter"
)

type PipeGetTransactionsContextFilter struct {
    And []PipeGetTransactionsContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientRequest *GetTransactionsRequestFilter `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *GetTransactionsResponseFilter `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    Not []PipeGetTransactionsContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeGetTransactionsContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceRequest *GetTransactionsRequestFilter `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *GetTransactionsResponseFilter `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}