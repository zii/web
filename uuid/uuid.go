// 生成全局唯一ID
package uuid

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

var ErrorInitUUID = fmt.Errorf("ErrorInitUUID")

func raise(err error) {
	if err != nil {
		panic(err)
	}
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func lower8BitPrivateIP() (uint8, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint8(ip[3]), nil
}

func InitUUID() error {
	st, err := time.Parse("2006-01-02", "2019-03-12")
	raise(err)
	sf = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: st, // 固定:项目开始日期
		MachineID: func() (uint16, error) {
			ip3, err := lower8BitPrivateIP()
			raise(err)
			id := (uint16(ip3) << 8) ^ uint16(os.Getpid())
			return id, nil
		}, // nil:以内网ip
	})
	if sf == nil {
		return ErrorInitUUID
	}
	return nil
}

// 桶的个数必须是奇数才会分布均匀
func NextID() int64 {
	v, err := sf.NextID()
	raise(err)
	return int64(v & 0x7fffffffffffffff)
}

func NextIDString() string {
	id := NextID()
	return strconv.FormatInt(id, 36)
}
