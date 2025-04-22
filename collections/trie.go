package collections

type TrieNode struct {
	Children map[string]*TrieNode
	IsEnd    bool
}

func (node *TrieNode) Insert(text string) {
	cur := node
	for _, char := range text {
		strChar := string(char)
		// if the character doesn't already exist, we create a new one
		if _, ok := cur.Children[strChar]; !ok {
			cur.Children[strChar] = &TrieNode{make(map[string]*TrieNode), false}
		}
		cur = cur.Children[strChar]
	}
	cur.IsEnd = true
}

func (node *TrieNode) Search(text string) bool {
	cur := node
	for _, char := range text {
		strChar := string(char)
		if _, ok := cur.Children[strChar]; !ok {
			return false
		}
		cur = cur.Children[strChar]
	}
	return cur.IsEnd
}

func (node *TrieNode) StartsWith(text string) bool {
	cur := node
	for _, char := range text {
		strChar := string(char)
		if _, ok := cur.Children[strChar]; !ok {
			return false
		}
		cur = cur.Children[strChar]
	}
	return true
}
