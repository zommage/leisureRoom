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

// 加入房间
func (c *RoomService) AddRoom(ctx context.Context, in *pb.AddRoomReq) (*pb.RoomNilResp, error) {
	if in == nil {
		tmpStr := "req is nil"
		Log.Errorf(tmpStr)
		return nil, errors.New(tmpStr)
	}

	// 参数校验以及组合
	row, err := CheckAddRoomReq(in)
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

	return &pb.RoomNilResp{}, nil
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
	row.MinSiteNum = gameInfo.MinSiteNum
	row.MaxSiteNum = gameInfo.MaxSiteNum
	row.SiteNum = siteNum
	row.UpdatedAt = nowTime
	row.Username = username
	row.CreatedAt = nowTime

	return row, nil
}

// 加入房间的参数校验以及组合, return int: error code
func CheckAddRoomReq(in *pb.AddRoomReq) (*models.LeisureRoom, error) {
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

	roomId := in.GetRoomId()
	err = common.NumLetter(16, 16, roomId)
	if err != nil {
		tmpStr := fmt.Sprintf("roomId=%v invalid %v", roomId, err)
		return nil, errors.New(tmpStr)
	}

	// 根据 room id 查询db
	adminRow, flag, err := models.GetRoomByRoomIdRole(roomId)
	if err != nil {
		tmpStr := fmt.Sprintf("query err: %v", err)
		return nil, errors.New(tmpStr)
	}

	if flag != true {
		tmpStr := "room not exist"
		return nil, errors.New(tmpStr)
	}

	if pwd != adminRow.Pwd {
		tmpStr := "password not match"
		return nil, errors.New(tmpStr)
	}

	nowTime := time.Now()
	row := &models.LeisureRoom{}
	row.MaxSiteNum = adminRow.MaxSiteNum
	row.GameType = adminRow.GameType
	row.CreatedAt = nowTime
	row.MinSiteNum = adminRow.MinSiteNum
	row.Pwd = adminRow.Pwd
	row.Role = 2
	row.RoomID = roomId
	row.RoomName = adminRow.RoomName
	row.SiteNum = adminRow.SiteNum
	row.UpdatedAt = nowTime
	row.Username = username

	return row, nil
}
