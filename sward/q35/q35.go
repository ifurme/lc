package q35

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	//原地复制并插入到当前节点后面
	cur := head
	for ; cur != nil; cur = cur.Next.Next {
		bak := &Node{Val: cur.Val}
		bak.Next = cur.Next
		cur.Next = bak
	}
	cur = head
	//Random指针赋值
	for ; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	//拆链
	cur = head
	ret := head.Next
	for ; cur != nil; cur = cur.Next {
		next := cur.Next
		cur.Next = next.Next
		if next.Next != nil {
			next.Next = next.Next.Next
		}
	}
	return ret
}
