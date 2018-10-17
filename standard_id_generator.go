package gid

import (
	"math/rand"
	"strconv"
	"time"
)

type StandardRule struct {
	
}
func (this *StandardRule) GetPrefix() string {
	return ""
}


func (this *StandardRule) GetRule() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Int())
}

func (this *StandardRule) GetSuffix() string {
	return ""
}