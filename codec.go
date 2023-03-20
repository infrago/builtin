package builtin

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/speps/go-hashids/v2"

	. "github.com/infrago/base"
	"github.com/infrago/infra"

	"github.com/BurntSushi/toml"
)

var (
	errInvalidData = errors.New("Invalid data.")

	jsonCoder     = jsoniter.ConfigCompatibleWithStandardLibrary
	textCoder     = base64.NewEncoding(infra.TextAlphabet())
	digitCoder, _ = hashids.NewWithData(&hashids.HashIDData{
		Alphabet: infra.DigitAlphabet(), Salt: infra.DigitSalt(), MinLength: infra.DigitLength(),
	})
)

func init() {

	gob.Register(time.Now())
	gob.Register(Map{})
	gob.Register([]Map{})
	gob.Register([]Any{})

	infra.Register(infra.JSON, infra.Codec{
		Name: "JSON编解码", Text: "JSON编解码",
		Encode: func(value Any) (Any, error) {
			return jsonCoder.Marshal(value)
		},
		Decode: func(data Any, value Any) (Any, error) {
			if bytes, ok := data.([]byte); ok {
				err := jsonCoder.Unmarshal(bytes, value)
				if err != nil {
					return nil, err
				}
				return value, nil
			}
			return nil, errInvalidData
		},
	}, false)

	infra.Register(infra.XML, infra.Codec{
		Name: "XML编解码", Text: "XML编解码",
		Encode: func(value Any) (Any, error) {
			return xml.Marshal(value)
		},
		Decode: func(data Any, value Any) (Any, error) {
			if dataBytes, ok := data.([]byte); ok {
				err := xml.Unmarshal(dataBytes, value)
				if err != nil {
					return nil, err
				}
				return value, nil
			}
			return nil, errInvalidData
		},
	}, false)

	infra.Register(infra.GOB, infra.Codec{
		Name: "GOB编解码", Text: "GOB编解码",
		Encode: func(value Any) (Any, error) {
			var buffer bytes.Buffer
			encoder := gob.NewEncoder(&buffer)
			err := encoder.Encode(value)
			if err != nil {
				return nil, err
			}
			return buffer.Bytes(), nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			if dataBytes, ok := data.([]byte); ok {
				buffer := bytes.NewReader(dataBytes)
				decoder := gob.NewDecoder(buffer)
				err := decoder.Decode(value)
				if err != nil {
					return nil, err
				}

				return value, nil
			}
			return nil, errInvalidData
		},
	}, false)

	infra.Register(infra.TOML, infra.Codec{
		Name: "toml编解码", Text: "toml编解码",
		Encode: func(value Any) (Any, error) {
			var buffer bytes.Buffer
			encoder := toml.NewEncoder(&buffer)
			err := encoder.Encode(value)
			if err != nil {
				return nil, err
			}
			return buffer.Bytes(), nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			if cont, ok := data.([]byte); ok {
				_, err := toml.Decode(string(cont), value)
				if err != nil {
					return nil, err
				}
				return value, nil
			}
			return nil, errInvalidData
		},
	}, false)

	infra.Register("base64", infra.Codec{
		Alias: []string{"base64std"},
		Name:  "BASE64加解密", Text: "BASE64加解密",
		Encode: func(value Any) (Any, error) {
			text := anyToString(value)
			return base64.StdEncoding.EncodeToString([]byte(text)), nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(data)
			bytes, err := base64.URLEncoding.DecodeString(text)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		},
	}, false)

	infra.Register("base64url", infra.Codec{
		Name: "BASE64url加解密", Text: "BASE64url加解密",
		Encode: func(value Any) (Any, error) {
			text := anyToString(value)
			return base64.StdEncoding.EncodeToString([]byte(text)), nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(data)
			bytes, err := base64.URLEncoding.DecodeString(text)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		},
	}, false)

	infra.Register(infra.TEXT, infra.Codec{
		Name: "文本加密", Text: "文本加密，自定义字符表的base64编码，字典：" + infra.TextAlphabet(),
		Encode: func(value Any) (Any, error) {
			var bytes []byte
			if vvs, ok := value.([]byte); ok {
				bytes = vvs
			} else {
				bytes = []byte(anyToString(value))
			}

			text := textCoder.EncodeToString(bytes)
			return text, nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			var text string
			if vvs, ok := value.(string); ok {
				text = vvs
			} else {
				text = anyToString(data)
			}
			return textCoder.DecodeString(text)
		},
	}, false)
	infra.Register(infra.TEXTS, infra.Codec{
		Name: "文本数组加密", Text: "文本数组加密，自定义字符表的base64编码，字典：" + infra.TextAlphabet(),
		Encode: func(value Any) (Any, error) {
			text := ""
			if vvs, ok := value.(string); ok {
				text = vvs
			} else if vvs, ok := value.([]string); ok {
				text = strings.Join(vvs, "\n")
			} else {
				text = anyToString(value)
			}
			return textCoder.EncodeToString([]byte(text)), nil
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(data)
			bytes, err := textCoder.DecodeString(text)
			if err != nil {
				return nil, err
			}
			return strings.Split(string(bytes), "\n"), nil
		},
	}, false)

	infra.Register(infra.DIGIT, infra.Codec{
		Name: "数字加密", Text: "数字加密",
		Encode: func(value Any) (Any, error) {
			num := int64(0)
			if vv, ok := value.(int); ok {
				num = int64(vv)
			} else if vv, ok := value.(int64); ok {
				num = int64(vv)
			} else if vv, ok := value.(string); ok {
				if v, e := strconv.ParseInt(vv, 10, 64); e == nil {
					num = v
				} else {
					return "", errors.New("无效数字")
				}
			} else {
				return "", errors.New("无效数字")
			}
			return digitCoder.EncodeInt64([]int64{num})
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(data)
			digits, err := digitCoder.DecodeInt64WithError(text)
			if err != nil {
				return nil, err
			}
			if len(digits) == 0 {
				return nil, errors.New("无效结果")
			}
			return digits[0], nil
		},
	}, false)

	infra.Register(infra.DIGITS, infra.Codec{
		Name: "数字数组加密", Text: "数字数组加密",
		Encode: func(value Any) (Any, error) {
			nums := []int64{}
			if vv, ok := value.(int); ok {
				nums = append(nums, int64(vv))
			} else if vv, ok := value.(int64); ok {
				nums = append(nums, vv)
			} else if vvs, ok := value.([]int); ok {
				for _, num := range vvs {
					nums = append(nums, int64(num))
				}
			} else if vvs, ok := value.([]int64); ok {
				for _, num := range vvs {
					nums = append(nums, num)
				}
			} else if vv, ok := value.(string); ok {
				if v, e := strconv.ParseInt(vv, 10, 64); e == nil {
					nums = append(nums, int64(v))
				} else {
					return "", errors.New("无效数字")
				}
			} else {
				return "", errors.New("无效数字")
			}
			return digitCoder.EncodeInt64(nums)
		},
		Decode: func(data Any, value Any) (Any, error) {
			text := anyToString(data)
			return digitCoder.DecodeInt64WithError(text)
		},
	}, false)

}
