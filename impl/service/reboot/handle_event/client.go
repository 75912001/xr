package handle_event

import (
	"fmt"

	"github.com/75912001/xr/impl/service/reboot"

	"github.com/75912001/xr/impl/service/common/proto_head"
	"github.com/75912001/xr/lib/tcp"
)

func OnEventDisConnClient(client *tcp.Client) int {
	if !client.IsConn() {
		return 0
	}
	//TODO 业务逻辑
	return 0
}

func OnEventPacketClient(client *tcp.Client, data []byte) int {
	//TODO 业务逻辑
	return 0
}

func OnParseProtoHeadClient(data []byte, length int) int {
	if uint32(length) < proto_head.GProtoHeadLength {
		//长度不足一个包头的长度大小
		return 0
	}
	packetLength := int(proto_head.GetPacketLength(data))
	if uint32(packetLength) < proto_head.GProtoHeadLength {
		reboot.GRebootMgr.Log.Error(fmt.Sprintf("packetLength:%v", packetLength))
		return -1
	}
	if reboot.GRebootMgr.BenchMgr.Json.Base.PacketLengthMax < uint32(packetLength) {
		reboot.GRebootMgr.Log.Error(fmt.Sprintf("PacketLengthMax:%v, packetLength:%v",
			reboot.GRebootMgr.BenchMgr.Json.Base.PacketLengthMax, packetLength))
		return -1
	}

	if length < int(packetLength) {
		return 0
	}

	return packetLength
}
