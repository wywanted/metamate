// generated by metactl sdk gen 
package sdk

const (
	GetBlueWhateversResponseName = "GetBlueWhateversResponse"
)

type GetBlueWhateversResponse struct {
    BlueWhatevers []BlueWhatever `json:"blueWhatevers,omitempty",yaml:"blueWhatevers,omitempty"`
    Meta *CollectionMeta `json:"meta,omitempty",yaml:"meta,omitempty"`
}