// generated by metactl sdk gen 
package sdk

const (
	PostSelectName = "PostSelect"
)

type PostSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    AlternativeIds *IdSelect `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    Content *TextSelect `json:"content,omitempty" yaml:"content,omitempty"`
    CreatedAt *TimestampSelect `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    Id *ServiceIdSelect `json:"id,omitempty" yaml:"id,omitempty"`
    IsPinned *bool `json:"isPinned,omitempty" yaml:"isPinned,omitempty"`
    IsSensitive *bool `json:"isSensitive,omitempty" yaml:"isSensitive,omitempty"`
    Links *HyperLinkSelect `json:"links,omitempty" yaml:"links,omitempty"`
    Relations *PostRelationsSelect `json:"relations,omitempty" yaml:"relations,omitempty"`
    Relationships *PostRelationshipsSelect `json:"relationships,omitempty" yaml:"relationships,omitempty"`
    SpoilerText *TextSelect `json:"spoilerText,omitempty" yaml:"spoilerText,omitempty"`
    Title *TextSelect `json:"title,omitempty" yaml:"title,omitempty"`
}