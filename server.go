package main

import (
	"net/http"

	"github.com/grandcolline/layered-demo/application/controller"
	"github.com/grandcolline/layered-demo/infrastructure/datastore"

	"github.com/go-chi/chi"
)

// serve サーバの起動
func serve() {
	// 設定の取得
	var conf serverConf
	conf.init()

	// ルーティング設定
	r := chi.NewRouter()
	r.Mount("/health", healthRouter())
	r.Mount("/items", itemRouter())

	http.ListenAndServe(":"+conf.Port, r)
}

// healthRouter はヘルスチェック用サブルータ.
// `/health` でルーティングされる
func healthRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	return r
}

// itemRouter は商品サービス用サブルータ.
// `/items` でルーティングされる
func itemRouter() http.Handler {
	// 商品レジストリの作成
	conn := mysqlConnect()
	itemRepo := datastore.NewItemRepoImpl(conn)

	// 商品コントローラの作成
	// logger := logger.NewLogger(conf.LogLevel)
	itemCtl := controller.NewItemCtl(itemRepo)

	// ルーティング
	r := chi.NewRouter()
	r.Get("/", handler(itemCtl.List).serve)
	r.Get("/{ID}", handler(itemCtl.FindByID).serve)
	r.Post("/", handler(itemCtl.Add).serve)

	return r
}

// handler はコントローラの抽象化
type handler func(*http.Request) (int, []byte)

// serve はハンドラが実際に書き込み処理をする部分をラッピングする
func (h handler) serve(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// panic対応
		if rv := recover(); rv != nil {
			// FIXME: ログの出力
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}()

	// ハンドラの実行
	status, res := h(r)

	// 整形をしてレスポンスを作る
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)

	return
}
