// generated by metactl sdk gen 
package mql

const (
	GetPostsEndpointFilterName = "GetPostsEndpointFilter"
)

type GetPostsEndpointFilter struct {
    And []GetPostsEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []GetPostsEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetPostsEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}