module github.com/metamatex/metamate/generic

go 1.13

replace github.com/metamatex/metamate/asg => ../asg

replace github.com/metamatex/metamate/gen => ../gen

require (
	github.com/metamatex/metamate/asg v0.0.0-00010101000000-000000000000
	github.com/metamatex/metamate/gen v0.0.0-00010101000000-000000000000
	github.com/mitchellh/hashstructure v1.0.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/stretchr/testify v1.5.1
	github.com/wolfeidau/unflatten v1.0.1
	gopkg.in/yaml.v2 v2.2.8
)
