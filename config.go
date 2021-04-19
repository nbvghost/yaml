package yaml

var Default = struct {
	UseFieldName bool //如果struct没有定义yaml标签，则默认使用struct 字段名
}{}
