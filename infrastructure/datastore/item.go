package datastore

import (
	"strconv"

	"github.com/grandcolline/layered-demo/domain/entity"
	"github.com/grandcolline/layered-demo/domain/repository"
	"github.com/grandcolline/layered-demo/infrastructure/datastore/model"

	"github.com/jinzhu/gorm"
)

// ItemRepoImpl 商品レポジトリの実装
type ItemRepoImpl struct {
	Conn *gorm.DB
}

// NewItemRepository 商品レポジトリの実装を作成する
func NewItemRepoImpl(conn *gorm.DB) repository.ItemRepo {
	return &ItemRepoImpl{
		Conn: conn,
	}
}

// Store ユーザの新規追加する
func (repo *ItemRepoImpl) Store(e *entity.Item) (*entity.Item, error) {
	itemMdl := &model.ItemMdl{
		// FIXME: ここうまく共通化かしら
		Name:          e.Name,
		Description:   e.Description,
		Price:         e.Price,
		SaleStartDate: e.SaleStartDate,
	}

	// データをインサートする
	if err := repo.Conn.Create(itemMdl).Error; err != nil {
		return nil, err
	}

	return itemMdl.ToEntity(), nil
}

// Store ユーザの新規追加する
func (repo *ItemRepoImpl) FindAll() (*[]entity.Item, error) {

	items := []model.ItemMdl{}

	// データを検索します
	if err := repo.Conn.Find(&items).Error; err != nil {
		return nil, err
	}

	// エンティティの作成
	ent := make([]entity.Item, len(items))
	for i, item := range items {
		ent[i] = *item.ToEntity()
	}

	return &ent, nil
}

// TODO: ここに実実装を実装
