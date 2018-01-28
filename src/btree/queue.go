package btree

// queue based on slice
type Queue []*BTree

func (q *Queue) Push(n *BTree)Queue {
	*q = append(*q, n)
	return *q
}

func (q *Queue) Pop() (n *BTree) {
	n = (*q)[0]
	*q = (*q)[1:]
	return n
}

func (q *Queue) Len() int{
	return len(*q)
}