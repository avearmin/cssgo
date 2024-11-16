package gomour

type ValueNode interface {
	Node
	valueNode()
}
