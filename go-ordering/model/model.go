package model

//model.go : db에 접속해 데이터를 핸들링, 결과 전달
import (
	"WBABEProject-20/go-ordering/logger"
	"context"
	"time"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client     *mongo.Client
	colPersons *mongo.Collection
}

type Orderer struct {
	OrdererID  string             `bson:"ordererID"`  //주문자ID
	Name       string             `bson:"name"`       //주문자 이름
	Address    string             `bson:"address"`    //주문자 주소
	Phone      int                `bson:"phone"`      //주문자 폰번호
	OrderCount int                `bson:"orderCount"` //주문 숫자
	CreateDate primitive.DateTime `bson:"createDate"` //등록일
	ModifyDate primitive.DateTime `bson:"modifyDate"` //수정일
}

type OrdererMenuLink struct {
	OrderNO        string             `bson:"orderNO"`        //주문번호
	MenuID         string             `bson:"menuID"`         //메뉴ID
	MenuName       string             `bson:"menuName"`       //메뉴ID
	OrdererID      int                `bson:"ordererID"`      //주문자ID
	Grade          int                `bson:"grade"`          //평점
	Comment        string             `bson:"comment"`        //후기
	Status         string             `bson:"status"`         //주문상태
	OrdererAddress string             `bson:"ordererAddress"` //주문자 주소
	OrdererPhone   int                `bson:"ordererPhone"`   //주문자 폰번호
	CreateDate     primitive.DateTime `bson:"createDate"`     //등록일
	ModifyDate     primitive.DateTime `bson:"modifyDate"`     //수정일
}

type Seller struct {
	Name       string             `bson:"name"`       //판매자이름
	Address    string             `bson:"address"`    //판매자 주소
	Phone      int                `bson:"phone"`      //판매자 전화번호
	SellCount  int                `bson:"sellCount"`  //판매자 판매갯수
	CreateDate primitive.DateTime `bson:"createDate"` //등록일
	ModifyDate primitive.DateTime `bson:"modifyDate"` //수정일
}

type Menu struct {
	MenuID     string             `bson:"menuID"`     //메뉴 이름
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

func NewModel() (*Model, error) {
	fmt.Println("Model.NewModel start")

	r := &Model{}

	var err error
	mgUrl := "mongodb://127.0.0.1:27017"
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		logger.Info("ConnectModel.")
		db := r.client.Database("go-ready")
		r.colPersons = db.Collection("tMenu")
	}
	return r, nil
}

// 메뉴 생성
func (p *Model) CreateMenu(menu Menu) Menu {
	logger.Info("[model.CreateMenu Param] ", menu)
	fmt.Println("[model.CreateMenu Param] ", menu)

	menu.CreateDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))
	menu.ModifyDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))

	result, err := p.colPersons.InsertOne(context.TODO(), menu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted with ID: %s\n", result.InsertedID)

	return menu
}

// 메뉴 수정
func (p *Model) UpdateMenu(menu Menu) Menu {
	fmt.Println("[model.UpdateMenu Param] ", menu)
	fmt.Println("[model.UpdateMenu menu.Name] ", menu.Name)
	fmt.Println("[model.UpdateMenu menu.Status] ", menu.Status)

	filter := bson.D{{"name", menu.Name}}

	// e, err := json.Marshal(menu)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("[model.UpdateMenu e] ", e)

	update := bson.M{
		"$set": bson.M{
			"status":     menu.Status,
			"maxCount":   menu.MaxCount,
			"countryOf":  menu.CountryOf,
			"price":      menu.Price,
			"spicy":      menu.Spicy,
			"popularity": menu.Popularity,
			"isDisabled": menu.IsDisabled,
			"todayMenu":  menu.TodayMenu,
			"category":   menu.Category,
			"createDate": menu.CreateDate,
			"modifyDate": menu.ModifyDate,
		},
	}

	fmt.Println("[model.UpdateMenu filter] ", filter)

	result, err := p.colPersons.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[model.UpdateMenu result] ", result)
	fmt.Println("[model.UpdateMenu result.ModifiedCount] ", result.ModifiedCount)

	return menu
}

// 메뉴 삭제
func (p *Model) DeleteMenu(menu Menu) Menu {

	fmt.Println("[model.DeleteMenu Param] ", menu)
	fmt.Println("[model.DeleteMenu menu.Name] ", menu.Name)

	filter := bson.D{{"name", menu.Name}}
	result, err := p.colPersons.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents Deleted: %v\n", result.DeletedCount)

	return menu
}

// 메뉴 검색
func (p *Model) SearchMenu(menu Menu) []Menu {
	fmt.Println("[model.SearchMenu Param] ", menu)

	var filter bson.D
	if menu.Name != "" {
		filter = append(filter, bson.E{"name", menu.Name})
		//filter = bson.D{{"name", menu.Name}}
	}
	if menu.Status != "" {
		filter = append(filter, bson.E{"status", menu.Status})
	}

	cursor, err := p.colPersons.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("[model.SearchMenu err] ", err)
		panic(err)
	}

	var menus []Menu
	if err = cursor.All(context.TODO(), &menus); err != nil {
		fmt.Println("[model.SearchMenu err] ", err)
		panic(err)
	}

	fmt.Println("[model.SearchMenu menus] ", menus)

	return menus
}

// 주문 생성
func (p *Model) NewOrder(omLink OrdererMenuLink) OrdererMenuLink {
	logger.Info("[model.NewOrder Param] ", omLink)
	fmt.Println("[model.NewOrder Param] ", omLink)

	omLink.CreateDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))
	omLink.ModifyDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))

	result, err := p.colPersons.InsertOne(context.TODO(), omLink)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted with ID: %s\n", result.InsertedID)

	return omLink
}

// 주문 수정/취소
func (p *Model) ChangeOrder(omLink OrdererMenuLink) OrdererMenuLink {
	logger.Info("[model.ChangeOrder Param] ", omLink)
	fmt.Println("[model.ChangeOrder Param] ", omLink)

	//omLink.OrderNO
	filter := bson.D{{"menuID", omLink.MenuID}, {"ordererID", omLink.OrdererID}}

	update := bson.M{
		"$set": bson.M{
			"status":     omLink.Grade,
			"maxCount":   omLink.Comment,
			"countryOf":  omLink.Status,
			"price":      omLink.OrdererAddress,
			"spicy":      omLink.OrdererPhone,
			"modifyDate": omLink.ModifyDate,
		},
	}

	fmt.Println("[model.ChangeOrder filter] ", filter)

	result, err := p.colPersons.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[model.ChangeOrder result] ", result)
	fmt.Println("[model.ChangeOrder result.ModifiedCount] ", result.ModifiedCount)

	return omLink
}

// 주문 검색
func (p *Model) SearchOrder(omLink OrdererMenuLink) []OrdererMenuLink {
	fmt.Println("[model.SearchOrder Param] ", omLink)

	filter := bson.D{{"name", omLink.OrdererID}}
	if omLink.MenuName != "" {
		filter = append(filter, bson.E{"menuName", omLink.MenuName})
	}

	cursor, err := p.colPersons.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("[model.SearchOrder err] ", err)
		panic(err)
	}

	var omLinks []OrdererMenuLink
	if err = cursor.All(context.TODO(), &omLinks); err != nil {
		fmt.Println("[model.SearchOrder err] ", err)
		panic(err)
	}

	fmt.Println("[model.SearchOrder omLinks] ", omLinks)

	return omLinks
}

// 주문 리스트
func (p *Model) OrderStates(menu Menu) Menu {
	fmt.Println("Model param : ", menu)

	result, err := p.colPersons.InsertOne(context.TODO(), menu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted with ID: %s\n", result.InsertedID)

	return menu
}

// 메뉴 상세
func (p *Model) ViewMenu(menu Menu) Menu {
	fmt.Println("[model.ViewMenu Param] ", menu)

	var filter bson.D
	if menu.Name != "" {
		filter = append(filter, bson.E{"name", menu.Name})
		//filter = bson.D{{"name", menu.Name}}
	}
	if menu.Status != "" {
		filter = append(filter, bson.E{"status", menu.Status})
	}

	cursor, err := p.colPersons.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("[model.ViewMenu err] ", err)
		panic(err)
	}

	var menus []Menu
	if err = cursor.All(context.TODO(), &menus); err != nil {
		fmt.Println("[model.ViewMenu err] ", err)
		panic(err)
	}

	fmt.Println("[model.ViewMenu menus] ", menus)

	return menu
}

// 메뉴 리뷰
func (p *Model) CreateReview(omLink OrdererMenuLink) OrdererMenuLink {
	logger.Info("[model.CreateReview Param] ", omLink)
	fmt.Println("[model.CreateReview Param] ", omLink)

	omLink.CreateDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))
	omLink.ModifyDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))

	result, err := p.colPersons.InsertOne(context.TODO(), omLink)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted with ID: %s\n", result.InsertedID)

	return omLink
}

// 	filter := bson.D{{"pnum", pnum}}

// 	age, _ := strconv.Atoi(agestr)
// 	update := bson.M{
// 		"$set": bson.M{
// 			"age": age,
// 		},
// 	}
