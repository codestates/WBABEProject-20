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
	client       *mongo.Client
	colMenu      *mongo.Collection
	colOrderLink *mongo.Collection
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
	}
	return r, nil
}

// 메뉴 등록 - 피주문자
func (p *Model) CreateMenu(menu Menu) Menu {
	logger.Info("[model.CreateMenu Param] ", menu)

	//menu.CreateDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))
	//menu.ModifyDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))

	result, err := p.colMenu.InsertOne(context.TODO(), menu)
	if err != nil {
		panic(err)
	}
	logger.Info("inserted with ID: %s\n", result.InsertedID)

	return menu
}

// 메뉴 수정 - 피주문자
func (p *Model) UpdateMenu(menu Menu) Menu {
	logger.Info("[model.UpdateMenu Param] ", menu)

	filter := bson.D{{"sellerID", menu.SellerID}, {"menuID", menu.MenuID}}

	update := bson.M{
		"$set": bson.M{
			// "status":     menu.Status,
			// "maxCount":   menu.MaxCount,
			// "countryOf":  menu.CountryOf,
			// "price":      menu.Price,
			// "spicy":      menu.Spicy,
			// "popularity": menu.Popularity,
			// "isDisabled": menu.IsDisabled,
			// "todayMenu":  menu.TodayMenu,
			// "category":   menu.Category,
			//"createDate": menu.CreateDate,
			//"modifyDate": menu.ModifyDate,
		},
	}

	result, err := p.colMenu.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("[model.UpdateMenu result] ", result)
	logger.Info("[model.UpdateMenu result.ModifiedCount] ", result.ModifiedCount)

	return p.ViewMenu(menu.MenuName)
}

// 메뉴 삭제 - 피주문자 (삭제하지않고 상태변경으로 비표시)
func (p *Model) DeleteMenu(menuName string, isDisabled string) Menu {

	logger.Info("[model.DeleteMenu Param menuName] ", menuName)
	logger.Info("[model.DeleteMenu Param isDisabled] ", isDisabled)

	filter := bson.D{{"menuName", menuName}}

	boolIsDisabled, err := strconv.ParseBool(isDisabled)
	if err != nil {
		logger.Error(err)
	}

	update := bson.M{
		"$set": bson.M{
			"isDisabled": boolIsDisabled,
		},
	}

	//삭제하지않고 상태변경으로 비표시
	// result, err := p.colPersons.DeleteOne(context.TODO(), filter)
	result, err := p.colMenu.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("UpdateOne() result:", result)
	fmt.Printf("Documents Updated: %v\n", result.ModifiedCount)

	//메뉴의 상세 내용을 리턴
	return p.ViewMenu(menuName)
}

// 메뉴 검색 - 주문자, 피주문자
func (p *Model) SearchMenu(menu Menu) []Menu {
	logger.Info("[model.SearchMenu Param] ", menu)

	var filter bson.D
	if menu.MenuName != "" {
		filter = append(filter, bson.E{"menuName", menu.MenuName})
	}
	if menu.Status != "" {
		filter = append(filter, bson.E{"status", menu.Status})
	}

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

	logger.Info("[model.SearchMenu menus] ", menus)

	return menus
}

// 메뉴 상세 - 주문자, 피주문자
func (p *Model) ViewMenu(menuName string) Menu {
	logger.Info("[model.ViewMenu Param] ", menuName)

	var menu Menu
	filter := bson.D{{"menuName", menuName}}

	err := p.colMenu.FindOne(context.TODO(), filter).Decode(&menu)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	logger.Info("[model.ViewMenu menu] ", menu)

	return menu
}

// 주문 등록 - 주문자
func (p *Model) NewOrder(omLink OrdererMenuLink) OrdererMenuLink {
	logger.Info("[model.NewOrder Param] ", omLink)

	//omLink.CreateDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))
	//omLink.ModifyDate = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))

	result, err := p.colOrderLink.InsertOne(context.TODO(), omLink)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	fmt.Printf("inserted with ID: %s\n", result.InsertedID)

	return omLink
}

// 주문 내역 조회 - 피주문자
func (p *Model) OrderStatus(menuName string, orderStatus string) []OrdererMenuLink {
	logger.Info("[model.OrderStates Param] ", menuName, ", ", orderStatus)

	var filter bson.D
	if menuName != "" { //메뉴에 따른 주문 들어온 리스트
		filter = append(filter, bson.E{"menuName", menuName})
	}
	if orderStatus != "" { //주문 들어온 리스트
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

	logger.Info("[model.SearchOrder omLinks] ", omLinks)

	return omLinks
}

// 주문 변경 - 주문자 (수정/취소)
func (p *Model) ChangeOrder(omLink OrdererMenuLink) OrdererMenuLink {
	logger.Info("[model.ChangeOrder Param] ", omLink)

	//omLink.OrderNO
	filter := bson.D{{"menuName", omLink.MenuName}, {"ordererID", omLink.OrdererID}}

	update := bson.M{
		"$set": bson.M{
			"orderStatus":    omLink.OrderStatus,
			"ordererAddress": omLink.OrdererAddress,
			"ordererPhone":   omLink.OrdererPhone,
			//"modifyDate": omLink.ModifyDate,
		},
	}

	logger.Info("[model.ChangeOrder filter] ", filter)

	result, err := p.colOrderLink.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("[model.ChangeOrder result] ", result)
	logger.Info("[model.ChangeOrder result.ModifiedCount] ", result.ModifiedCount)

	return omLink
}

// 주문 내역 조회 기능 - 주문자
func (p *Model) SearchOrder(omLink OrdererMenuLink) []OrdererMenuLink {
	logger.Info("[model.SearchOrder Param] ", omLink)

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

	logger.Info("[model.SearchOrder omLinks] ", omLinks)

	return omLinks
}

// 리뷰 등록 - 주문자
func (p *Model) CreateReview(omLink OrdererMenuLink) OrdererMenuLink {
	logger.Info("[model.CreateReview Param] ", omLink)

	//omLink.OrderNO
	filter := bson.D{{"menuName", omLink.MenuName}, {"ordererID", omLink.OrdererID}}

	update := bson.M{
		"$set": bson.M{
			"orderComment":   omLink.OrderComment,
			"orderStarGrade": omLink.OrderStarGrade,
			//"modifyDate": omLink.ModifyDate,
		},
	}

	logger.Info("[model.CreateReview filter] ", filter)

	result, err := p.colOrderLink.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("[model.CreateReview result] ", result)
	logger.Info("[model.CreateReview result.ModifiedCount] ", result.ModifiedCount)

	return omLink
}
