package datastore

import (
	"context"

	"github.com/grandcolline/layered-demo/domain/entity"
	"github.com/grandcolline/layered-demo/domain/repository"
	"github.com/grandcolline/layered-demo/infrastructure/datastore/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ItemRepoImpl 商品レポジトリの実装
type ItemRepoImpl struct {
	Conn *mongo.Database
}

// NewItemRepository 商品レポジトリの実装を作成する
func NewItemRepoImpl(conn *mongo.Database) repository.ItemRepo {
	return &ItemRepoImpl{
		Conn: conn,
	}
}

// Store ユーザの新規追加する
func (repo *ItemRepoImpl) Store(e *entity.Item) (*entity.Item, error) {
	itemMdl := &model.ItemMdl{
		Name:          e.Name,
		Description:   e.Description,
		Price:         e.Price,
		SaleStartDate: e.SaleStartDate,
	}

	// データをインサートする
	collection := repo.Conn.Collection("items")
	// FIXME: context.TODO()って何か調べる？使い方あってる？
	insertResult, err := collection.InsertOne(context.TODO(), itemMdl)
	if err != nil {
		return nil, err
	}

	// 登録したオブジェクトIDの取得
	objectID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return itemMdl.ToEntity(objectID), nil
}

// FindAll 商品の全権検索をする
func (repo *ItemRepoImpl) FindAll() (*[]entity.Item, error) {
	// TODO
	return &[]entity.Item{}, nil
}

func (repo *ItemRepoImpl) FindByID(idStr string) (*entity.Item, error) {
	var result model.ItemMdl
	objID, _ := primitive.ObjectIDFromHex(idStr)
	collection := repo.Conn.Collection("items")
	// FIXME: context.TODO()って何か調べる？使い方あってる？
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.ToEntity(idStr), nil
}

func (repo *ItemRepoImpl) AddAccess(id string) error {
	// TODO
	return nil
}
