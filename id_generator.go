package gid

import "fmt"

//生成id的规则
type IdGenerateRule interface{

	//规则
	GetRule()string

	//前缀
	GetPrefix()string

	//后缀
	GetSuffix()string
}

//id生成器
//num：生成的数量
func Generate(rule IdGenerateRule)string{
	return fmt.Sprintf("%s%s%s",rule.GetPrefix(),rule.GetRule(),rule.GetSuffix())
}