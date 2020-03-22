// generated by metactl sdk gen 
package sdk

const (
	AttachmentSelectName = "AttachmentSelect"
)

type AttachmentSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    AlternativeIds *IdSelect `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    Description *bool `json:"description,omitempty" yaml:"description,omitempty"`
    Id *ServiceIdSelect `json:"id,omitempty" yaml:"id,omitempty"`
    Meta *TypeMetaSelect `json:"meta,omitempty" yaml:"meta,omitempty"`
    Relations *AttachmentRelationsSelect `json:"relations,omitempty" yaml:"relations,omitempty"`
}