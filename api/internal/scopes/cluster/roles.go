package cluster

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
	"dashboard/api/internal/utils"
	"errors"
)

func (c *Cluster) Roles(ctx context.Context) ([]entities.RoleView, error) {

	roles, err := c.storage.Roles(ctx)
	if err != nil {
		c.logger.ErrorContext(ctx, "get roles", "error", err)
		return nil, errors.New("failed to get cluster roles")
	}

	views := make([]entities.RoleView, 0, len(roles))

	for _, role := range roles {
		views = append(views, computeRoleAccessLevel(role))
	}

	return views, nil
}

func buildMembership(role entities.Role) []entities.RoleMembership {
	var res []entities.RoleMembership

	for _, r := range role.MemberOf {
		desc, ok := entities.RoleDescriptions[r]
		if !ok {
			desc = r
		}

		res = append(res, entities.RoleMembership{
			Name:        r,
			Description: desc,
		})
	}

	return res
}

func computeRoleAccessLevel(role entities.Role) entities.RoleView {
	rv := entities.RoleView{
		ID:          utils.IntToString(role.ID),
		Name:        role.Name,
		Membership:  buildMembership(role),
		IsGroupRole: !role.CanLogin,
	}

	var flags []entities.RoleFlag

	if role.IsSuper {
		flags = append(flags, entities.RoleFlagSuperuser)
	}
	if role.CanLogin {
		flags = append(flags, entities.RoleFlagLogin)
	}
	if role.CanCreateRole {
		flags = append(flags, entities.RoleFlagCreateRole)
	}
	if role.CanCreateDB {
		flags = append(flags, entities.RoleFlagCreateDB)
	}
	if role.Replication {
		flags = append(flags, entities.RoleFlagReplication)
	}

	rv.Flags = flags

	rv.AccessLevel = entities.RoleAccessLevelLimited

	for _, rule := range entities.AccessRules {
		if rule.Check(role) {
			rv.AccessLevel = rule.Level
			break
		}
	}

	return rv
}
