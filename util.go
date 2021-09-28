package LRU_Cache

func MList2Ints(lru *CLRUCache) [][]interface{} {
	res := [][]interface{}{}
	for head := lru.list.Front(); head != nil; head = head.Next() {
		tmp := []interface{}{head.Value.(Pair).key, head.Value.(Pair).value}
		res = append(res, tmp)
	}
	return res
}
