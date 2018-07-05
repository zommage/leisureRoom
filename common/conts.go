package common

const (
	Sql_Desc = "desc"
	Sql_Asec = "asec"
	Success  = "success"

	Time_Format = "2006-01-02 15:04:05"
)

// 游戏的信息
type GameInfo struct {
	GameType   int // 游戏类型
	MinSiteNum int // 最少人数开始
	MaxSiteNum int // 最大人数开始
}

var (
	InvalidStr = "^%@!~·#*()+{}[]|:,;=?$\\\""

	GameCodeMap = map[string]*GameInfo{
		"paohuzi": &GameInfo{
			GameType:   1,
			MinSiteNum: 3,
			MaxSiteNum: 4,
		},
	}
)
