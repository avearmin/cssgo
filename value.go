package cssgo

type ValueNode interface {
	Node
	valueNode()
}
