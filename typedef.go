package builtin

import (
	"github.com/infrago/base"
	"github.com/infrago/infra"
)

func Bool(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["group"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["group"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"group": true}})
	}
	return infra.Define("bool", require, def, name, exts...)
}
func Bools(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("bools", require, def, name, exts...)
}

func Int(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("int", require, def, name, exts...)
}
func Ints(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["array"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["array"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"array": true}})
	}
	return infra.Define("[int]", require, def, name, exts...)
}

func IntKey(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["group"] = true
			mm["key"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["group"] = true
			vv.Setting["key"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"key": true, "group": true, "order": true}})
	}
	return infra.Define("int", require, def, name, exts...)
}
func IntGroup(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["group"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["group"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"group": true}})
	}
	return infra.Define("int", require, def, name, exts...)
}
func IntOrder(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["order"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["order"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"order": true}})
	}
	return infra.Define("int", require, def, name, exts...)
}

func Float(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("float", require, def, name, exts...)
}
func Floats(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("[float]", require, def, name, exts...)
}
func FloatOrder(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["order"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["group"] = true
			vv.Setting["order"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"order": true}})
	}
	return infra.Define("float", require, def, name, exts...)
}
func String(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("string", require, def, name, exts...)
}

func StringKey(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["group"] = true
			mm["query"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["group"] = true
			vv.Setting["query"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"group": true, "query": true}})
	}
	return infra.Define("string", require, def, name, exts...)
}
func StringQuery(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["query"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["query"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"query": true}})
	}
	return infra.Define("string", require, def, name, exts...)
}
func Strings(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("[string]", require, def, name, exts...)
}
func Datetime(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["order"] = true
			mm["range"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["order"] = true
			vv.Setting["range"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"order": true, "range": true}})
	}
	return infra.Define("datetime", require, def, name, exts...)
}
func Datetimes(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("[datetime]", require, def, name, exts...)
}
func Enum(require bool, def base.Any, name string, options base.Map, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["options"] = options
			mm["group"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			vv.Options = options
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["group"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Options: options, Setting: base.Map{"group": true}})
	}
	return infra.Define("enum", require, def, name, exts...)
}
func Enums(require bool, def base.Any, name string, options base.Map, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["options"] = options
			mm["group"] = true
			mm["array"] = true
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			vv.Options = options
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["group"] = true
			vv.Setting["array"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Options: options, Setting: base.Map{"group": true, "array": true}})
	}
	return infra.Define("[enum]", require, def, name, exts...)
}
func Json(require bool, def base.Any, name string, children base.Vars, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["children"] = children
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			vv.Children = children
			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Children: children})
	}
	return infra.Define("json", require, def, name, exts...)
}
func Jsons(require bool, def base.Any, name string, children base.Vars, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["children"] = children
			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			vv.Children = children
			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Children: children})
	}
	return infra.Define("[json]", require, def, name, exts...)
}

func Map(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("map", require, def, name, exts...)
}
func Maps(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("[map]", require, def, name, exts...)
}

func Any(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("any", require, def, name, exts...)
}
func Anys(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("[any]", require, def, name, exts...)
}

func Password(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("password", require, def, name, exts...)
}
func Email(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["query"] = true

			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["query"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"query": true}})
	}
	return infra.Define("email", require, def, name, exts...)
}
func Mobile(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["query"] = true

			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["query"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"query": true}})
	}
	return infra.Define("mobile", require, def, name, exts...)
}
func Phone(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	if len(exts) > 0 {
		if mm, ok := exts[0].(base.Map); ok {
			mm["query"] = true

			exts[0] = mm
		} else if vv, ok := exts[0].(base.Var); ok {
			if vv.Setting == nil {
				vv.Setting = base.Map{}
			}
			vv.Setting["query"] = true

			exts[0] = vv
		}
	} else {
		exts = append(exts, base.Var{Setting: base.Map{"query": true}})
	}
	return infra.Define("phone", require, def, name, exts...)
}

func Timestamp(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("timestamp", require, def, name, exts...)
}

func File(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("file", require, def, name, exts...)
}
func Files(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("[file]", require, def, name, exts...)
}

func Image(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("image", require, def, name, exts...)
}
func Images(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("[image]", require, def, name, exts...)
}

func Audio(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("audio", require, def, name, exts...)
}
func Audios(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("[audio]", require, def, name, exts...)
}

func Video(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("video", require, def, name, exts...)
}
func Videos(require bool, def base.Any, name string, exts ...base.Any) base.Var {
	return infra.Define("[video]", require, def, name, exts...)
}

// func TimestampConfig(name string, texts ...string) base.Var {
// 	text := name
// 	if len(texts) > 0 {
// 		text = texts[0]
// 	}
// 	return infra.Define("timestamp", true, 0, name, base.Var{Text: text})
// }
