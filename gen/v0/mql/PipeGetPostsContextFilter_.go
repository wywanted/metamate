// generated by metactl sdk gen 
package mql

const (
	PipeGetPostsContextFilterName = "PipeGetPostsContextFilter"
)

type PipeGetPostsContextFilter struct {
    And []PipeGetPostsContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientRequest *GetPostsRequestFilter `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *GetPostsResponseFilter `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    Not []PipeGetPostsContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeGetPostsContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceRequest *GetPostsRequestFilter `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *GetPostsResponseFilter `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}