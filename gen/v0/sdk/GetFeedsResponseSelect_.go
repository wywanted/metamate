// generated by metactl sdk gen 
package sdk

const (
	GetFeedsResponseSelectName = "GetFeedsResponseSelect"
)

type GetFeedsResponseSelect struct {
    All *bool `json:"all,omitempty",yaml:"all,omitempty"`
    Feeds *FeedSelect `json:"feeds,omitempty",yaml:"feeds,omitempty"`
    Meta *CollectionMetaSelect `json:"meta,omitempty",yaml:"meta,omitempty"`
}