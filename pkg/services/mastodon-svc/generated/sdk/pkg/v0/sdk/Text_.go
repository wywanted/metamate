// generated by metactl sdk gen 
package sdk

const (
	TextName = "Text"
)

type Text struct {
    Formatting *string `json:"formatting,omitempty",yaml:"formatting,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    Language *string `json:"language,omitempty",yaml:"language,omitempty"`
    Value *string `json:"value,omitempty",yaml:"value,omitempty"`
}