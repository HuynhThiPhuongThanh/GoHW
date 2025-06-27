// Last updated: 6/28/2025, 3:44:27 AM
type MyLinkedList struct {
    head *ListNode
    size int
}

func Constructor() MyLinkedList {
    return MyLinkedList{head: nil, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    current := this.head
    for i := 0; i < index; i++ {
        current = current.Next
    }
    return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
    newNode := &ListNode{Val: val, Next: this.head}
    this.head = newNode
    this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
    newNode := &ListNode{Val: val}
    if this.size == 0 {
        this.head = newNode
    } else {
        current := this.head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index < 0 || index > this.size {
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    newNode := &ListNode{Val: val}
    current := this.head
    for i := 0; i < index-1; i++ {
        current = current.Next
    }
    newNode.Next = current.Next
    current.Next = newNode
    this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }
    if index == 0 {
        this.head = this.head.Next 
    } else {
        current := this.head
        for i := 0; i < index-1; i++ {
            current = current.Next
        }
        current.Next = current.Next.Next
    }
    this.size--
}