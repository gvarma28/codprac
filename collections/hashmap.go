package collections

type MyHashMap struct {
	buckets [][]mapNode
	size    int
}

type mapNode struct {
	key int
	val int
}

func MyHashMapConstructor() MyHashMap {
	const defaultSize = 10007
	return MyHashMap{
		size:    defaultSize,
		buckets: make([][]mapNode, defaultSize),
	}
}

func (hashmap *MyHashMap) getHash(key int) int {
	return key % hashmap.size
}

func (hashmap *MyHashMap) Put(key int, value int) {
	hash := hashmap.getHash(key)
	for i := range hashmap.buckets[hash] {
		if hashmap.buckets[hash][i].key == key {
			hashmap.buckets[hash][i].val = value
			return
		}
	}
	hashmap.buckets[hash] = append(hashmap.buckets[hash], mapNode{key, value})
}

func (hashmap *MyHashMap) Get(key int) int {
	hash := hashmap.getHash(key)
	for _, node := range hashmap.buckets[hash] {
		if node.key == key {
			return node.val
		}
	}
	return -1
}

func (hashmap *MyHashMap) Remove(key int) {
	hash := hashmap.getHash(key)
	temp := hashmap.buckets[hash]
	for i, node := range hashmap.buckets[hash] {
		if node.key == key {
			hashmap.buckets[hash] = append(temp[:i], temp[i+1:]...)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := MyHashMapConstructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
