package service

import (
	"time"

	"github.com/grandcolline/layered-demo/domain/entity"
	"github.com/grandcolline/layered-demo/domain/repository"
)

// ItemSvc 商品サービス
// 商品サービスは、商品に関するビジネスロジックを持ちます.
type ItemSvc struct {
	repo repository.ItemRepo
}

// NewItemInteractor は商品サービスの作成を行う
func NewItemSvc(repo repository.ItemRepo) *ItemSvc {
	return &ItemSvc{
		repo: repo,
	}
}

// Add は商品を新規追加する
func (svc *ItemSvc) Add(e *entity.Item) (*entity.Item, error) {
	var item *entity.Item

	// 商品の追加を行う
	item, err := svc.repo.Store(e)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// FindAll はすべての商品を検索します
func (svc *ItemSvc) FindAll() (*[]entity.Item, error) {
	var items *[]entity.Item

	items, err := svc.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}

// TODO: ここにドメインロジックをかく
