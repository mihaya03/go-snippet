package main

// LinkedListの各操作の計算量は以下の通り
//
// | Operation       | Time Complexity |
// |-----------------|-----------------|
// | Read            | O(n)            |
// | IndexOf         | O(n)            |
// | Append          | O(n)            |
// | InsertAtIndex   | O(n)            |
// | DeleteAtIndex   | O(n)            |

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

// Read メソッドはリストの指定された位置のノードを返す
func (l *LinkedList) Read(index int) *Node {
	current := l.Head
	for i := 0; i < index; i++ {
		if current.Next == nil {
			return nil
		}
		current = current.Next
	}
	return current
}

// IndexOf メソッドはリストの中で指定された値を持つノードの位置を返す
func (l *LinkedList) IndexOf(value int) int {
	// リストの先頭点から走査する
	current := l.Head
	current_index := 0
	for current != nil {
		if current.Value == value {
			return current_index
		}
		current = current.Next
		current_index++
	}
	return -1
}

// Append メソッドはリストの末尾に新しいノードを追加する
func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// InsertAtIndex メソッドはリストの指定された位置に新しいノードを追加する
func (l *LinkedList) InsertAtIndex(index, value int) {
	newNode := &Node{Value: value}
	// リストの先頭に挿入する場合はHeadを更新する
	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		return
	}

	// リストの先頭からindex-1番目のノードを探す
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
}

// DeleteAtIndex メソッドはリストの指定された位置のノードを削除する
func (l *LinkedList) DeleteAtIndex(index int) {
	// リストの先頭を削除する場合はHeadを更新する
	if index == 0 {
		l.Head = l.Head.Next
		return
	}
	// リストの先頭からindex-1番目のノードを探す
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
}

func main() {
	// リストの作成
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// リストの読み込み
	println(list.Read(0).Value) // 1

	// リストの検索
	println(list.IndexOf(3)) // 2

	// リストの挿入
	list.InsertAtIndex(2, 10)
	println(list.Read(2).Value) // 10

	// リストの削除
	list.DeleteAtIndex(2)
	println(list.Read(2).Value) // 3

}
