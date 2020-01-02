//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package proxy

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/fatih/structs"
	logr "github.com/sirupsen/logrus"
)

func caculateAKSN(url string, req interface{}) (string, error) {
	// delete sn element because 'sn' is not needed to be caculated
	m := structs.Map(req)
	delete(m, "sn")
	logr.Debugln("m:", m)
	// get all keys and sort them
	keys := []string{}
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Sort(sort.StringSlice(keys))

	// generate query string
	queryString := ""
	s := structs.New(req)
	logr.Debugln("keys:", keys)
	for index, key := range keys {
		field := s.Field(key)
		// get json tag name
		tagValue := field.Tag("json")
		tags := strings.Split(tagValue, ",")
		logr.Debugln("tagvalue:", tagValue)
		logr.Debugln("tags", tags)
		if len(tags) != 2 {
			return "", errors.New("invalid json tag")
		}
		fieldName := tags[0]
		val, err := getFieldValueString(field)
		if err != nil {
			return "", err
		}
		// get field value
		queryString = fieldName + "=" + val
		if index < len(keys)-1 {
			queryString += "&"
		}
	}
	logr.Debugf("query string is '%s'", queryString)

	h := md5.New()
	h.Write([]byte(queryString))
	cipherStr := h.Sum(nil)
	sn := hex.EncodeToString(cipherStr)
	return sn, nil
}

func getFieldValueString(field *structs.Field) (string, error) {
	switch field.Kind() {
	case reflect.String:
		return field.Value().(string), nil
	case reflect.Int:
		return fmt.Sprintf("%d", (field.Value().(int))), nil
	case reflect.Float64:
		return fmt.Sprintf("%f", (field.Value().(float64))), nil
	default:
		return "", fmt.Errorf("invalid filed type '%s'", field.Kind())
	}
}
