package model

import (
	"time"

	"github.com/grandcolline/layered-demo/domain/entity"
)

// ItemMdl 商品モデルのデータ構造
type ItemMdl struct {
	Name          string
	Description   string
	Price         uint32
	SaleStartDate time.Time
	AccessCount   int
}

// ToEntity 商品リクエストを商品エンティティに詰めて返却します
func (mdl *ItemMdl) ToEntity(id string) *entity.Item {
	return &entity.Item{
		ID:            id,
		Name:          mdl.Name,
		Description:   mdl.Description,
		Price:         mdl.Price,
		SaleStartDate: mdl.SaleStartDate,
	}
}
