// Copyright 2018 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/casbin/casbin/model"
)

type AbacAttrList struct {
	V0      string
	V1      string
	V2      string
	V3      string
	V4      string
	V5      string
	V6      string
	V7      string
	V8      string
	V9      string
	V10     string
	nameMap map[string]string
}

func toUpperFirstChar(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func MakeABAC(obj interface{}) (string, error) {
	data, err := json.Marshal(&obj)
	if err != nil {
		return "", err
	}
	return "ABAC::" + string(data), nil
}

func resolveABAC(obj string) (AbacAttrList, error) {
	var jsonMap map[string]interface{}
	attrList := AbacAttrList{nameMap: map[string]string{}}

	err := json.Unmarshal([]byte(obj[len("ABAC::"):]), &jsonMap)
	if err != nil {
		return attrList, err
	}

	i := 0
	for k, v := range jsonMap {
		key := toUpperFirstChar(k)
		value := fmt.Sprintf("%v", v)
		attrList.nameMap[key] = "V" + strconv.Itoa(i)
		switch i {
		case 0:
			attrList.V0 = value
		case 1:
			attrList.V1 = value
		case 2:
			attrList.V2 = value
		case 3:
			attrList.V3 = value
		case 4:
			attrList.V4 = value
		case 5:
			attrList.V5 = value
		case 6:
			attrList.V6 = value
		case 7:
			attrList.V7 = value
		case 8:
			attrList.V8 = value
		case 9:
			attrList.V9 = value
		case 10:
			attrList.V10 = value
		}
		i++
	}

	return attrList, nil
}

func parseAbacParam(param string, m *model.Assertion) interface{} {
	if strings.HasPrefix(param, "ABAC::") {
		attrList, err := resolveABAC(param)
		if err != nil {
			panic(err)
		}
		for k, v := range attrList.nameMap {
			old := "." + k
			if strings.Contains(m.Value, old) {
				m.Value = strings.Replace(m.Value, old, "."+v, -1)
			}
		}

		return attrList
	} else {
		return param
	}
}
