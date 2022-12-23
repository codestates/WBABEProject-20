package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// type BodyMenu struct {
// 	Category   string //`json:"category"`
// 	Name       string //`json:"name" binding:"required"`
// 	Status     string //`json:"status"`
// 	MaxCount   string //`json:"maxCount"`
// 	CountryOf  string //`json:"countryOf"`
// 	Price      string //`json:"orice"`
// 	Spicy      string //`json:"spicy"`
// 	Popularity string //`json:"popularity"`
// 	IsDisabled string //`json:"isDisabled"`
// 	TodayMenu  string //`json:"todayMenu"`
// }

// type BodyMenu struct {
// 	Name       string             //메뉴 이름
// 	Status     string             //주문 가능 상태
// 	MaxCount   int                //판매 가능 갯수
// 	CountryOf  string             //원산지
// 	Price      int                //가격
// 	Spicy      string             //맵기
// 	Popularity int                //인기도
// 	IsDisabled bool               //판매여부
// 	TodayMenu  bool               //오늘의 추천메뉴 여부
// 	Category   string             //메뉴 카테고리
// 	CreateDate primitive.DateTime //등록일
// 	ModifyDate primitive.DateTime //수정일

// }

type BodyMenu struct {
	Name       string             `bson:"name"`       //메뉴 이름
	Status     string             `bson:"status"`     //주문 가능 상태
	MaxCount   int                `bson:"maxCount"`   //판매 가능 갯수
	CountryOf  string             `bson:"countryOf"`  //원산지
	Price      int                `bson:"price"`      //가격
	Spicy      string             `bson:"spicy"`      //맵기
	Popularity int                `bson:"popularity"` //인기도
	IsDisabled bool               `bson:"isDisabled"` //판매여부
	TodayMenu  bool               `bson:"todayMenu"`  //오늘의 추천메뉴 여부
	Category   string             `bson:"category"`   //메뉴 카테고리
	CreateDate primitive.DateTime `bson:"createDate"` //등록일
	ModifyDate primitive.DateTime `bson:"modifyDate"` //수정일
}

// func getMenuByBody(body BodyMenu) Menu {

// 	var menu Menu
// 	menu.Name = body.Name
// 	menu.Status = body.Status

// 	maxCount, err := strconv.Atoi(body.MaxCount)
// 	if err != nil {

// 	} else {
// 		menu.MaxCount = maxCount
// 	}

// 	menu.CountryOf = body.CountryOf
// 	price, err := strconv.Atoi(body.Price)
// 	if err != nil {
// 		logger.Warn(err)
// 	} else {
// 		menu.Price = price
// 	}

// 	popularity, err := strconv.Atoi(body.Popularity)
// 	if err != nil {
// 		logger.Warn(err)
// 	} else {
// 		menu.Popularity = popularity
// 	}

// 	if body.IsDisabled != "" {
// 		menu.IsDisabled = true
// 	} else {
// 		menu.IsDisabled = false
// 	}

// 	if body.TodayMenu != "" {
// 		menu.TodayMenu = true
// 	} else {
// 		menu.TodayMenu = false
// 	}

// 	menu.Category = body.Category

// 	return menu
// }
