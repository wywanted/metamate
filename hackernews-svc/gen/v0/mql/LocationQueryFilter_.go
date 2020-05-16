// generated by metactl sdk gen 
package mql

const (
	LocationQueryFilterName = "LocationQueryFilter"
)

type LocationQueryFilter struct {
    And []LocationQueryFilter `json:"and,omitempty" yaml:"and,omitempty"`
    City *StringFilter `json:"city,omitempty" yaml:"city,omitempty"`
    CityDistrict *StringFilter `json:"cityDistrict,omitempty" yaml:"cityDistrict,omitempty"`
    Country *StringFilter `json:"country,omitempty" yaml:"country,omitempty"`
    CountryState *StringFilter `json:"countryState,omitempty" yaml:"countryState,omitempty"`
    CountryStateDistrict *StringFilter `json:"countryStateDistrict,omitempty" yaml:"countryStateDistrict,omitempty"`
    Not []LocationQueryFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []LocationQueryFilter `json:"or,omitempty" yaml:"or,omitempty"`
    RadiusLt *LengthScalarFilter `json:"radiusLt,omitempty" yaml:"radiusLt,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Street *StringFilter `json:"street,omitempty" yaml:"street,omitempty"`
    ZipCode *StringFilter `json:"zipCode,omitempty" yaml:"zipCode,omitempty"`
}