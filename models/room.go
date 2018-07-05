package dbbase

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// platform site gateway users
type LeisureRoom struct {
	ID        int       `gorm:"column:id;primary_key;AUTO_INCREMENT;"` // id
	Username  string    `gorm:"column:username;unique_index"`          // 用户名
	Pwd       string    `gorm:"column:pwd"`                            // 用户名密码
	Role      int       `gorm:"column:role"`                           // 角色, 如: admin, normal, 判断是否为房间创建人
	SiteNum   int       `gorm:"column:site_num"`                       // 座位人数
	RoomID    string    `gorm:"column:room_id"`                        // 房间号
	RoomName  string    `gorm:"column:room_name"`                      // 房间名称
	GameType  int       `gorm:"column:game_type"`                      // 游戏类型
	UpdatedAt time.Time `gorm:"column:updated_at"`                     // 更新时间
	CreatedAt time.Time `gorm:"column:created_at"`                     // 创建时间
}
