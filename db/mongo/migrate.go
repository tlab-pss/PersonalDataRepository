package mongo

import (
	"context"
	"github.com/yuuis/PersonalDataRepository/models/big_category"
	"github.com/yuuis/PersonalDataRepository/models/small_category"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"strconv"
	"time"
)

func Seed(c *mongo.Client) error {
	_, err := c.Database("pss").Collection("big_category").InsertMany(context.TODO(), bigCategories())

	if err != nil {
		return err
	}

	_, err = c.Database("pss").Collection("small_category").InsertMany(context.TODO(), smallCategories(bigCategories()))

	if err != nil {
		return err
	}

	return nil
}

func bigCategories() []interface{} {
	i := 0
	bc := []big_category.BigCategory{
		{
			ID:        incrementID(&i),
			Name:      "コマース",
			CreatedAt: time.Time{},
		},
		{
			ID:        incrementID(&i),
			Name:      "グルメ",
			CreatedAt: time.Time{},
		},
		{
			ID:        incrementID(&i),
			Name:      "天気",
			CreatedAt: time.Time{},
		},
		{
			ID:        incrementID(&i),
			Name:      "マップ",
			CreatedAt: time.Time{},
		},
		{
			ID:        incrementID(&i),
			Name:      "メール",
			CreatedAt: time.Time{},
		},
		{
			ID:        incrementID(&i),
			Name:      "ミュージック",
			CreatedAt: time.Time{},
		},
		{
			ID:        incrementID(&i),
			Name:      "メッセージ",
			CreatedAt: time.Time{},
		},
		{
			ID:        incrementID(&i),
			Name:      "検索",
			CreatedAt: time.Time{},
		},
		{
			ID:        incrementID(&i),
			Name:      "翻訳",
			CreatedAt: time.Time{},
		},
		{
			ID:        incrementID(&i),
			Name:      "ニュース",
			CreatedAt: time.Time{},
		},
	}

	var s []interface{}

	for _, c := range bc {
		s = append(s, reflect.ValueOf(c).Interface())
	}

	return s
}

func smallCategories(bc []interface{}) []interface{} {
	i := 0
	g, _ := bc[1].(big_category.BigCategory)

	// todo: とりあえず、グルメのもののみ
	sc := []small_category.SmallCategory{
		{
			ID:            incrementID(&i),
			Name:          "居酒屋",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "ダイニングバー・バル",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "創作料理",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "和食",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "洋食",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "イタリアン・フレンチ",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "中華",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "焼肉・ホルモン",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "韓国料理",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          " アジア・エスニック料理",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "各国料理",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "カラオケ・パーティ",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "バー・カクテル",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "ラーメン",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "お好み焼き・もんじゃ",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          "カフェ・スイーツ",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
		{
			ID:            incrementID(&i),
			Name:          " その他グルメ",
			BigCategoryId: g.ID,
			CreatedAt:     time.Time{},
		},
	}

	var s []interface{}

	for _, c := range sc {
		s = append(s, reflect.ValueOf(c).Interface())
	}

	return s
}

func incrementID(i *int) string {
	*i = *i + 1
	return strconv.Itoa(*i)
}
