package room

import (
	"context"
	"errors"
	"fmt"

	. "github.com/zommage/leisureRoom/logs"
	pb "github.com/zommage/leisureRoom/proto"
)

type RoomService struct {
}

// 创建房间
func (c *RoomService) CreateRoom(ctx context.Context, in *pb.NewRoomReq) (*pb.NewRoomResp, error) {
	if in == nil {
		tmpStr := fmt.Sprintf("new room req is nil")
		Log.Errorf(tmpStr)
		return nil, errors.New(tmpStr)
	}

	return nil, nil
}
