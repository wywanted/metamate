// generated by metactl sdk gen 
package sdk

const (
	StatusesCollectionName = "StatusesCollection"
)

type StatusesCollection struct {
    Meta *CollectionMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
    Statuses []Status `json:"statuses,omitempty" yaml:"statuses,omitempty"`
}