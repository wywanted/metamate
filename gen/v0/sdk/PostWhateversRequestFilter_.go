// generated by metactl sdk gen 
package sdk

const (
	PostWhateversRequestFilterName = "PostWhateversRequestFilter"
)

type PostWhateversRequestFilter struct {
    And []PostWhateversRequestFilter `json:"and,omitempty",yaml:"and,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty",yaml:"meta,omitempty"`
    Mode *PostModeFilter `json:"mode,omitempty",yaml:"mode,omitempty"`
    Not []PostWhateversRequestFilter `json:"not,omitempty",yaml:"not,omitempty"`
    Or []PostWhateversRequestFilter `json:"or,omitempty",yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty",yaml:"set,omitempty"`
    Whatevers *WhateverListFilter `json:"whatevers,omitempty",yaml:"whatevers,omitempty"`
}