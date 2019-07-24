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
	"fmt"

	casbinpb "github.com/unistack-org/casbin-micro/casbinpb"
)

func (s *Server) wrapPlainPolicy(policy [][]string) *casbinpb.Array2DReply {
	if len(policy) == 0 {
		return &casbinpb.Array2DReply{}
	}

	policyReply := &casbinpb.Array2DReply{}
	policyReply.D2 = make([]*casbinpb.Array2DReplyD, len(policy))
	for e := range policy {
		policyReply.D2[e] = &casbinpb.Array2DReplyD{D1: policy[e]}
	}

	return policyReply
}

// GetAllSubjects gets the list of subjects that show up in the current policy.
func (s *Server) GetAllSubjects(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.ArrayReply) error {
	return s.GetAllNamedSubjects(ctx, &casbinpb.SimpleGetRequest{EnforcerHandler: req.Handler, PType: "p"}, rsp)
}

// GetAllNamedSubjects gets the list of subjects that show up in the current named policy.
func (s *Server) GetAllNamedSubjects(ctx context.Context, req *casbinpb.SimpleGetRequest, rsp *casbinpb.ArrayReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	rsp.Array = enf.GetModel().GetValuesForFieldInPolicy("p", req.PType, 0)

	return nil
}

// GetAllObjects gets the list of objects that show up in the current policy.
func (s *Server) GetAllObjects(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.ArrayReply) error {
	return s.GetAllNamedObjects(ctx, &casbinpb.SimpleGetRequest{EnforcerHandler: req.Handler, PType: "p"}, rsp)
}

// GetAllNamedObjects gets the list of objects that show up in the current named policy.
func (s *Server) GetAllNamedObjects(ctx context.Context, req *casbinpb.SimpleGetRequest, rsp *casbinpb.ArrayReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	rsp.Array = enf.GetModel().GetValuesForFieldInPolicy("p", req.PType, 1)

	return nil
}

// GetAllActions gets the list of actions that show up in the current policy.
func (s *Server) GetAllActions(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.ArrayReply) error {
	return s.GetAllNamedActions(ctx, &casbinpb.SimpleGetRequest{EnforcerHandler: req.Handler, PType: "p"}, rsp)
}

// GetAllNamedActions gets the list of actions that show up in the current named policy.
func (s *Server) GetAllNamedActions(ctx context.Context, req *casbinpb.SimpleGetRequest, rsp *casbinpb.ArrayReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	rsp.Array = enf.GetModel().GetValuesForFieldInPolicy("p", req.PType, 2)

	return nil
}

// GetAllRoles gets the list of roles that show up in the current policy.
func (s *Server) GetAllRoles(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.ArrayReply) error {
	return s.GetAllNamedRoles(ctx, &casbinpb.SimpleGetRequest{EnforcerHandler: req.Handler, PType: "g"}, rsp)
}

// GetAllNamedRoles gets the list of roles that show up in the current named policy.
func (s *Server) GetAllNamedRoles(ctx context.Context, req *casbinpb.SimpleGetRequest, rsp *casbinpb.ArrayReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	rsp.Array = enf.GetModel().GetValuesForFieldInPolicy("g", req.PType, 1)

	return nil
}

// GetPolicy gets all the authorization rules in the policy.
func (s *Server) GetPolicy(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.Array2DReply) error {
	return s.GetNamedPolicy(ctx, &casbinpb.PolicyRequest{EnforcerHandler: req.Handler, PType: "p"}, rsp)
}

// GetNamedPolicy gets all the authorization rules in the named policy.
func (s *Server) GetNamedPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.Array2DReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res := enf.GetModel().GetPolicy("p", req.PType)
	rsp = s.wrapPlainPolicy(res)

	return nil
}

// GetFilteredPolicy gets all the authorization rules in the policy, field filters can be specified.
func (s *Server) GetFilteredPolicy(ctx context.Context, req *casbinpb.FilteredPolicyRequest, rsp *casbinpb.Array2DReply) error {
	req.PType = "p"

	return s.GetFilteredNamedPolicy(ctx, req, rsp)
}

// GetFilteredNamedPolicy gets all the authorization rules in the named policy, field filters can be specified.
func (s *Server) GetFilteredNamedPolicy(ctx context.Context, req *casbinpb.FilteredPolicyRequest, rsp *casbinpb.Array2DReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res := enf.GetModel().GetFilteredPolicy("p", req.PType, int(req.FieldIndex), req.FieldValues...)
	rsp = s.wrapPlainPolicy(res)

	return nil
}

// GetGroupingPolicy gets all the role inheritance rules in the policy.
func (s *Server) GetGroupingPolicy(ctx context.Context, req *casbinpb.EmptyRequest, rsp *casbinpb.Array2DReply) error {
	return s.GetNamedGroupingPolicy(ctx, &casbinpb.PolicyRequest{EnforcerHandler: req.Handler, PType: "g"}, rsp)
}

// GetNamedGroupingPolicy gets all the role inheritance rules in the policy.
func (s *Server) GetNamedGroupingPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.Array2DReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res := enf.GetModel().GetPolicy("g", req.PType)
	rsp = s.wrapPlainPolicy(res)

	return nil
}

// GetFilteredGroupingPolicy gets all the role inheritance rules in the policy, field filters can be specified.
func (s *Server) GetFilteredGroupingPolicy(ctx context.Context, req *casbinpb.FilteredPolicyRequest, rsp *casbinpb.Array2DReply) error {
	req.PType = "g"

	return s.GetFilteredNamedGroupingPolicy(ctx, req, rsp)
}

// GetFilteredNamedGroupingPolicy gets all the role inheritance rules in the policy, field filters can be specified.
func (s *Server) GetFilteredNamedGroupingPolicy(ctx context.Context, req *casbinpb.FilteredPolicyRequest, rsp *casbinpb.Array2DReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res := enf.GetModel().GetFilteredPolicy("g", req.PType, int(req.FieldIndex), req.FieldValues...)
	rsp = s.wrapPlainPolicy(res)

	return nil
}

// HasPolicy determines whether an authorization rule exists.
func (s *Server) HasPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	return s.HasNamedPolicy(ctx, req, rsp)
}

// HasNamedPolicy determines whether a named authorization rule exists.
func (s *Server) HasNamedPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	rsp.Res = enf.GetModel().HasPolicy("p", req.PType, req.Params)

	return nil
}

// HasGroupingPolicy determines whether a role inheritance rule exists.
func (s *Server) HasGroupingPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	req.PType = "g"

	return s.HasNamedGroupingPolicy(ctx, req, rsp)
}

// HasNamedGroupingPolicy determines whether a named role inheritance rule exists.
func (s *Server) HasNamedGroupingPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	rsp.Res = enf.GetModel().HasPolicy("g", req.PType, req.Params)

	return nil
}

func (s *Server) AddPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	req.PType = "p"

	return s.AddNamedPolicy(ctx, req, rsp)
}

func (s *Server) AddNamedPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := enf.AddNamedPolicy(req.PType, req.Params)
	if err != nil {
		return err
	}

	rsp.Res = res

	return nil
}

func (s *Server) RemovePolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	e, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := e.RemovePolicy(req.Params)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove policy %v", req.Params)
	}

	rsp.Res = res

	return nil
}

func (s *Server) RemoveNamedPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := enf.RemoveNamedPolicy(req.PType, req.Params)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove named policy %v %v", req.PType, req.Params)
	}

	rsp.Res = res

	return nil
}

// RemoveFilteredPolicy removes an authorization rule from the current policy, field filters can be specified.
func (s *Server) RemoveFilteredPolicy(ctx context.Context, req *casbinpb.FilteredPolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := enf.RemoveFilteredNamedPolicy("p", int(req.FieldIndex), req.FieldValues...)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove filtered named policy")
	}

	rsp.Res = res

	return nil
}

// RemoveFilteredNamedPolicy removes an authorization rule from the current named policy, field filters can be specified.
func (s *Server) RemoveFilteredNamedPolicy(ctx context.Context, req *casbinpb.FilteredPolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := enf.RemoveFilteredNamedPolicy(req.PType, int(req.FieldIndex), req.FieldValues...)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove filtered named policy")
	}

	rsp.Res = res

	return nil
}

// AddGroupingPolicy adds a role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (s *Server) AddGroupingPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	req.PType = "g"

	return s.AddNamedGroupingPolicy(ctx, req, rsp)
}

// AddNamedGroupingPolicy adds a named role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (s *Server) AddNamedGroupingPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := enf.AddNamedGroupingPolicy(req.PType, req.Params)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to add named grouping policy")
	}

	rsp.Res = res

	return nil
}

// RemoveGroupingPolicy removes a role inheritance rule from the current policy.
func (s *Server) RemoveGroupingPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := enf.RemoveNamedGroupingPolicy("g", req.Params)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove named grouping policy")
	}

	rsp.Res = res

	return nil
}

// RemoveNamedGroupingPolicy removes a role inheritance rule from the current named policy.
func (s *Server) RemoveNamedGroupingPolicy(ctx context.Context, req *casbinpb.PolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := enf.RemoveNamedGroupingPolicy(req.PType, req.Params)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove named grouping policy")
	}

	rsp.Res = res

	return nil
}

// RemoveFilteredGroupingPolicy removes a role inheritance rule from the current policy, field filters can be specified.
func (s *Server) RemoveFilteredGroupingPolicy(ctx context.Context, req *casbinpb.FilteredPolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := enf.RemoveFilteredNamedGroupingPolicy("g", int(req.FieldIndex), req.FieldValues...)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove filtered named grouping policy")
	}

	rsp.Res = res

	return nil
}

// RemoveFilteredNamedGroupingPolicy removes a role inheritance rule from the current named policy, field filters can be specified.
func (s *Server) RemoveFilteredNamedGroupingPolicy(ctx context.Context, req *casbinpb.FilteredPolicyRequest, rsp *casbinpb.BoolReply) error {
	enf, err := s.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		return err
	}

	res, err := enf.RemoveFilteredNamedGroupingPolicy(req.PType, int(req.FieldIndex), req.FieldValues...)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove filtered named grouping policy")
	}

	rsp.Res = res

	return nil
}
