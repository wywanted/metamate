// generated by metactl sdk gen 
package sdk

const (
	PipeGetBlueWhateversContextFilterName = "PipeGetBlueWhateversContextFilter"
)

type PipeGetBlueWhateversContextFilter struct {
    And []PipeGetBlueWhateversContextFilter `json:"and,omitempty",yaml:"and,omitempty"`
    ClientRequest *GetBlueWhateversRequestFilter `json:"clientRequest,omitempty",yaml:"clientRequest,omitempty"`
    ClientResponse *GetBlueWhateversResponseFilter `json:"clientResponse,omitempty",yaml:"clientResponse,omitempty"`
    Not []PipeGetBlueWhateversContextFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PipeGetBlueWhateversContextFilter `json:"or,omitempty",yaml:"or,omitempty"`
    ServiceRequest *GetBlueWhateversRequestFilter `json:"serviceRequest,omitempty",yaml:"serviceRequest,omitempty"`
    ServiceResponse *GetBlueWhateversResponseFilter `json:"serviceResponse,omitempty",yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
}