package controller

import (
	"encoding/json"
	"net/http"

	"github.com/grandcolline/layered-demo/application/controller/request"
	"github.com/grandcolline/layered-demo/application/controller/response"
	"github.com/grandcolline/layered-demo/domain/repository"
	"github.com/grandcolline/layered-demo/domain/service"

	"github.com/go-chi/chi"
)

// ItemCtl 商品コントローラ
type ItemCtl struct {
	ItemSvc *service.ItemSvc
}

// NewItemCtl 商品コントローラの作成
func NewItemCtl(repo repository.ItemRepo) *ItemCtl {
	return &ItemCtl{
		ItemSvc: service.NewItemSvc(repo),
	}
}

// Add は新規商品の追加をする
func (ctl *ItemCtl) Add(r *http.Request) (int, []byte) {

	// POSTのデータを読み取る
	var req request.ItemReq
	json.NewDecoder(r.Body).Decode(&req)

	// FIXME: 必須チェック
	// FIXME: バリデーションチェック

	// エンティティに詰める
	e := req.ToEntity()

	// 商品登録のビジネスロジック（商品サービス）を実行
	res, err := ctl.ItemSvc.Add(e)
	if err != nil {
		// FIXME: ログを出す
		return response.ServerErrorRender()
	}

	// レスポンスを返す
	return response.ItemRender(res)
}

// List は商品一覧を取得する
func (ctl *ItemCtl) List(r *http.Request) (int, []byte) {

	// 商品全件検索のビジネスロジック（商品サービス）を実行
	res, err := ctl.ItemSvc.FindAll()
	if err != nil {
		// FIXME: ログを出す
		return response.ServerErrorRender()
	}

	// レスポンスを返す
	return response.ItemListRender(res)
}

// FindByID は商品をIDで検索する
func (ctl *ItemCtl) FindByID(r *http.Request) (int, []byte) {

	// POSTのデータを読み取る
	id := chi.URLParam(r, "ID")

	// TODO: 商品検索のビジネスロジック（商品サービス）を実行

	// TODO: レスポンスを返す
	return 200, []byte("まだ未実装（id:" + id + "）")
}
