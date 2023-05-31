package fee

import (
	"githun.com/captainstdin/airenche-temp/service"
	"testing"
)

func TestGetFreeForTmp(t *testing.T) {

	err := service.LoadEnv()
	if err != nil {
		t.Error(err)
		return
	}

	tmp, err := GetFreeForTmp(service.EnvSite.Sid, service.EnvSite.SecretKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("检测到临时车有的收费配置方案：%d个", tmp.Data.Total)
	for index, item := range tmp.Data.Items {
		t.Logf("编号%d： \n 免费时长：%d分，\n 起步收费：%d元，\n 收费最高：%d元，\n 收费步长：%d分",
			index,
			item.Policy.DurationFree,
			item.Policy.PriceFix,
			item.Policy.Percaps, item.Policy.DurationUnit)
	}

}
