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

	"context"

	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	casbinpb "github.com/unistack-org/casbin-micro/casbinpb"
)

type Server struct {
	enforcerMap map[int]*casbin.Enforcer
	adapterMap  map[int]persist.Adapter
}

func NewServer() *Server {
	s := Server{}

	s.enforcerMap = map[int]*casbin.Enforcer{}
	s.adapterMap = map[int]persist.Adapter{}

	return &s
}

func (s *Server) getEnforcer(handle int) (*casbin.Enforcer, error) {
	if _, ok := s.enforcerMap[handle]; ok {
		return s.enforcerMap[handle], nil
	} else {
		return nil, errors.New("enforcer not found")
	}
}

func (s *Server) getAdapter(handle int) (persist.Adapter, error) {
	if _, ok := s.adapterMap[handle]; ok {
		return s.adapterMap[handle], nil
	} else {
		return nil, errors.New("adapter not found")
	}
}

func (s *Server) addEnforcer(enf *casbin.Enforcer) int {
	cnt := len(s.enforcerMap)
	s.enforcerMap[cnt] = enf
	return cnt
}

func (s *Server) addAdapter(a persist.Adapter) int {
	cnt := len(s.adapterMap)
	s.adapterMap[cnt] = a
	return cnt
}

func (s *Server) NewEnforcer(ctx context.Context, req *casbinpb.NewEnforcerRequest, rsp *casbinpb.NewEnforcerReply) error {
	var a persist.Adapter
	var enf *casbin.Enforcer
	var err error

	if req.AdapterHandle != -1 {
		a, err = s.getAdapter(int(req.AdapterHandle))
		if err != nil {
			return err
		}
	}

	m, err := model.NewModelFromString(req.ModelText)
	if err != nil {
		return err
	}

	if a == nil {
		enf, err = casbin.NewEnforcer(m)
	} else {
		enf, err = casbin.NewEnforcer(m, a)
	}

	if err != nil {
		return err
	}

	h := s.addEnforcer(enf)

	rsp.Handler = int32(h)

	return nil
}

func (s *Server) NewAdapter(ctx context.Context, req *casbinpb.NewAdapterRequest, rsp *casbinpb.NewAdapterReply) error {
	a, err := newAdapter(req)
	if err != nil {
		return err
	}

	h := s.addAdapter(a)

	rsp.Handler = int32(h)

	return nil
}

func (s *Server) Enforce(ctx context.Context, req *casbinpb.EnforceRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	params := make([]interface{}, 0, len(req.Params))

	m := enf.GetModel()["m"]["m"]
	sourceValue := m.Value

	for index := range req.Params {
		params = append(params, parseAbacParam(req.Params[index], m))
	}

	res, err := enf.Enforce(params...)
	if err != nil {
		return err
	}

	rsp.Res = res
	m.Value = sourceValue

	return nil
}

func (s *Server) LoadPolicy(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.EmptyReply) error {
	enf, err := s.getEnforcer(int(req.Handler))
	if err != nil {
		return err
	}

	if err = enf.LoadPolicy(); err != nil {
		return err
	}

	return nil
}

func (s *Server) SavePolicy(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.EmptyReply) error {
	enf, err := s.getEnforcer(int(req.Handler))
	if err != nil {
		return err
	}

	if err = enf.SavePolicy(); err != nil {
		return err
	}

	return nil
}
