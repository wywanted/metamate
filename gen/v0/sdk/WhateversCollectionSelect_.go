// generated by metactl sdk gen 
package sdk

const (
	WhateversCollectionSelectName = "WhateversCollectionSelect"
)

type WhateversCollectionSelect struct {
    All *bool `json:"all,omitempty",yaml:"all,omitempty"`
    Meta *CollectionMetaSelect `json:"meta,omitempty",yaml:"meta,omitempty"`
    Whatevers *WhateverSelect `json:"whatevers,omitempty",yaml:"whatevers,omitempty"`
}