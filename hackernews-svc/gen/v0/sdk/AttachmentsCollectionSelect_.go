// generated by metactl sdk gen 
package sdk

const (
	AttachmentsCollectionSelectName = "AttachmentsCollectionSelect"
)

type AttachmentsCollectionSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    Attachments *AttachmentSelect `json:"attachments,omitempty" yaml:"attachments,omitempty"`
    Meta *CollectionMetaSelect `json:"meta,omitempty" yaml:"meta,omitempty"`
}