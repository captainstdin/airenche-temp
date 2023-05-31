package fee

import (
	"encoding/json"
	"githun.com/captainstdin/airenche-temp/service"
	"githun.com/captainstdin/airenche-temp/service/jsonStruct"
)

func GetFreeForTmp(siteSid, secretKey string) (jsonStruct.CarshopTmp, error) {

	body, err := service.CallRemoteCurl("/carshop/tmpchargeGetList", nil)
	if err != nil {
		return jsonStruct.CarshopTmp{}, err
	}
	var CarShop jsonStruct.CarshopTmp
	jsonerr := json.Unmarshal(body, &CarShop)

	if jsonerr != nil {
		return CarShop, jsonerr
	}
	return CarShop, nil
}
