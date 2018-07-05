package common

const (
	Sql_Desc = "desc"
	Sql_Asec = "asec"
	Success  = "success"

	Time_Format = "2006-01-02 15:04:05"

	Jsctssc    = "jsctssc"
	Jslhc      = "jslhc"
	Jsctkl10gd = "jsctkl10gd"
	Jsctpk10   = "jsctpk10"
	Jsft       = "jsft"
	Jsctxyft   = "jsctxyft"
	Jsfc3d     = "jsfc3d"
	Jspl3      = "jspl3"
)

// at present, support lottery code
var LotteryCode map[string]string = map[string]string{
	"cqssc":      "重庆时时彩",
	"gd11x5":     "广东11选5",
	"bjpk10":     "北京pk10",
	"gxkl10fz":   "广西快乐10分",
	"gdkl10fz":   "广东快乐10分",
	"pl3":        "排列三",
	"ahk3":       "安徽快三",
	"fc3d":       "福彩3D",
	"xglhc":      "香港六合彩",
	"jsctssc":    "极速时时彩",
	"jslhc":      "极速六合彩",
	"jsctkl10gd": "极速快乐10分",
	"jsctpk10":   "极速pk10",
	"jsctxyft":   "极速传统幸运飞艇",
	"xyft":       "幸运飞艇",
	"jsfc3d":     "极速福彩3D",
	"jspl3":      "极速排列三",
	"tjssc":      "天津时时彩",
}
