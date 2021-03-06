// generated by metactl sdk gen 
package mql

const (
	PipeGetDummiesContextFilterName = "PipeGetDummiesContextFilter"
)

type PipeGetDummiesContextFilter struct {
    And []PipeGetDummiesContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientRequest *GetDummiesRequestFilter `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *GetDummiesResponseFilter `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    Not []PipeGetDummiesContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeGetDummiesContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceRequest *GetDummiesRequestFilter `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *GetDummiesResponseFilter `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}