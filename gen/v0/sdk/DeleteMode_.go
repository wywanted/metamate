// generated by metactl sdk gen 
package sdk

const (
	DeleteModeName = "DeleteMode"
)

type DeleteMode struct {
    Archive *bool `json:"archive,omitempty",yaml:"archive,omitempty"`
    Kind *string `json:"kind,omitempty",yaml:"kind,omitempty"`
    Permanent *bool `json:"permanent,omitempty",yaml:"permanent,omitempty"`
}