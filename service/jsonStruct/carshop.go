package jsonStruct

type CarshopTmp struct {
	Data struct {
		Items []struct {
			Chid int `json:"chid"`
			//名称
			Name    string `json:"name"`
			Type    int    `json:"type"`
			Subtype int    `json:"subtype"` //`2` 小时收费
			Time    int    `json:"time"`
			Sid     string `json:"sid"`
			Policy  struct {
				Caps         int `json:"caps"`         //每天多次限额
				DurationFix  int `json:"durationFix"`  //固定时长
				DurationFree int `json:"durationFree"` //免费时长
				DurationUnit int `json:"durationUnit"` //步长
				Percaps      int `json:"percaps"`      //单日最高收费
				PriceFix     int `json:"priceFix"`     //固定收费
				PriceUnit    int `json:"priceUnit"`    //步长收费
			} `json:"policy"`
			Policy2 interface{} `json:"policy2"`
		} `json:"items"`
		Total int `json:"total"`
	} `json:"data"`
	Msg     string `json:"msg"`
	Result  int    `json:"result"`
	Traceid string `json:"traceid"`
}
