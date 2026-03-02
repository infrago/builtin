package builtin

import (
	. "github.com/infrago/base"
)

// Define creates a Var with common fields.
func Define(kind string, require bool, def Any, name string, exts ...Any) Var {
	v := Var{
		Type:     kind,
		Required: require,
		Name:     name,
		Default:  def,
	}
	return mergeExt(v, exts...)
}

func Bool(require bool, def Any, name string, exts ...Any) Var {
	v := Define("bool", require, def, name)
	v = setFlags(v, Map{"group": true})
	return mergeExt(v, exts...)
}

func Bools(require bool, def Any, name string, exts ...Any) Var {
	// keep type consistent with builtin registered array type.
	return Define("[bool]", require, def, name, exts...)
}

func Int(require bool, def Any, name string, exts ...Any) Var {
	return Define("int", require, def, name, exts...)
}

func Ints(require bool, def Any, name string, exts ...Any) Var {
	v := Define("[int]", require, def, name)
	v = setFlags(v, Map{"array": true})
	return mergeExt(v, exts...)
}

func IntKey(require bool, def Any, name string, exts ...Any) Var {
	v := Define("int", require, def, name)
	v = setFlags(v, Map{"group": true, "key": true, "order": true})
	return mergeExt(v, exts...)
}

func IntGroup(require bool, def Any, name string, exts ...Any) Var {
	v := Define("int", require, def, name)
	v = setFlags(v, Map{"group": true})
	return mergeExt(v, exts...)
}

func IntOrder(require bool, def Any, name string, exts ...Any) Var {
	v := Define("int", require, def, name)
	v = setFlags(v, Map{"order": true})
	return mergeExt(v, exts...)
}

func Float(require bool, def Any, name string, exts ...Any) Var {
	return Define("float", require, def, name, exts...)
}

func Floats(require bool, def Any, name string, exts ...Any) Var {
	return Define("[float]", require, def, name, exts...)
}

func FloatOrder(require bool, def Any, name string, exts ...Any) Var {
	v := Define("float", require, def, name)
	v = setFlags(v, Map{"order": true})
	return mergeExt(v, exts...)
}

func String(require bool, def Any, name string, exts ...Any) Var {
	return Define("string", require, def, name, exts...)
}

func StringKey(require bool, def Any, name string, exts ...Any) Var {
	v := Define("string", require, def, name)
	v = setFlags(v, Map{"group": true, "query": true})
	return mergeExt(v, exts...)
}

func StringQuery(require bool, def Any, name string, exts ...Any) Var {
	v := Define("string", require, def, name)
	v = setFlags(v, Map{"query": true})
	return mergeExt(v, exts...)
}

func Strings(require bool, def Any, name string, exts ...Any) Var {
	return Define("[string]", require, def, name, exts...)
}

func Datetime(require bool, def Any, name string, exts ...Any) Var {
	v := Define("datetime", require, def, name)
	v = setFlags(v, Map{"order": true, "range": true})
	return mergeExt(v, exts...)
}

func Datetimes(require bool, def Any, name string, exts ...Any) Var {
	return Define("[datetime]", require, def, name, exts...)
}

func Enum(require bool, def Any, name string, options Map, exts ...Any) Var {
	v := Define("enum", require, def, name)
	v.Options = options
	v = setFlags(v, Map{"group": true})
	return mergeExt(v, exts...)
}

func Enums(require bool, def Any, name string, options Map, exts ...Any) Var {
	v := Define("[enum]", require, def, name)
	v.Options = options
	v = setFlags(v, Map{"group": true, "array": true})
	return mergeExt(v, exts...)
}

func Json(require bool, def Any, name string, children Vars, exts ...Any) Var {
	v := Define("json", require, def, name)
	v.Children = children
	return mergeExt(v, exts...)
}

func Password(require bool, def Any, name string, exts ...Any) Var {
	v := Define("string", require, def, name)
	v.Check = "password"
	return mergeExt(v, exts...)
}

func Email(require bool, def Any, name string, exts ...Any) Var {
	v := Define("string", require, def, name)
	v.Check = "email"
	return mergeExt(v, exts...)
}

func Mobile(require bool, def Any, name string, exts ...Any) Var {
	v := Define("string", require, def, name)
	v.Check = "mobile"
	return mergeExt(v, exts...)
}

func ObjectID(require bool, def Any, name string, exts ...Any) Var {
	return Define("oid", require, def, name, exts...)
}

func setFlags(v Var, flags Map) Var {
	if len(flags) == 0 {
		return v
	}
	if v.Setting == nil {
		v.Setting = Map{}
	}
	for k, val := range flags {
		v.Setting[k] = val
	}
	return v
}

func mergeExt(v Var, exts ...Any) Var {
	for _, ext := range exts {
		switch vv := ext.(type) {
		case Map:
			if v.Setting == nil {
				v.Setting = Map{}
			}
			for k, val := range vv {
				v.Setting[k] = val
			}
		case Var:
			if vv.Type != "" {
				v.Type = vv.Type
			}
			if vv.Name != "" {
				v.Name = vv.Name
			}
			if vv.Text != "" {
				v.Text = vv.Text
			}
			if vv.Default != nil {
				v.Default = vv.Default
			}
			if vv.Check != "" {
				v.Check = vv.Check
			}
			if vv.Collation != "" {
				v.Collation = vv.Collation
			}
			if vv.Comment != "" {
				v.Comment = vv.Comment
			}
			if vv.Options != nil {
				v.Options = vv.Options
			}
			if vv.Children != nil {
				v.Children = vv.Children
			}
			if vv.Encode != "" {
				v.Encode = vv.Encode
			}
			if vv.Decode != "" {
				v.Decode = vv.Decode
			}
			if vv.Empty != nil {
				v.Empty = vv.Empty
			}
			if vv.Error != nil {
				v.Error = vv.Error
			}
			if vv.Valid != nil {
				v.Valid = vv.Valid
			}
			if vv.Value != nil {
				v.Value = vv.Value
			}
			if vv.Setting != nil {
				if v.Setting == nil {
					v.Setting = Map{}
				}
				for k, val := range vv.Setting {
					v.Setting[k] = val
				}
			}
			// bool flags only elevate when explicit true is passed in ext Var.
			if vv.Required {
				v.Required = true
			}
			if vv.Nullable {
				v.Nullable = true
			}
			if vv.Unique {
				v.Unique = true
			}
		}
	}
	return v
}
