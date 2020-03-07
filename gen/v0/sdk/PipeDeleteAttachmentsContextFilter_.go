// generated by metactl sdk gen 
package sdk

const (
	PipeDeleteAttachmentsContextFilterName = "PipeDeleteAttachmentsContextFilter"
)

type PipeDeleteAttachmentsContextFilter struct {
    And []PipeDeleteAttachmentsContextFilter `json:"and,omitempty",yaml:"and,omitempty"`
    ClientRequest *DeleteAttachmentsRequestFilter `json:"clientRequest,omitempty",yaml:"clientRequest,omitempty"`
    ClientResponse *DeleteAttachmentsResponseFilter `json:"clientResponse,omitempty",yaml:"clientResponse,omitempty"`
    Not []PipeDeleteAttachmentsContextFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PipeDeleteAttachmentsContextFilter `json:"or,omitempty",yaml:"or,omitempty"`
    ServiceRequest *DeleteAttachmentsRequestFilter `json:"serviceRequest,omitempty",yaml:"serviceRequest,omitempty"`
    ServiceResponse *DeleteAttachmentsResponseFilter `json:"serviceResponse,omitempty",yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}