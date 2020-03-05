// generated by metactl sdk gen 
package sdk

const (
	FeedsCollectionName = "FeedsCollection"
)

type FeedsCollection struct {
    Feeds []Feed `json:"feeds,omitempty",yaml:"feeds,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Meta *CollectionMeta `json:"meta,omitempty",yaml:"meta,omitempty"`
}