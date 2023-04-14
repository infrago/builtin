package builtin

import (
	"fmt"
	"strconv"

	"github.com/infrago/base"
	"github.com/infrago/infra"
	"github.com/infrago/util"
)

// 密码加密格式
func password(str string) string {
	return util.Sha1(str)
}

func anyToString(val base.Any) string {
	sv := ""
	switch v := val.(type) {
	case string:
		sv = v
	case int:
		sv = strconv.Itoa(v)
	case int64:
		sv = strconv.FormatInt(v, 10)
	case bool:
		sv = strconv.FormatBool(v)
	case base.Map:
		d, e := infra.MarshalJSON(v)
		if e == nil {
			sv = string(d)
		} else {
			sv = "{}"
		}
	case []base.Map:
		d, e := infra.MarshalJSON(v)
		if e == nil {
			sv = string(d)
		} else {
			sv = "[]"
		}
	case []int, []int8, []int16, []int32, []int64, []float32, []float64, []string, []bool, []base.Any:
		d, e := infra.MarshalJSON(v)
		if e == nil {
			sv = string(d)
		} else {
			sv = "[]"
		}
	default:
		sv = fmt.Sprintf("%v", v)
	}

	return sv
}
