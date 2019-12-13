package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CodeParameter struct {
	Letter       string
	IsExpression bool
	IsDriverId   bool
	stringValue  string
	IsString     bool
	parsedValue  interface{}
}

func NewSimpleCodeParameter(letter string, value interface{}) *CodeParameter {
	_, isString := value.(string)
	return &CodeParameter{
		Letter:      letter,
		parsedValue: value,
		stringValue: fmt.Sprintf("%v", value),
		IsString:    isString,
	}
}

func (cp CodeParameter) String() string {
	if cp.IsString {
		return fmt.Sprintf(`%s"%s"`, cp.Letter, strings.ReplaceAll(cp.stringValue, `"`, `""`))
	} else {
		return fmt.Sprintf("%s%s", cp.Letter, cp.stringValue)
	}
}

func (cp *CodeParameter) ConvertDriverIds() error {
	if cp.IsExpression {
		return nil
	}

	d := make([]uint64, 0)
	params := strings.Split(cp.stringValue, ":")
	for _, p := range params {
		s := strings.Split(p, ".")
		if len(s) == 1 {
			u, err := strconv.ParseUint(s[0], 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("Failed to parse driver number from %s parameter", cp.Letter))
			}
			d = append(d, u)
		} else if len(s) == 2 {
			board, err := strconv.ParseUint(s[0], 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("Failed to parse board number from %s parameter", cp.Letter))
			}
			port, err := strconv.ParseUint(s[1], 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("Failed to parse driver number from %s parameter", cp.Letter))
			}
			driver := (board << 16) | port
			d = append(d, driver)
		} else {
			return errors.New(fmt.Sprintf("Driver value from %s parameter is invalid.", cp.Letter))
		}
	}

	if len(d) == 1 {
		cp.parsedValue = d[0]
	} else {
		cp.parsedValue = d
	}
	cp.IsDriverId = true

	return nil
}

func (cp *CodeParameter) AsFloat64() (float64, error) {
	if f, ok := cp.parsedValue.(float64); ok {
		return f, nil
	}
	return 0, errors.New(fmt.Sprintf("Cannot convert %s parameter to float64 (value %s)", cp.Letter, cp.stringValue))
}

func (cp *CodeParameter) AsInt64() (int64, error) {
	if i, ok := cp.parsedValue.(int64); ok {
		return i, nil
	}
	return 0, errors.New(fmt.Sprintf("Cannot convert %s parameter to int64 (value %s)", cp.Letter, cp.stringValue))
}

func (cp *CodeParameter) AsUint64() (uint64, error) {
	if u, ok := cp.parsedValue.(uint64); ok {
		return u, nil
	}
	return 0, errors.New(fmt.Sprintf("Cannot convert %s parameter to uint64 (value %s)", cp.Letter, cp.stringValue))
}

func (cp *CodeParameter) AsBool() (bool, error) {
	return strconv.ParseBool(cp.stringValue)
}

func (cp *CodeParameter) AsString() string {
	return cp.stringValue
}

func (cp *CodeParameter) AsFloat64Slice() ([]float64, error) {
	switch v := cp.parsedValue.(type) {
	case []float64:
		return v, nil
	case float64:
		return []float64{v}, nil
	case int64:
		return []float64{float64(v)}, nil
	case uint64:
		return []float64{float64(v)}, nil
	case []int64:
		fs := make([]float64, 0, len(v))
		for _, i := range v {
			fs = append(fs, float64(i))
		}
		return fs, nil
	case []uint64:
		fs := make([]float64, 0, len(v))
		for _, i := range v {
			fs = append(fs, float64(i))
		}
		return fs, nil
	}

	return nil, errors.New(fmt.Sprintf("Cannot convert %s parameter to []float64 (value %s)", cp.Letter, cp.stringValue))
}

func (cp *CodeParameter) AsInt64Slice() ([]int64, error) {
	switch v := cp.parsedValue.(type) {
	case []int64:
		return v, nil
	case float64:
		return []int64{int64(v)}, nil
	case int64:
		return []int64{int64(v)}, nil
	case uint64:
		return []int64{int64(v)}, nil
	case []float64:
		fs := make([]int64, 0, len(v))
		for _, i := range v {
			fs = append(fs, int64(i))
		}
		return fs, nil
	case []uint64:
		fs := make([]int64, 0, len(v))
		for _, i := range v {
			fs = append(fs, int64(i))
		}
		return fs, nil
	}

	return nil, errors.New(fmt.Sprintf("Cannot convert %s parameter to []int64 (value %s)", cp.Letter, cp.stringValue))
}

func (cp *CodeParameter) AsUint64Slice() ([]uint64, error) {
	switch v := cp.parsedValue.(type) {
	case []uint64:
		return v, nil
	case float64:
		return []uint64{uint64(v)}, nil
	case int64:
		return []uint64{uint64(v)}, nil
	case uint64:
		return []uint64{uint64(v)}, nil
	case []float64:
		fs := make([]uint64, 0, len(v))
		for _, i := range v {
			fs = append(fs, uint64(i))
		}
		return fs, nil
	case []int64:
		fs := make([]uint64, 0, len(v))
		for _, i := range v {
			fs = append(fs, uint64(i))
		}
		return fs, nil
	}

	return nil, errors.New(fmt.Sprintf("Cannot convert %s parameter to []uint64 (value %s)", cp.Letter, cp.stringValue))
}

func (cp *CodeParameter) Init(letter, value string, isString bool) error {
	cp.Letter = letter
	cp.stringValue = value
	cp.IsString = isString
	if cp.IsString {
		cp.parsedValue = cp.stringValue
		return nil
	}
	value = strings.TrimSpace(value)
	if value == "" {
		cp.parsedValue = 0
	} else if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
		cp.IsExpression = true
	} else if strings.Contains(value, ":") {

		split := strings.Split(value, ":")
		success := true
		// FIXME: This is ugly
		if strings.Contains(value, ".") {
			floats := make([]float64, 0)
			for _, s := range split {
				f, err := strconv.ParseFloat(s, 64)
				if err != nil {
					success = false
					break
				}
				floats = append(floats, f)
			}
			if success {
				cp.parsedValue = floats
			}
		} else {
			ints := make([]int64, 0)
			for _, s := range split {
				i, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					success = false
					break
				}
				ints = append(ints, i)
			}
			if success {
				cp.parsedValue = ints
			} else {
				success = true
				uints := make([]uint64, 0)
				for _, s := range split {
					u, err := strconv.ParseUint(s, 10, 64)
					if err != nil {
						success = false
						break
					}
					uints = append(uints, u)
				}
				if success {
					cp.parsedValue = uints
				}
			}
		}

		if !success {
			cp.parsedValue = value
		}
	} else if i, err := strconv.ParseInt(value, 10, 64); err == nil {
		cp.parsedValue = i
	} else if u, err := strconv.ParseUint(value, 10, 64); err == nil {
		cp.parsedValue = u
	} else if f, err := strconv.ParseFloat(value, 64); err == nil {
		cp.parsedValue = f
	} else {
		cp.parsedValue = value
	}
	return nil
}

func (cp *CodeParameter) MarshalJSON() ([]byte, error) {
	ss := make(map[string]interface{})
	ss["letter"] = cp.Letter
	ss["value"] = cp.stringValue
	if cp.IsString {
		ss["isString"] = 1
	} else {
		ss["isString"] = 0
	}

	return json.Marshal(ss)
}

func (cp *CodeParameter) UnmarshalJSON(data []byte) error {
	ss := make(map[string]interface{})
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}
	var letter, value string
	var isString bool
	for k, v := range ss {
		if k == "letter" {
			letter = v.(string)
		} else if k == "value" {
			value = v.(string)
		} else if k == "isString" {
			isString = v.(float64) == 1
		}
	}
	return cp.Init(letter, value, isString)
}
