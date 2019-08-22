package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
)

// BindJSON BindJSON
func BindJSON(data io.Reader, dest interface{}) error {
	value := reflect.ValueOf(dest)

	if value.Kind() != reflect.Ptr {
		return errors.New("BindJSON not a pointer")
	}

	decoder := json.NewDecoder(data)
	decoder.UseNumber()
	if err := decoder.Decode(dest); err != nil {
		return err
	}

	return nil
}

// APICalcSign api 签名规则 md5key为签名参数的key，appendkey是salt
func APICalcSign(mReq map[string]interface{}, appendkey, md5key string) string {

	//fmt.Println("========STEP3, 在键值对的最后加上key=API_KEY========")
	//STEP1, 在键值对的最后加上key=API_KEY
	var _buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	if appendkey != "" {
		_buffer.WriteString(appendkey)
	}

	//fmt.Println("========STEP 2, 对key进行升序排序.========")
	//fmt.Println("微信支付签名计算, API KEY:", key)
	//STEP 2, 对key进行升序排序.
	_sortedKeys := make([]string, 0)
	for k := range mReq {
		_sortedKeys = append(_sortedKeys, k)
	}

	sort.Strings(_sortedKeys)

	//fmt.Println("========STEP3, 对key=value的键值对用&连接起来，略过空值========")
	//STEP3, 对key=value的键值对用&连接起来，略过空值

	for _, _k := range _sortedKeys {
		//fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		_value := fmt.Sprintf("%v", mReq[_k])
		if _k != md5key {
			_buffer.WriteString(_k)
			_buffer.WriteString("=")
			_buffer.WriteString(_value)
			_buffer.WriteString("&")
		}
	}

	// remove lasted &
	_buf := make([]byte, _buffer.Len()-1)
	_buffer.Read(_buf)

	//fmt.Println("========STEP4, 进行MD5签名并且将所有字符转为大写.========")
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	_md5Ctx := md5.New()
	_md5Ctx.Write(_buf)
	_cipherStr := _md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(_cipherStr))

	return upperSign
}
