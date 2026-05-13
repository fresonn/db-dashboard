package database

type StatDirection string

const (
	StatDirectionUp   StatDirection = "up"
	StatDirectionDown StatDirection = "down"
)

type DatabaseStatsOverview struct {
	ID      string                    `json:"id"`
	Name    string                    `json:"name"`
	Size    DatabaseSizeStaticStat    `json:"size"`
	Tables  DatabaseTablesStaticStat  `json:"tables"`
	Indexes DatabaseIndexesStaticStat `json:"indexes"`
}

type DatabaseSizeStaticStat struct {
	SizeBytes  int64           `json:"sizeBytes"`
	SizePretty string          `json:"sizePretty"`
	Trend      StaticStatTrend `json:"trend"`
}

type DatabaseTablesStaticStat struct {
	Total int64           `json:"total"`
	Trend StaticStatTrend `json:"trend"`
}

type DatabaseIndexesStaticStat struct {
	Total int64           `json:"total"`
	Trend StaticStatTrend `json:"trend"`
}

type StaticStatTrend struct {
	Value     string        `json:"value"`
	Direction StatDirection `json:"direction"`
}

type PostgresDbOverviewStats struct {
	ID      int64
	Name    string
	Size    int64
	Tables  int64
	Indexes int64
}
