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

	// view := computeAccessLevel(roles[0])

	for _, role := range roles {
		views = append(views, computeAccessLevel(role))
	}

	// fmt.Printf("%+v\n", view)

	return views, nil
}

func computeAccessLevel(role entities.Role) entities.RoleView {
	rv := entities.RoleView{
		ID:       utils.IntToString(role.ID),
		Name:     role.Name,
		MemberOf: role.MemberOf,
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

	access := entities.RoleAccessLevelLimited

	if role.IsSuper {
		access = entities.RoleAccessLevelAdmin
	} else {

		for _, r := range role.MemberOf {
			if _, ok := entities.RoleDescriptions[r]; ok {
				access = entities.RoleAccessLevelAdmin
			}
		}

		// Elevated
		if access != entities.RoleAccessLevelAdmin {
			if role.CanCreateRole {
				access = entities.RoleAccessLevelElevated
			}
			for _, r := range role.MemberOf {
				if r == "pg_write_all_data" {
					access = entities.RoleAccessLevelElevated
				}
			}
		}

		// Standard
		if access != entities.RoleAccessLevelAdmin && access != entities.RoleAccessLevelElevated {
			if role.CanCreateDB {
				access = entities.RoleAccessLevelStandard
			}
			for _, r := range role.MemberOf {
				if r == "pg_read_all_data" || r == "pg_monitor" {
					access = entities.RoleAccessLevelStandard
				}
			}
		}
	}

	rv.AccessLevel = access
	rv.Capabilities = buildCapabilitiesDescriptions(role)
	return rv
}

func buildCapabilitiesDescriptions(role entities.Role) []string {
	var descriptions []string

	// Flags
	if role.IsSuper {
		descriptions = append(descriptions, entities.RoleDescriptions["superuser"])
	}
	if role.Replication {
		descriptions = append(descriptions, entities.RoleDescriptions["replication"])
	}
	if role.CanCreateDB {
		descriptions = append(descriptions, entities.RoleDescriptions["createdb"])
	}
	if role.CanCreateDB {
		descriptions = append(descriptions, entities.RoleDescriptions["createrole"])
	}
	if role.CanLogin {
		descriptions = append(descriptions, entities.RoleDescriptions["rolcanlogin"])
	}

	// Membership
	for _, r := range role.MemberOf {
		if desc, ok := entities.RoleDescriptions[r]; ok {
			descriptions = append(descriptions, desc)
		} else {
			descriptions = append(descriptions, r)
		}
	}

	return descriptions
}
