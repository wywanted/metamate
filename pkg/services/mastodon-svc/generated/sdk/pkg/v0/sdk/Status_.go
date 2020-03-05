// generated by metactl sdk gen 
package sdk

const (
	StatusName = "Status"
)

type Status struct {
    AlternativeIds []Id `json:"alternativeIds,omitempty",yaml:"alternativeIds,omitempty"`
    Content *Text `json:"content,omitempty",yaml:"content,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Id *ServiceId `json:"id,omitempty",yaml:"id,omitempty"`
    Meta *TypeMeta `json:"meta,omitempty",yaml:"meta,omitempty"`
    Pinned *bool `json:"pinned,omitempty",yaml:"pinned,omitempty"`
    Relations *StatusRelations `json:"relations,omitempty",yaml:"relations,omitempty"`
    Sensitive *bool `json:"sensitive,omitempty",yaml:"sensitive,omitempty"`
    SpoilerText *Text `json:"spoilerText,omitempty",yaml:"spoilerText,omitempty"`
}