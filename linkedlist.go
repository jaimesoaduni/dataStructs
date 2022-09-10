package dataStructs

type node[T any] struct {
	parent, child *node[T]
	data          T
}

func (n *node[T]) Next() *node[T] {
	return n.child
}

func (n *node[T]) Prev() *node[T] {
	return n.parent
}

type List[T any] struct {
	root, leaf *node[T]
	length     int
}

func (l *List[T]) First() *node[T] {
	return l.root
}

func (l *List[T]) Last() *node[T] {
	return l.leaf
}

func (l *List[T]) PushBack(data T) {
	if l.root == nil {
		l.root = new(node[T])
		l.root.data = data
		l.leaf = l.root
		l.length = 1
		return
	}

	newNode := new(node[T])

	newNode.data = data
	newNode.parent = l.leaf

	l.leaf.child = newNode
	l.leaf = newNode

	l.length++
}
