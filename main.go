package main

import "SkipList/skiplist"

func main() {
	skipList := skiplist.NewSkipList(7)
	skipList.Put(778, "sda")
	skipList.Put(7738, 1)
	skipList.Put(72178, true)
	skipList.Put(71378, 114)
	skipList.Put(72378, "114")
	skipList.Delete(71378)
	skipList.PrintSkipList()

}
