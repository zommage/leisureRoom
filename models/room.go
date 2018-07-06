package dbbase

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// platform site gateway users
type LeisureRoom struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT;"` // id
	Username   string    `gorm:"column:username;unique_index"`          // 用户名
	Pwd        string    `gorm:"column:pwd"`                            // 用户名密码
	Role       int       `gorm:"column:role"`                           // 角色,1: 为房间创建人, 2: 随从者
	SiteNum    int       `gorm:"column:site_num"`                       // 座位人数
	MinSiteNum int       `gorm:"column:min_site_num"`                   // 最少座位人数
	MaxSiteNum int       `gorm:"column:max_site_num"`                   // 最多座位人数
	RoomID     string    `gorm:"column:room_id"`                        // 房间号
	RoomName   string    `gorm:"column:room_name"`                      // 房间名称
	GameType   int       `gorm:"column:game_type"`                      // 游戏类型
	UpdatedAt  time.Time `gorm:"column:updated_at"`                     // 更新时间
	CreatedAt  time.Time `gorm:"column:created_at"`                     // 创建时间
}

// 查找房间
func GetRoomByRoomIdRole(roomId string) (*LeisureRoom, bool, error) {
	x := db.Dbs.Table("leisure_room").Where("room_id = ? AND role = 1", roomId)
	count := 0
	err := x.Count(&count).Error
	if err != nil {
		return nil, false, err
	}

	if count == 0 {
		return nil, false, nil
	}

	row := &LeisureRoom{}
	err = x.First(row).Error
	if err != nil {
		return nil, false, err
	}

	return row, true, nil
}
