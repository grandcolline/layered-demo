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

func (repo *ItemRepoImpl) FindByID(id string) (*entity.Item, error) {
	var item entity.Item
	i, _ := strconv.Atoi(id)
	if err := repo.Conn.First(&item, uint32(i)).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (repo *ItemRepoImpl) AddAccess(id string) error {
	i, _ := strconv.Atoi(id)
	err := repo.Conn.First(&model.ItemMdl{}, uint32(i)).UpdateColumn("access_count", gorm.Expr("access_count + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}
