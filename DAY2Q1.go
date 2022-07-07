package main

import "fmt"


func main() {
	wordList:=[]string{"quick","brown","fox","lazy","dog"}
	m:=make(map[string]int)
	count:=func (s string) {
		for i:=0;i<len(s);i++{
			m[string(s[i])]++
		}
	}
	for _,entity:=range wordList {
		go count(entity)
	}
	fmt.Println(m)
}

//version 2
//package main
//
//import "fmt"
//
//
//func main() {
//	wordList:=[]string{"quick","brown","fox","lazy","dog"}
//	ch :=make(chan string)
//	m:=make(map[string]int)
//	count:=func (ch <-chan string) {
//
//		s:=<-ch
//		for i:=0;i<len(s);i++{
//			m[string(s[i])]++
//		}
//	}
//	for _,entity:=range wordList {
//
//		go count(ch)
//		ch<-entity
//	}
//	fmt.Println(m)
//}
