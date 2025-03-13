package excel

import (
	"fmt"
	"log"
	"testing"
)

var (
	floatFilePath  = "./testdata/float.xlsx"
	floatSheetName = "Sheet1"
)

type FloatItem struct {
	CoinReward   int `json:"coin_reward" xlsx:"column(coin_reward)"`
	ChargeFactor int `json:"charge_factor" xlsx:"column(charge_factor)"`
}

func TestFloat(t *testing.T) {
	conn := NewConnecter()

	// see the Advancd.suffix sheet in simple.xlsx
	err := conn.Open(floatFilePath)
	if err != nil {
		t.Error(err)
	}

	rd, err := conn.NewReader(conn.GetSheetNames()[0])
	if err != nil {
		t.Error(err)
	}
	for rd.Next() {
		var s FloatItem
		if err = rd.Read(&s); err != nil {
			fmt.Println(err)
			return
		}
		log.Println(s.CoinReward)
	}

}
