// generated by metactl sdk gen 
package sdk

const (
	VerifyTokenOutputName = "VerifyTokenOutput"
)

type VerifyTokenOutput struct {
    ClientAccountId *ServiceId `json:"clientAccountId,omitempty",yaml:"clientAccountId,omitempty"`
    IsValid *bool `json:"isValid,omitempty",yaml:"isValid,omitempty"`
}