package gid

import (
	"github.com/zouyx/gomatcher"
	"testing"
)

func TestGenerate(t *testing.T) {
	rule := &StandardRule{}
	id := Generate(rule)
	gomatcher.Assert(t,id,gomatcher.Not(gomatcher.NilVal()))
}

func TestGenerateNotRepeat(t *testing.T) {
	uniqueIdMap:=make(map[string]int)
	rule := &StandardRule{}
	for i := 0; i < 1000; i++ {
		id := Generate(rule)
		uniqueIdMap[id]+=1
	}

	for _, num := range uniqueIdMap {
		gomatcher.Assert(t,num,gomatcher.Equal(1))
	}
}