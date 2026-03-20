package cluster

import (
	"context"
	"dashboard/api/internal/scopes/cluster/entities"
	"errors"
)

func (c *Cluster) Roles(ctx context.Context) ([]entities.RoleView, error) {

	roles, err := c.storage.Roles(ctx)
	if err != nil {
		c.logger.ErrorContext(ctx, "get roles", "error", err)
		return nil, errors.New("failed to get cluster roles")
	}

	views := make([]entities.RoleView, 0, len(roles))

	// view := computeAccessLevel(roles[0])

	for _, role := range roles {
		views = append(views, computeAccessLevel(role))
	}

	// fmt.Printf("%+v\n", view)

	return views, nil
}

func computeAccessLevel(role entities.Role) entities.RoleView {
	rv := entities.RoleView{
		Name:     role.Name,
		MemberOf: role.MemberOf,
	}

	flags := []string{}
	if role.IsSuper {
		flags = append(flags, "superuser")
	}
	if role.CanLogin {
		flags = append(flags, "canLogin")
	}
	if role.CanCreateRole {
		flags = append(flags, "createRole")
	}
	if role.CanCreateDB {
		flags = append(flags, "createDb")
	}
	if role.Replication {
		flags = append(flags, "replication")
	}

	rv.Flags = flags

	reasons := []string{}
	access := entities.RoleAccessLevelLimited

	if role.IsSuper {
		access = entities.RoleAccessLevelAdmin
		reasons = append(reasons, "Superuser")
	} else {
		dangerousRoles := map[string]string{
			"pg_execute_server_program": "Can execute server programs",
			"pg_write_server_files":     "Can write server files",
			"pg_read_server_files":      "Can read server files",
		}
		for _, r := range role.MemberOf {
			if reason, ok := dangerousRoles[r]; ok {
				access = entities.RoleAccessLevelAdmin
				reasons = append(reasons, reason)
			}
		}

		// Elevated
		if access != entities.RoleAccessLevelAdmin {
			if role.CanCreateRole {
				access = entities.RoleAccessLevelElevated
				reasons = append(reasons, "Can create roles")
			}
			for _, r := range role.MemberOf {
				if r == "pg_write_all_data" {
					access = entities.RoleAccessLevelElevated
					reasons = append(reasons, "Can modify all data")
				}
			}
		}

		// Standard
		if access != entities.RoleAccessLevelAdmin && access != entities.RoleAccessLevelElevated {
			if role.CanCreateDB {
				access = entities.RoleAccessLevelStandard
				reasons = append(reasons, "Can create DBs")
			}
			for _, r := range role.MemberOf {
				if r == "pg_read_all_data" || r == "pg_monitor" {
					access = entities.RoleAccessLevelStandard
					reasons = append(reasons, "Extended read or monitoring access")
				}
			}
		}
	}

	rv.AccessLevel = access
	rv.Reasons = reasons
	return rv
}
