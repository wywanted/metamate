// generated by metactl sdk gen 
package mql

const (
	GetDummiesCollectionName = "GetDummiesCollection"
)

type GetDummiesCollection struct {
    Filter *DummyFilter `json:"filter,omitempty" yaml:"filter,omitempty"`
    Pages []ServicePage `json:"pages,omitempty" yaml:"pages,omitempty"`
    Relations *GetDummiesRelations `json:"relations,omitempty" yaml:"relations,omitempty"`
    Select *DummiesCollectionSelect `json:"select,omitempty" yaml:"select,omitempty"`
    ServiceFilter *ServiceFilter `json:"serviceFilter,omitempty" yaml:"serviceFilter,omitempty"`
    Sort *DummySort `json:"sort,omitempty" yaml:"sort,omitempty"`
}