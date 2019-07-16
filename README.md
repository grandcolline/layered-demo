# layered-demo

## 開発

```bash
# 起動
$ make up

# 停止
$ make down

# ログ確認
$ make log

# お片付け
$ make cleanup
```

## APIテスト

```bash
# 商品登録
$ curl -X POST "http://localhost:8080/items" \
    -d "{\"name\":\"商品A\",\"description\":\"販売前の商品です\",\"price\":3000,\"sale_start_date\":\"2019-08-01T00:00:00Z\"}" \
    | jq .
$ curl -X POST "http://localhost:8080/items" \
    -d "{\"name\":\"商品B\",\"description\":\"販売中の商品です\",\"price\":2999,\"sale_start_date\":\"2019-07-01T00:00:00Z\"}" \
    | jq .

# 一覧取得
$ curl -X GET "http://localhost:8080/items" | jq .
```
