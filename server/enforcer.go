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
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	casbinpb "github.com/unistack-org/casbin-micro/casbinpb"
)

type Server struct {
	enf *casbin.Enforcer
}

func NewServer(m model.Model, a persist.Adapter) (*Server, error) {
	enf, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	s := &Server{enf: enf}
	return s, nil
}

func (s *Server) Enforce(ctx context.Context, req *casbinpb.EnforceRequest, rsp *casbinpb.BoolReply) error {
	params := make([]interface{}, 0, len(req.Params))

	m := s.enf.GetModel()["m"]["m"]
	sourceValue := m.Value

	for index := range req.Params {
		params = append(params, parseAbacParam(req.Params[index], m))
	}

	res, err := s.enf.Enforce(params...)
	if err != nil {
		return err
	}

	rsp.Res = res
	m.Value = sourceValue

	return nil
}

func (s *Server) LoadPolicy(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.EmptyReply) error {
	return s.enf.LoadPolicy()
}

func (s *Server) SavePolicy(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.EmptyReply) error {
	return s.enf.SavePolicy()
}
