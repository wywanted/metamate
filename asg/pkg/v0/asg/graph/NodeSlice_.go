// generated by go:generate go run gen/nodeslice.go
package graph

import (
	"sort"
)

type BasicTypeNodeSlice []*BasicTypeNode

func (nm BasicTypeNodeSlice) GetIds() (ids []NodeId) {
	for _, n := range nm {
		ids = append(ids, n.Id())
	}

	return
}

func (ns BasicTypeNodeSlice) Sort() (BasicTypeNodeSlice) {
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Namex < ns[j].Namex
	})

	return ns
}

func (ns BasicTypeNodeSlice) Map() (fnm BasicTypeNodeMap) {
	fnm = BasicTypeNodeMap{}

	fnm.Add(ns...)

	return
}

type EndpointNodeSlice []*EndpointNode

func (nm EndpointNodeSlice) GetIds() (ids []NodeId) {
	for _, n := range nm {
		ids = append(ids, n.Id())
	}

	return
}

func (ns EndpointNodeSlice) Sort() (EndpointNodeSlice) {
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Namex < ns[j].Namex
	})

	return ns
}

func (ns EndpointNodeSlice) Map() (fnm EndpointNodeMap) {
	fnm = EndpointNodeMap{}

	fnm.Add(ns...)

	return
}

type EnumNodeSlice []*EnumNode

func (nm EnumNodeSlice) GetIds() (ids []NodeId) {
	for _, n := range nm {
		ids = append(ids, n.Id())
	}

	return
}

func (ns EnumNodeSlice) Sort() (EnumNodeSlice) {
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Namex < ns[j].Namex
	})

	return ns
}

func (ns EnumNodeSlice) Map() (fnm EnumNodeMap) {
	fnm = EnumNodeMap{}

	fnm.Add(ns...)

	return
}

type FieldNodeSlice []*FieldNode

func (nm FieldNodeSlice) GetIds() (ids []NodeId) {
	for _, n := range nm {
		ids = append(ids, n.Id())
	}

	return
}

func (ns FieldNodeSlice) Sort() (FieldNodeSlice) {
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Namex < ns[j].Namex
	})

	return ns
}

func (ns FieldNodeSlice) Map() (fnm FieldNodeMap) {
	fnm = FieldNodeMap{}

	fnm.Add(ns...)

	return
}

type RelationNodeSlice []*RelationNode

func (nm RelationNodeSlice) GetIds() (ids []NodeId) {
	for _, n := range nm {
		ids = append(ids, n.Id())
	}

	return
}

func (ns RelationNodeSlice) Sort() (RelationNodeSlice) {
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Namex < ns[j].Namex
	})

	return ns
}

func (ns RelationNodeSlice) Map() (fnm RelationNodeMap) {
	fnm = RelationNodeMap{}

	fnm.Add(ns...)

	return
}

type TypeNodeSlice []*TypeNode

func (nm TypeNodeSlice) GetIds() (ids []NodeId) {
	for _, n := range nm {
		ids = append(ids, n.Id())
	}

	return
}

func (ns TypeNodeSlice) Sort() (TypeNodeSlice) {
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Namex < ns[j].Namex
	})

	return ns
}

func (ns TypeNodeSlice) Map() (fnm TypeNodeMap) {
	fnm = TypeNodeMap{}

	fnm.Add(ns...)

	return
}

type PathNodeSlice []*PathNode

func (nm PathNodeSlice) GetIds() (ids []NodeId) {
	for _, n := range nm {
		ids = append(ids, n.Id())
	}

	return
}

func (ns PathNodeSlice) Sort() (PathNodeSlice) {
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Namex < ns[j].Namex
	})

	return ns
}

func (ns PathNodeSlice) Map() (fnm PathNodeMap) {
	fnm = PathNodeMap{}

	fnm.Add(ns...)

	return
}
