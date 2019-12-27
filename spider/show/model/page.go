package model

type SearchResultData struct {
	Items []interface{} //列表
	Hits int //结果总数
	Count int //当前结果数
	PrevPage int //上一页
	NextPage int // 下一页
	Query string
}
