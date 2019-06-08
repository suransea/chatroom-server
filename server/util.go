package server

import "container/list"

func Remove(list *list.List, element interface{}) {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value == element {
			list.Remove(e)
			break
		}
	}
}
