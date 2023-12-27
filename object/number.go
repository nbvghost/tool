package object

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
	"strings"

	"strconv"
)

// ParseFloat /*转化为float64*/
func ParseFloat(value interface{}) float64 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case int:
		return float64(value.(int))
	case int8:
		return float64(value.(int8))
	case int32:
		return float64(value.(int32))
	case int64:
		return float64(value.(int64))
	case uint:
		return float64(value.(uint))
	case uint8:
		return float64(value.(uint8))
	case uint32:
		return float64(value.(uint32))
	case uint64:
		return float64(value.(uint64))
	case float32:
		return float64(value.(float32))
	case float64:
		return value.(float64)
	case string:
		numberStr := value.(string)
		if strings.EqualFold(numberStr, "") {
			return 0
		}
		numb, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			log.Println(err.Error())
		}
		return numb
	case []uint8:
		var pi float64
		buf := bytes.NewReader(value.([]byte))
		err := binary.Read(buf, binary.LittleEndian, &pi)
		if err != nil {
			log.Println(err.Error())
		}
	default:
		raw := ParseRaw(reflect.ValueOf(value))
		if raw == nil {
			log.Println(errors.New("未支持的数据类型：" + fmt.Sprintf("%v", reflect.TypeOf(value))))
		} else {
			return ParseFloat(raw)
		}

	}

	return 0
}
func ParseString(value interface{}) string {
	if value == nil {
		return ""
	}
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int))
	case int8:
		return strconv.Itoa(int(value.(int8)))
	case int32:
		return strconv.Itoa(int(value.(int32)))
	case int64:
		return strconv.Itoa(int(value.(int64)))
	case uint:
		return strconv.Itoa(int(value.(uint)))
	case uint8:
		return strconv.Itoa(int(value.(uint8)))
	case uint32:
		return strconv.Itoa(int(value.(uint32)))
	case uint64:
		u := value.(uint64)
		if u > math.MaxInt64 {
			return strconv.FormatUint(math.MaxInt64, 10)
		}
		return strconv.Itoa(int(u))
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case string:
		numberStr := value.(string)
		return numberStr
	default:
		raw := ParseRaw(reflect.ValueOf(value))
		if raw == nil {
			log.Println(errors.New("未支持的数据类型：" + fmt.Sprintf("%v", reflect.TypeOf(value))))
		} else {
			return ParseString(raw)
		}

	}

	return ""
}
func Convert(value reflect.Value, types reflect.Type) reflect.Value {
	vv := value.Convert(types)
	return vv
}
func ParseRaw(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.Struct:
		f := value.MethodByName("String")
		if f.IsValid() {
			results := f.Call([]reflect.Value{})
			if len(results) != 1 {
				log.Println("String方法没有返回值")
				return nil
			}
			return results[0].String()
		} else {
			return nil
		}
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		vv := value.Convert(reflect.TypeOf(float64(0)))
		return vv.Interface().(float64)
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		vv := value.Convert(reflect.TypeOf(int(0)))
		return vv.Interface().(int)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		vv := value.Convert(reflect.TypeOf(uint(0)))
		return vv.Interface().(uint)
	case reflect.String:
		vv := value.Convert(reflect.TypeOf(""))
		return vv.Interface().(string)
	case reflect.Slice:
		return value.Interface()
	case reflect.Ptr:
		return ParseRaw(value.Elem())
	default:
		panic(errors.New(fmt.Sprintf("不支持类型：%v", value.Kind())))
	}

	return nil
}
func ParseInt(value interface{}) int {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case int:
		return value.(int)
	case int8:
		return int(value.(int8))
	case int32:
		return int(value.(int32))
	case int64:
		return int(value.(int64))
	case uint:
		return int(value.(uint))
	case uint8:
		return int(value.(uint8))
	case uint32:
		return int(value.(uint32))
	case uint64:
		u := value.(uint64)
		if u > math.MaxInt64 {
			return math.MaxInt64
		}
		return int(u)
	case float32:
		return int(value.(float32))
	case float64:
		u := value.(float64)
		if u > math.MaxInt64 {
			return math.MaxInt64
		}
		return int(u)
	case string:
		numberStr := value.(string)
		if strings.EqualFold(numberStr, "") {
			return 0
		}
		numb, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Println(err.Error())
		}
		return numb
	default:
		raw := ParseRaw(reflect.ValueOf(value))
		if raw == nil {
			log.Println(errors.New("未支持的数据类型：" + fmt.Sprintf("%v", reflect.TypeOf(value))))
		} else {
			return ParseInt(raw)
		}

	}
	return 0
}
func ParseUint(value interface{}) uint {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case int:
		return uint(value.(int))
	case int8:
		return uint(value.(int8))
	case int32:
		return uint(value.(int32))
	case int64:
		return uint(value.(int64))
	case uint:
		return uint(value.(uint))
	case uint8:
		return uint(value.(uint8))
	case uint32:
		return uint(value.(uint32))
	case uint64:
		return uint(value.(uint64))
	case float32:
		return uint(value.(float32))
	case float64:
		u := value.(float64)
		if u > math.MaxUint64 {
			return math.MaxUint64
		}
		return uint(u)
	case string:
		numberStr := value.(string)
		if strings.EqualFold(numberStr, "") {
			return 0
		}
		numb, err := strconv.ParseUint(numberStr, 10, 64)
		if err != nil {
			log.Println(err.Error())
		}
		return uint(numb)
	default:
		vv := reflect.ValueOf(value).Convert(reflect.TypeOf(uint(0)))
		v, ok := vv.Interface().(uint)
		if ok {
			return ParseUint(v)
		}
		log.Println(errors.New("未支持的数据类型：" + fmt.Sprintf("%v", reflect.TypeOf(value))))
	}
	return 0
}
func ParseBool(value interface{}) bool {
	if value == nil {
		return false
	}
	switch value.(type) {
	case int:
		return (value.(int)) > 0
	case int8:
		return (value.(int8)) > 0
	case int32:
		return (value.(int32)) > 0
	case int64:
		return (value.(int64)) > 0
	case uint:
		return (value.(uint)) > 0
	case uint8:
		return (value.(uint8)) > 0
	case uint32:
		return (value.(uint32)) > 0
	case uint64:
		return (value.(uint64)) > 0
	case float32:
		return (value.(float32)) > 0
	case float64:
		return (value.(float64)) > 0
	case bool:
		return value.(bool)
	case string:
		numberStr := value.(string)
		if strings.EqualFold(numberStr, "") {
			return false
		}
		numb, err := strconv.ParseBool(numberStr)
		if err != nil {
			log.Println(err.Error())
		}
		return numb
	default:
		vv := reflect.ValueOf(value).Convert(reflect.TypeOf(false))
		v, ok := vv.Interface().(uint)
		if ok {
			return ParseBool(v)
		}
		log.Println(errors.New("未支持的数据类型：" + fmt.Sprintf("%v", reflect.TypeOf(value))))
	}
	return false
}
