// generated by metactl sdk gen 
package mql

const (
	PostFeedName = "PostFeed"
)

type PostFeed struct {
    AlternativeIds []Id `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    CreatedAt *Timestamp `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    Id *ServiceId `json:"id,omitempty" yaml:"id,omitempty"`
    Info *Info `json:"info,omitempty" yaml:"info,omitempty"`
    Kind *string `json:"kind,omitempty" yaml:"kind,omitempty"`
    Relations *PostFeedRelations `json:"relations,omitempty" yaml:"relations,omitempty"`
    Relationships *PostFeedRelationships `json:"relationships,omitempty" yaml:"relationships,omitempty"`
}