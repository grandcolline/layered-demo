package request

import (
	"time"

	"github.com/grandcolline/layered-demo/domain/entity"
)

// Item 商品リクエスト
type ItemReq struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         uint32    `json:"price"`
	SaleStartDate time.Time `json:"sale_start_date"`
}

// ToEntity 商品リクエストを商品エンティティに詰めて返却します
func (req *ItemReq) ToEntity() *entity.Item {
	return &entity.Item{
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		SaleStartDate: req.SaleStartDate,
	}
}
