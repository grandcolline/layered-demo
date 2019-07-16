package response

import (
	"encoding/json"
	"net/http"

	"github.com/grandcolline/layered-demo/domain/entity"
)

// item 商品レスポンス
type itemRes struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         uint32 `json:"price"`
	SaleStartDate string `json:"sale_start_date"`
	// アクセス数はレスポンスには含まない
}

// set 商品エンティティから商品レスポンスを取得する
func (res *itemRes) set(e *entity.Item) {
	res.ID = e.ID
	res.Name = e.Name
	res.Description = e.Description
	res.Price = e.Price
	res.SaleStartDate = e.SaleStartDate.Format("2006-01-02")
}

// ItemRender 商品エンティティからレスポンスのJsonを生成する
func ItemRender(e *entity.Item) (int, []byte) {
	var res itemRes
	res.set(e)
	json, err := json.Marshal(res)
	if err != nil {
		// FIXME: ログを出す
		return ServerErrorRender()
	}
	return http.StatusOK, json
}

// ItemListRender 商品エンティティのリストからレスポンスのJsonを生成する
func ItemListRender(es *[]entity.Item) (int, []byte) {
	var items []itemRes
	for _, e := range *es {
		var res itemRes
		res.set(&e)
		items = append(items, res)
	}
	json, err := json.Marshal(items)
	if err != nil {
		// FIXME: ログを出す
		return ServerErrorRender()
	}
	return http.StatusOK, json
}
