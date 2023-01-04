package model

//model.go : db에 접속해 데이터를 핸들링, 결과 전달
import (
	"WBABEProject-20/go-ordering/conf"
	"WBABEProject-20/go-ordering/logger"
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client         *mongo.Client
	colMenu        *mongo.Collection
	colOrderLink   *mongo.Collection
	colUserAccount *mongo.Collection
}

// Model mongodb Connection
func NewModel(cfg *conf.Config) (*Model, error) {
	logger.Info("[model.NewModel] start...")

	cf := cfg.MongoDB
	r := &Model{}

	var err error
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cf.Host)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		logger.Info("[model.NewModel] mongodb connection... ")

		db := r.client.Database(cf.Database)
		r.colMenu = db.Collection(cf.MenuCollection)
		r.colOrderLink = db.Collection(cf.OrderCollection)
		r.colUserAccount = db.Collection(cf.UserCollection)

	}
	return r, nil
}

// 메뉴 등록 - 피주문자
func (p *Model) CreateMenu(menu Menu) Menu {
	fmt.Println("[model.CreateMenu Param] ", menu)

	menu.MenuID = CreateUUID() //메뉴 ID는 uuid로 설정

	result, err := p.colMenu.InsertOne(context.TODO(), menu)
	if err != nil {
		panic(err)
	}
	fmt.Println("inserted with ID: %s\n", result.InsertedID)

	return menu
}

// 메뉴 수정 - 피주문자
func (p *Model) UpdateMenu(menuID string, menu Menu, updateFilter bson.M) Menu {
	fmt.Println("[model.UpdateMenu Param] ", menu)

	/*
		업데이트 시에는 MenuId를 request body로 받는 것이 아니라 URI로서 값을 받아와 이용하는 것이 일반적입니다.
		e.g PATCH /menu/1(메뉴아이디 값)
	*/
	/*
		수정 내용
		controller.go에서 URI로서 값을 받아 처리
	*/
	//메뉴ID 기준으로 메뉴 업데이트
	filter := bson.D{{"menuID", menuID}}

	fmt.Println("[model.UpdateMenu filter] ", filter)

	result, err := p.colMenu.UpdateOne(context.Background(), filter, updateFilter)
	if err != nil {
		logger.Error(err)
	}
	fmt.Println("[model.UpdateMenu result] ", result)
	fmt.Println("[model.UpdateMenu result.ModifiedCount] ", result.ModifiedCount)

	return p.ViewMenu(menuID)
}

// 메뉴 삭제 - 피주문자 (삭제하지않고 상태변경으로 비표시)
func (p *Model) DeleteMenu(menuID string, isRecommededstr string) Menu {

	fmt.Println("[model.DeleteMenu Param menuID] ", menuID)
	fmt.Println("[model.DeleteMenu Param isRecommeded] ", isRecommededstr)

	isRecommeded, _ := strconv.ParseBool(isRecommededstr)

	filter := bson.D{{"menuID", menuID}}
	update := bson.M{
		"$set": bson.M{
			"isRecommeded": isRecommeded,
		},
	}

	//삭제하지않고 상태변경으로 비표시
	// result, err := p.colPersons.DeleteOne(context.TODO(), filter)
	result, err := p.colMenu.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
	}
	fmt.Println("UpdateOne() result:", result)
	fmt.Printf("Documents Updated: %v\n", result.ModifiedCount)

	//메뉴의 상세 내용을 리턴
	return p.ViewMenu(menuID)
}

// 메뉴 검색 - 주문자, 피주문자
func (p *Model) SearchMenu(filter bson.D) []Menu {
	fmt.Println("[model.SearchMenu Param] ", filter)

	cursor, err := p.colMenu.Find(context.TODO(), filter)
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	var menus []Menu
	if err = cursor.All(context.TODO(), &menus); err != nil {
		logger.Error(err)
		panic(err)
	}

	fmt.Println("[model.SearchMenu menus] ", menus)

	return menus
}

// 메뉴 상세 - 주문자, 피주문자
func (p *Model) ViewMenu(menuID string) Menu {
	fmt.Println("[model.ViewMenu Param] ", menuID)

	var menu Menu
	filter := bson.D{{"menuID", menuID}}

	err := p.colMenu.FindOne(context.TODO(), filter).Decode(&menu)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	fmt.Println("[model.ViewMenu menu] ", menu)

	return menu
}

// 메뉴 ID 검색
func (p *Model) GetMenuID(sellerID string, menuName string) string {
	sMenu := p.SearchMenuFindOne(sellerID, menuName)
	return sMenu.MenuID
}

// 메뉴 상세 - 주문자, 피주문자
func (p *Model) SearchMenuFindOne(sellerID string, menuName string) Menu {
	fmt.Println("[model.ViewMenu Param] ", sellerID)
	fmt.Println("[model.ViewMenu Param] ", menuName)

	var menu Menu
	filter := bson.D{{"sellerID", sellerID}, {"menuName", menuName}}

	err := p.colMenu.FindOne(context.TODO(), filter).Decode(&menu)
	if err != nil {
		logger.Error(err)
		//panic(err)
	}
	fmt.Println("[model.ViewMenu menu] ", menu)

	return menu
}

// 주문 등록 - 주문자
func (p *Model) NewOrder(omLink OrdererMenuLink) OrdererMenuLink {
	fmt.Println("[model.NewOrder Param] ", omLink)

	omLink.OrderNo = CreateUUID() //Order번호는 uuid로 설정
	omLink.OrderStatus = "주문확인중"  //주문시 상태 설정

	result, err := p.colOrderLink.InsertOne(context.TODO(), omLink)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	fmt.Printf("inserted with ID: %s\n", result.InsertedID)

	return omLink
}

// 주문 내역 조회 - 피주문자
func (p *Model) OrderStatus(orderStatus string, sellerID string) []OrdererMenuLink {
	fmt.Println("[model.OrderStates Param] ", orderStatus, sellerID)

	var filter bson.D
	if sellerID != "" { //판매자 ID 필수
		filter = append(filter, bson.E{"sellerID", sellerID})
	}
	if orderStatus != "" { //주문 들어온 리스트 (상태값에 따른 조회)
		filter = append(filter, bson.E{"orderStatus", orderStatus})
	}

	cursor, err := p.colOrderLink.Find(context.TODO(), filter)
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	var omLinks []OrdererMenuLink
	if err = cursor.All(context.TODO(), &omLinks); err != nil {
		logger.Error(err)
		panic(err)
	}

	fmt.Println("[model.SearchOrder omLinks] ", omLinks)

	return omLinks
}

// 주문 상세 - 주문자, 피주문자
func (p *Model) ViewOrder(orderNo string) OrdererMenuLink {
	fmt.Println("[model.ViewOrder Param] ", orderNo)

	var omLink OrdererMenuLink
	filter := bson.D{{"orderNo", orderNo}}

	err := p.colOrderLink.FindOne(context.TODO(), filter).Decode(&omLink)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	fmt.Println("[model.ViewOrder menu] ", omLink)

	return omLink
}

// 주문 변경 - 주문자 (수정/취소)
func (p *Model) ChangeOrder(omLink OrdererMenuLink) OrdererMenuLink {
	fmt.Println("[model.ChangeOrder Param] ", omLink)

	//omLink.OrderNO
	filter := bson.D{{"orderNo", omLink.OrderNo}}

	update := bson.M{
		"$set": bson.M{
			"orderStatus":    omLink.OrderStatus,
			"ordererAddress": omLink.OrdererAddress,
			"ordererPhone":   omLink.OrdererPhone,
			//"modifyDate": omLink.ModifyDate,
		},
	}

	fmt.Println("[model.ChangeOrder filter] ", filter)

	result, err := p.colOrderLink.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
	}
	fmt.Println("[model.ChangeOrder result] ", result)
	fmt.Println("[model.ChangeOrder result.ModifiedCount] ", result.ModifiedCount)

	return omLink
}

// 주문 내역 조회 기능 - 주문자
func (p *Model) SearchOrder(omLink OrdererMenuLink) []OrdererMenuLink {
	fmt.Println("[model.SearchOrder Param] ", omLink)

	filter := bson.D{{"ordererID", omLink.OrdererID}}
	if omLink.MenuName != "" {
		filter = append(filter, bson.E{"menuName", omLink.MenuName})
	}
	if omLink.OrderStatus != "" {
		filter = append(filter, bson.E{"orderStatus", omLink.OrderStatus})
	}

	cursor, err := p.colOrderLink.Find(context.TODO(), filter)
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	var omLinks []OrdererMenuLink
	if err = cursor.All(context.TODO(), &omLinks); err != nil {
		logger.Error(err)
		panic(err)
	}

	fmt.Println("[model.SearchOrder omLinks] ", omLinks)

	return omLinks
}

// 리뷰 등록 - 주문자
func (p *Model) CreateReview(omLink OrdererMenuLink) OrdererMenuLink {
	fmt.Println("[model.CreateReview Param] ", omLink)

	//omLink.OrderNO
	filter := bson.D{{"orderNo", omLink.OrderNo}}

	update := bson.M{
		"$set": bson.M{
			"orderComment":   omLink.OrderComment,
			"orderStarGrade": omLink.OrderStarGrade,
		},
	}

	fmt.Println("[model.CreateReview filter] ", filter)

	result, err := p.colOrderLink.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
	}
	fmt.Println("[model.CreateReview result] ", result)
	fmt.Println("[model.CreateReview result.ModifiedCount] ", result.ModifiedCount)

	return omLink
}

// 메뉴 상세 - 주문자, 피주문자
func (p *Model) GetUserAccount(userID string) UserAccount {
	fmt.Println("[model.GetUserAccount userID] ", userID)

	var user UserAccount
	filter := bson.D{{"userID", userID}}

	err := p.colUserAccount.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	fmt.Println("[model.GetUserAccount user] ", user)

	return user
}
