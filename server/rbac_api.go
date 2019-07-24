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

// GetRolesForUser gets the roles that a user has.
func (s *Server) GetRolesForUser(ctx context.Context, req *casbinpb.UserRoleRequest, rsp *casbinpb.ArrayReply) error {
	res, err := s.enf.GetModel()["g"]["g"].RM.GetRoles(req.User)
	if err != nil {
		return err
	}

	rsp.Array = res

	return nil
}

// GetUsersForRole gets the users that has a role.
func (s *Server) GetUsersForRole(ctx context.Context, req *casbinpb.UserRoleRequest, rsp *casbinpb.ArrayReply) error {
	res, err := s.enf.GetModel()["g"]["g"].RM.GetUsers(req.User)
	if err != nil {
		return err
	}

	rsp.Array = res
	return nil
}

// HasRoleForUser determines whether a user has a role.
func (s *Server) HasRoleForUser(ctx context.Context, req *casbinpb.UserRoleRequest, rsp *casbinpb.BoolReply) error {
	roles, err := s.enf.GetRolesForUser(req.User)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r == req.Role {
			rsp.Res = true
			break
		}
	}

	return nil
}

// AddRoleForUser adds a role for a user.
// Returns false if the user already has the role (aka not affected).
func (s *Server) AddRoleForUser(ctx context.Context, req *casbinpb.UserRoleRequest, rsp *casbinpb.BoolReply) error {
	res, err := s.enf.AddGroupingPolicy(req.User, req.Role)
	if err != nil {
		return err
	}

	rsp.Res = res

	return nil
}

// DeleteRoleForUser deletes a role for a user.
// Returns false if the user does not have the role (aka not affected).
func (s *Server) DeleteRoleForUser(ctx context.Context, req *casbinpb.UserRoleRequest, rsp *casbinpb.BoolReply) error {
	res, err := s.enf.RemoveGroupingPolicy(req.User, req.Role)
	if err != nil {
		return err
	}

	rsp.Res = res

	return nil
}

// DeleteRolesForUser deletes all roles for a user.
// Returns false if the user does not have any roles (aka not affected).
func (s *Server) DeleteRolesForUser(ctx context.Context, req *casbinpb.UserRoleRequest, rsp *casbinpb.BoolReply) error {
	res, err := s.enf.RemoveFilteredGroupingPolicy(0, req.User)
	if err != nil {
		return err
	}

	rsp.Res = res

	return nil
}

// DeleteUser deletes a user.
// Returns false if the user does not exist (aka not affected).
func (s *Server) DeleteUser(ctx context.Context, req *casbinpb.UserRoleRequest, rsp *casbinpb.BoolReply) error {
	res, err := s.enf.RemoveFilteredGroupingPolicy(0, req.User)
	if err != nil {
		return err
	}

	rsp.Res = res

	return nil
}

// DeleteRole deletes a role.
func (s *Server) DeleteRole(ctx context.Context, req *casbinpb.UserRoleRequest, rsp *casbinpb.EmptyReply) error {
	res, err := s.enf.RemoveFilteredGroupingPolicy(1, req.Role)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove filtered group policy %v", req.Role)
	}

	res, err = s.enf.RemoveFilteredPolicy(0, req.Role)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove filtered policy %v", req.Role)
	}

	return nil
}

// DeletePermission deletes a permission.
// Returns false if the permission does not exist (aka not affected).
func (s *Server) DeletePermission(ctx context.Context, req *casbinpb.PermissionRequest, rsp *casbinpb.BoolReply) error {
	res, err := s.enf.RemoveFilteredPolicy(1, req.Permissions...)
	if err != nil {
		return err
	}

	rsp.Res = res

	return nil
}

// AddPermissionForUser adds a permission for a user or role.
// Returns false if the user or role already has the permission (aka not affected).
func (s *Server) AddPermissionForUser(ctx context.Context, req *casbinpb.PermissionRequest, rsp *casbinpb.BoolReply) error {
	res, err := s.enf.AddPolicy(s.convertPermissions(req.User, req.Permissions...)...)
	if err != nil {
		return err
	}

	rsp.Res = res

	return nil
}

// DeletePermissionForUser deletes a permission for a user or role.
// Returns false if the user or role does not have the permission (aka not affected).
func (s *Server) DeletePermissionForUser(ctx context.Context, req *casbinpb.PermissionRequest, rsp *casbinpb.BoolReply) error {
	res, err := s.enf.RemovePolicy(s.convertPermissions(req.User, req.Permissions...)...)
	if err != nil {
		return err
	}

	rsp.Res = res

	return nil
}

// DeletePermissionsForUser deletes permissions for a user or role.
// Returns false if the user or role does not have any permissions (aka not affected).
func (s *Server) DeletePermissionsForUser(ctx context.Context, req *casbinpb.PermissionRequest, rsp *casbinpb.BoolReply) error {
	res, err := s.enf.RemoveFilteredPolicy(0, req.User)
	if err != nil {
		return err
	} else if !res {
		return fmt.Errorf("casbin unable to remove filtered policy %v", req.User)
	}

	rsp.Res = res

	return nil
}

// GetPermissionsForUser gets permissions for a user or role.
func (s *Server) GetPermissionsForUser(ctx context.Context, req *casbinpb.PermissionRequest, rsp *casbinpb.Array2DReply) error {
	rsp = s.wrapPlainPolicy(s.enf.GetFilteredPolicy(0, req.User))

	return nil
}

// HasPermissionForUser determines whether a user has a permission.
func (s *Server) HasPermissionForUser(ctx context.Context, req *casbinpb.PermissionRequest, rsp *casbinpb.BoolReply) error {
	rsp.Res = s.enf.HasPolicy(s.convertPermissions(req.User, req.Permissions...)...)

	return nil
}

func (s *Server) convertPermissions(user string, permissions ...string) []interface{} {
	params := make([]interface{}, 0, len(permissions)+1)
	params = append(params, user)
	for _, perm := range permissions {
		params = append(params, perm)
	}

	return params
}
