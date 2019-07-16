package model

import (
	"strconv"
	"time"

	"github.com/grandcolline/layered-demo/domain/entity"
	"github.com/jinzhu/gorm"
)

// ItemMdl 商品モデルのデータ構造
type ItemMdl struct {
	gorm.Model
	Name          string    `gorm:"NOT NULL;type:varchar(255) COMMENT '商品名'"`
	Description   string    `gorm:"NOT NULL;type:varchar(255) COMMENT '商品詳細説明'"`
	Price         uint32    `gorm:"NOT NULL;type:INT(10) COMMENT '商品価格'"`
	SaleStartDate time.Time `gorm:"type:Timestamp COMMENT '売り出し開始日'"`
	AccessCount   int       `gorm:"NOT NULL;type:INT(10) COMMENT 'アクセス数'"`
}

func (*ItemMdl) TableName() string {
	return "items"
}

// ToEntity 商品リクエストを商品エンティティに詰めて返却します
func (mdl *ItemMdl) ToEntity() *entity.Item {
	return &entity.Item{
		ID:            strconv.FormatUint(uint64(mdl.ID), 10),
		Name:          mdl.Name,
		Description:   mdl.Description,
		Price:         mdl.Price,
		SaleStartDate: mdl.SaleStartDate,
	}
}
