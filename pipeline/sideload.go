package pipeline

// Sideload adds data to points based on hierarchical data from various sources.
// Sideload behaves similarly to the DefaultNode but loads values from external sources using a hierarchy.
//
// Example:
//        |sideload()
//             .source('file:///path/to/dir')
//             .order('host/{host}.yml', 'hostgroup/{hostgroup}.yml', default.yml'
//             .field('cpu_threshold')
//             .tag('foo')
//
// Add a field `cpu_threshold` and a tag `foo` to each point based on the value loaded from the hierarchical source.
type SideloadNode struct {
	chainnode

	// Source for the data, currently only `file://` based sources are supported
	Source string

	// Order is a list of paths that indicate the hierarchical order.
	// The paths are relative to the source and can have `{}` that will be replaced with the tag value from the point.
	// This allows for values to be overridden based on a hierarchy of tags.
	// tick:ignore
	OrderList []string `tick:"Order"`

	// Fields is a list of fields to load.
	// tick:ignore
	Fields map[string]interface{} `tick:"Field"`
	// Tags is a list of tags to load.
	// tick:ignore
	Tags map[string]string `tick:"Tag"`
}

func newSideloadNode(wants EdgeType) *SideloadNode {
	return &SideloadNode{
		chainnode: newBasicChainNode("sideload", wants, wants),
		Fields:    make(map[string]interface{}),
		Tags:      make(map[string]string),
	}
}

// Order is a list of paths that indicate the hierarchical order.
// The paths are relative to the source and can have `{}` that will be replaced with the tag value from the point.
// This allows for values to be overridden based on a hierarchy of tags.
// tick:property
func (n *SideloadNode) Order(order ...string) *SideloadNode {
	n.OrderList = order
	return n
}

// Field is the name of a field to load from the source and its default value.
// The type loaded must match the type of the default value.
// Otherwise an error is recorded and the default value is used.
// tick:property
func (n *SideloadNode) Field(f string, v interface{}) *SideloadNode {
	n.Fields[f] = v
	return n
}

// Tag is the name of a tag to load from the source and its default value.
// tick:property
func (n *SideloadNode) Tag(t string, v string) *SideloadNode {
	n.Tags[t] = v
	return n
}