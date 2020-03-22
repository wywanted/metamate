// generated by metactl sdk gen 
package sdk

const (
	ImageSelectName = "ImageSelect"
)

type ImageSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    Description *TextSelect `json:"description,omitempty" yaml:"description,omitempty"`
    Height *bool `json:"height,omitempty" yaml:"height,omitempty"`
    IsPreview *bool `json:"isPreview,omitempty" yaml:"isPreview,omitempty"`
    Url *UrlSelect `json:"url,omitempty" yaml:"url,omitempty"`
    Width *bool `json:"width,omitempty" yaml:"width,omitempty"`
}