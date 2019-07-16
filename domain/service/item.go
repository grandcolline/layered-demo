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

// FindByID は商品をIDで検索します
func (svc *ItemSvc) FindByID(id string) (*entity.Item, error) {

	// IDで商品をデータストアから取得
	item, err := svc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// もし、購入期間前なら金額を0円にする
	if time.Now().Before(item.SaleStartDate) {
		item.Price = 0
	}

	// アクセス履歴を登録する
	if err := svc.repo.AddAccess(item.ID); err != nil {
		return nil, err
	}

	return item, nil
}
