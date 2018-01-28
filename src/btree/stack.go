package btree

// stack based on slice
type Stack []*BTree

func (q *Stack) Push(n *BTree) {
	*q = append(*q, n)
}

func (q *Stack) Pop() (n *BTree) {
	x := q.Len() - 1
	n = (*q)[x]
	*q = (*q)[:x]
	return n
}
func (q *Stack) Len() int {
	return len(*q)
}