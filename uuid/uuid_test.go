package uuid

import (
	"fmt"
	"testing"
	"time"

	"github.com/zii/web"

	cmap "github.com/orcaman/concurrent-map"

	"github.com/sony/sonyflake"
)

func aTest1(t *testing.T) {
	st, err := time.Parse("2006-01-02", "2019-03-12")
	web.Raise(err)
	sf1 := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: st, // 固定:项目开始日期
		MachineID: func() (uint16, error) {
			return 1, nil
		}, // nil:以内网ip
	})
	sf2 := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: st, // 固定:项目开始日期
		MachineID: func() (uint16, error) {
			return 1, nil
		}, // nil:以内网ip
	})
	nextID := func(sf *sonyflake.Sonyflake) int64 {
		v, err := sf.NextID()
		web.Raise(err)
		return int64(v & 0x00000fffffffffff)
	}
	var set = cmap.New() // 在线用户信息表
	run := func(sf *sonyflake.Sonyflake) {
		for {
			id := fmt.Sprintf("%d", nextID(sf))
			if set.Has(id) {
				panic("那你?")
			}
			set.Set(id, 1)
			time.Sleep(1 * time.Millisecond)
		}
	}
	go run(sf1)
	go run(sf2)
	time.Sleep(10 * time.Second)
	fmt.Println("OK", set.Count())
}

func Test2(t *testing.T) {
	InitUUID()
	id := NextID() & 0x7ffffff
	fmt.Println(id)
}
