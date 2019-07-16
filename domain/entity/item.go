package entity

import "time"

// Item 商品エンティティ
type Item struct {
	ID            string    // 商品ID
	Name          string    // 商品名
	Description   string    // 商品詳細説明
	Price         uint32    // 価格
	SaleStartDate time.Time // 売り出し開始日
	AccessCount   int       // アクセス数
}
