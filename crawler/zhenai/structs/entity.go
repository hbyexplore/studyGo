package structs

type Request struct {
	Url        string
	ParserFunc func(*[]byte) *ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type People struct {
	/*
		当前用户基础信息
	*/
	BasicInfo []string
	/*
		当前用户详细信息
	*/
	DetailInfo []string
	/*
		当前用户择偶信息
	*/
	ObjectInfo []string
	/*
		内心独白
	*/
	InnerOs string
	/*
		兴趣爱好
	*/
	PeopleHobbies map[string]string
}
