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
	"errors"

	"github.com/casbin/casbin/persist"
	gormadapter "github.com/casbin/gorm-adapter"
	casbinpb "github.com/unistack-org/casbin-micro/casbinpb"

	//_ "github.com/jinzhu/gorm/dialects/mssql"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var errDriverName = errors.New("currently supported DriverName: mysql | postgres | mssql")

func newAdapter(req *casbinpb.NewAdapterRequest) (persist.Adapter, error) {
	var a persist.Adapter
	supportDriverNames := [...]string{"mysql", "postgres", "mssql"}

	switch req.DriverName {
	default:
		var support = false
		for _, driverName := range supportDriverNames {
			if driverName == req.DriverName {
				support = true
				break
			}
		}
		if support {
			a = gormadapter.NewAdapter(req.DriverName, req.ConnectString, req.DbSpecified)
			break
		}
		return nil, errDriverName
	}

	return a, nil
}
