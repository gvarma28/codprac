package collections

type MyHashSet struct {
	buckets [][]int
	size    int
}

func MyHashSetConstructor() MyHashSet {
	const defaultSize = 10007
	return MyHashSet{
		size:    defaultSize,
		buckets: make([][]int, defaultSize),
	}
}

func (set *MyHashSet) getHash(key int) int {
	return key % set.size
}

func (set *MyHashSet) Add(key int) {
	hash := set.getHash(key)
	for _, val := range set.buckets[hash] {
		if val == key {
			return
		}
	}
	set.buckets[hash] = append(set.buckets[hash], key)
}

func (set *MyHashSet) Remove(key int) {
	hash := set.getHash(key)
	temp := set.buckets[hash]
	for i, val := range set.buckets[hash] {
		if val == key {
			set.buckets[hash] = append(temp[:i], temp[i+1:]...)
			return
		}
	}
}

func (set *MyHashSet) Contains(key int) bool {
	hash := set.getHash(key)
	for _, val := range set.buckets[hash] {
		if val == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := MyHashSetConstructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
