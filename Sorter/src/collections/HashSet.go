package collections

import (
	"strings"
	"fmt"
	"bytes"
)


/*
go 提供的map中key不能为func,map或者slice
*/
type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}


func (set *HashSet)Delete(e interface){
	delete(set.m,e)
}


func (set *HashSet)Clear(){
	set.m=make(map[interface{}]bool)
}


func (set *HashSet)Contains(e interface{}) bool{
	return set.m[e]
}

func (set *HashSet)Len() int{
	return len(set.m)
}


func (set *HashSet)Elements() []interface{}{
	initialLen:=len(set.m)
	snapshot:=make([]interface{},initialLen)
	actualLen:=0
	for k:=range set.m{
		if actualLen<initialLen{
			snapshot[actualLen]=k
		}else{
			snapshot=append(snapshot,k)
		}
		actualLen++
	}
	if actualLen<initialLen{
		snapshot=snapshot[:actualLen]
	}
	return snapshot
}


func (set *HashSet)String() string{
	var buf bytes.Buffer
	buf.WriteString("HashSet{")
	for element:=range set.m{
		buf.WriteString(fmt.Sprintf("%v ",element))
	}
	str:=buf.String()
	return strings.Replace(str," ","}",len(str)-1)
}