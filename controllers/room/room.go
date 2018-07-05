package room

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zommage/leisureRoom/common"
	. "github.com/zommage/leisureRoom/logs"
	models "github.com/zommage/leisureRoom/models"
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

	// 参数校验以及组合
	row, err := CheckNewRoomReq(in)
	if err != nil {
		tmpStr := fmt.Sprintf("invalid params: %v", err)
		Log.Errorf(tmpStr)
		return nil, errors.New(tmpStr)
	}

	err = models.InsertDbs(row)
	if err != nil {
		tmpStr := fmt.Sprintf("insert room to db err: %v", err)
		Log.Errorf(tmpStr)
		return nil, errors.New(tmpStr)
	}

	resp := &pb.NewRoomResp{}
	resp.RoomId = row.RoomID

	return resp, nil
}

// 参数组合以及校验
func CheckNewRoomReq(in *pb.NewRoomReq) (*models.LeisureRoom, error) {
	gameCode := in.GetGameCode()
	gameInfo, exist := common.GameCodeMap[gameCode]
	if !exist {
		tmpStr := "gameCode=" + gameCode + " is not exist"
		return nil, errors.New(tmpStr)
	}

	username := in.GetUsername()
	err := common.NumLetterLine(6, 12, username)
	if err != nil {
		tmpStr := fmt.Sprintf("username=%v invalid %v", username, err)
		return nil, errors.New(tmpStr)
	}

	pwd := in.GetPwd()
	if len(pwd) < 3 || len(pwd) > 20 {
		tmpStr := "pwd len invalid 3-20"
		return nil, errors.New(tmpStr)
	}

	roomName := in.GetRoomName()
	flag := strings.ContainsAny(roomName, common.InvalidStr)
	if flag != false {
		tmpStr := "roomName=" + roomName + " is invalid"
		return nil, errors.New(tmpStr)
	}

	siteNum := int(in.GetSiteNum())
	if siteNum > gameInfo.MinSiteNum || siteNum < gameInfo.MaxSiteNum {
		tmpStr := fmt.Sprintf("siteNum=%v is invalid %v-%v", siteNum, gameInfo.MinSiteNum, gameInfo.MaxSiteNum)
		return nil, errors.New(tmpStr)
	}

	nowTime := time.Now()

	row := &models.LeisureRoom{}
	row.GameType = gameInfo.GameType
	row.Pwd = pwd
	row.Role = 1
	row.RoomID = common.Get16RandId()
	row.RoomName = roomName
	row.SiteNum = siteNum
	row.UpdatedAt = nowTime
	row.Username = username
	row.CreatedAt = nowTime

	return row, nil
}
