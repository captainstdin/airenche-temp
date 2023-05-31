package sluice

import (
	"fmt"
	"githun.com/captainstdin/airenche-temp/service"
)

type ChannelInstance struct {
	Cid string //通道ID
	Sid string //车场ID
}

func (this *ChannelInstance) Open(plate string) error {
	/**
	1   蓝牌车
	2   黄牌车
	4   新能源小车
	8   新能源黄牌车
	101 警车
	102 救护车
	103 消防车
	*/
	var data = map[string]string{
		"sid":     this.Sid,
		"cid":     this.Cid,
		"carno":   plate,
		"cartype": "1",
		"regtype": "0", //0 临时车 1 注册车 2 月卡车 3 Vip车 4 黑名单车 21 共组月卡车 22 共组临时车
	}
	body, err := service.CallRemoteCurl("/carshop/MonitorChannelOpen", data)

	if err != nil {
		return err
	}

	fmt.Println(string(body))
	//json.Unmarshal(body, "")

	return nil
}
