// generated by metactl sdk gen 
package sdk

const (
	PostStatusesResponseName = "PostStatusesResponse"
)

type PostStatusesResponse struct {
    Meta *ResponseMeta `json:"meta,omitempty",yaml:"meta,omitempty"`
    Statuses []Status `json:"statuses,omitempty",yaml:"statuses,omitempty"`
}