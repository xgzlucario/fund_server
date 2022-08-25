package pro

import (
	"fmt"
	"fund/db"
	"fund/model"

	"go.mongodb.org/mongo-driver/bson"
)

// 策略1：低买高卖
func test1(arg float64) {
	trade := model.NewTrade(fmt.Sprintf("test arg:%.2f", arg))

	klineMap.Range(func(id string, k []model.Kline) {
		trade.Init()

		for i := range k {
			if k[i].WinnerRate < 2.7 && k[i].Tr < 3.5 && k[i].Pe < 33 {
				// log
				_id := bson.M{"code": id, "time": k[i].Time}
				db.Backtest.UpsertId(ctx,
					_id, bson.M{"type": "b", "close": k[i].Close, "arg": arg, "winner_rate": k[i].WinnerRate})

				trade.Buy(k[i])

			} else if k[i].WinnerRate > arg {
				// log
				_id := bson.M{"code": id, "time": k[i].Time}
				db.Backtest.UpsertId(ctx,
					_id, bson.M{"type": "s", "close": k[i].Close, "arg": arg, "winner_rate": k[i].WinnerRate})

				trade.Sell(k[i], id)
			}
		}
	})
	trade.RecordsInfo()
}

func Test1() {
	go test1(20)
	go test1(30)
	go test1(40)
	go test1(50)
	go test1(60)
}
