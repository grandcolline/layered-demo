package repository

import "github.com/grandcolline/layered-demo/domain/entity"

// ItemRepo 商品レポジトリポート
type ItemRepo interface {
	// Store 新規商品の追加
	Store(*entity.Item) (*entity.Item, error)
	// FindAll 商品の全権検索
	FindAll() (*[]entity.Item, error)
	// TODO: ここにインターフェースを定義
}
