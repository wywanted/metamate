// generated by metactl sdk gen 
package mql

const (
	LengthScalarName = "LengthScalar"
)

type LengthScalar struct {
    IsEstimate *bool `json:"isEstimate,omitempty" yaml:"isEstimate,omitempty"`
    Unit *string `json:"unit,omitempty" yaml:"unit,omitempty"`
    Value *float64 `json:"value,omitempty" yaml:"value,omitempty"`
}