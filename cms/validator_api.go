package cms

import "context"

type Validator interface {
	Check(ctx context.Context, value interface{}) (ret bool)
}

type ValidRule struct {
	RuleName	string `json:"rule_name"`
	Rule        interface{} `json:"rule"`
}

//正则
type RegularRule struct {
	Rule string `json:"rule"`
}

func (r RegularRule) Check(ctx context.Context, value interface{}) (ret bool) {
	return
}

//最大值
type MaxValue struct {
	Rule interface{} `json:"rule"`
}

func (m MaxValue) Check(ctx context.Context, value interface{}) (ret bool) {
	return
}

///...........
