// generated by metactl sdk gen 
package mql

const (
	PostFeedSelectName = "PostFeedSelect"
)

type PostFeedSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    AlternativeIds *IdSelect `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    CreatedAt *TimestampSelect `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    Id *ServiceIdSelect `json:"id,omitempty" yaml:"id,omitempty"`
    Info *InfoSelect `json:"info,omitempty" yaml:"info,omitempty"`
    Kind *bool `json:"kind,omitempty" yaml:"kind,omitempty"`
    Relations *PostFeedRelationsSelect `json:"relations,omitempty" yaml:"relations,omitempty"`
    Relationships *PostFeedRelationshipsSelect `json:"relationships,omitempty" yaml:"relationships,omitempty"`
}