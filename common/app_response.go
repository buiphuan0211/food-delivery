package common

type successRes struct {
	Data   interface{} `json:"data"`
	Page   interface{} `json:"page,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, page, filter interface{}) *successRes {
	return &successRes{Data: data, Page: page, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
