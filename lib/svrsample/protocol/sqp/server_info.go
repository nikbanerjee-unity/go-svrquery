package sqp

import (
	"github.com/multiplay/go-svrquery/lib/svrsample/common"
)

// ServerInfo holds the ServerInfo chunk data.
type ServerInfo struct {
	CurrentPlayers uint16
	MaxPlayers     uint16
	ServerName     string
	GameType       string
	BuildID        string
	GameMap        string
	Port           uint16
}

// QueryStateToServerInfo converts proto.QueryState to ServerInfo.
func QueryStateToServerInfo(qs common.QueryState) ServerInfo {
	return ServerInfo{
		CurrentPlayers: uint16(qs.CurrentPlayers),
		MaxPlayers:     uint16(qs.MaxPlayers),
		ServerName:     qs.ServerName,
		GameType:       qs.GameType,
		GameMap:        qs.Map,
		Port:           qs.Port,
	}
}

// Size returns the number of bytes sqpServerInfo will use on the wire.
func (si ServerInfo) Size() uint32 {
	return uint32(
		2 + // CurrentPlayers
			2 + // MaxPlayers
			len([]byte(si.ServerName)) + 1 +
			len([]byte(si.GameType)) + 1 +
			len([]byte(si.BuildID)) + 1 +
			len([]byte(si.GameMap)) + 1 +
			2, // Port
	)
}
