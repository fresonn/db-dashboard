package database

type StatDirection string

const (
	StatDirectionUp   StatDirection = "up"
	StatDirectionDown StatDirection = "down"
)

type DatabaseStatsOverview struct {
	Size DatabaseSizeStaticStat `json:"size"`
}

type DatabaseSizeStaticStat struct {
	SizeBytes  int64           `json:"sizeBytes"`
	SizePretty string          `json:"sizePretty"`
	Trend      StaticStatTrend `json:"trend"`
}

type StaticStatTrend struct {
	Value     string        `json:"value"`
	Direction StatDirection `json:"direction"`
}
