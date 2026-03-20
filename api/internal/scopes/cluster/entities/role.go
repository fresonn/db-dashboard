package entities

type RoleAccessLevel string

const (
	RoleAccessLevelAdmin    RoleAccessLevel = "admin"
	RoleAccessLevelStandard RoleAccessLevel = "standard"
	RoleAccessLevelElevated RoleAccessLevel = "elevated"
	RoleAccessLevelLimited  RoleAccessLevel = "limited"
)

type RoleFlag string

const (
	RoleFlagSuperuser   RoleFlag = "superuser"
	RoleFlagLogin       RoleFlag = "login"
	RoleFlagCreateRole  RoleFlag = "create_role"
	RoleFlagCreateDB    RoleFlag = "create_db"
	RoleFlagReplication RoleFlag = "replication"
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
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	MemberOf     []string        `json:"memberOf"`
	Flags        []RoleFlag      `json:"flags"`
	Capabilities []string        `json:"capabilities"`
	AccessLevel  RoleAccessLevel `json:"accessLevel"`
}

var RoleDescriptions = map[string]string{
	// Data access
	"pg_read_all_data":  "Can read all data",
	"pg_write_all_data": "Can modify all data",

	// Monitoring
	"pg_monitor":           "Can access monitoring stats",
	"pg_read_all_settings": "Can read all configuration settings",
	"pg_read_all_stats":    "Can read all statistics",
	"pg_stat_scan_tables":  "Can scan tables for statistics",

	// Server access
	"pg_read_server_files":      "Can read server files",
	"pg_write_server_files":     "Can write server files",
	"pg_execute_server_program": "Can execute server programs",

	// Admin actions
	"pg_signal_backend": "Can terminate backend processes",
	"pg_checkpoint":     "Can trigger checkpoints",

	// Replication
	"pg_create_subscription": "Can create logical replication subscriptions",

	// Flags
	"superuser":   "Full access to cluster",
	"replication": "Can manage replication",
	"createdb":    "Can create databases",
	"createrole":  "Can create roles",
	"rolcanlogin": "Can log in",
}
