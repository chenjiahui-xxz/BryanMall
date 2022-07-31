package enum

type from int

const (
	Nature from = 1
	Market from = 2
)

func (f from) string() string {
	switch f {
	case Nature:
		return "自然来源"
	case Market:
		return "行销来源"
	default:
		return "第三方流量平台"
	}
}