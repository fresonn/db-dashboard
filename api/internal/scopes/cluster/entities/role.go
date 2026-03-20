package entities

type RoleAccessLevel string

const (
	RoleAccessLevelAdmin    RoleAccessLevel = "admin"
	RoleAccessLevelStandard RoleAccessLevel = "standard"
	RoleAccessLevelElevated RoleAccessLevel = "elevated"
	RoleAccessLevelLimited  RoleAccessLevel = "limited"
)

type Role struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	CanLogin      bool     `json:"canLogin"`
	IsSuper       bool     `json:"isSuper"`
	CanCreateRole bool     `json:"canCreateRole"`
	CanCreateDB   bool     `json:"canCreateDb"`
	Replication   bool     `json:"replication"`
	MemberOf      []string `json:"memberOf"`
}

type RoleView struct {
	Name        string          `json:"name"`
	MemberOf    []string        `json:"memberOf"`
	Flags       []string        `json:"flags"`
	AccessLevel RoleAccessLevel `json:"accessLevel"`
	Reasons     []string        `json:"reasons"`
}
