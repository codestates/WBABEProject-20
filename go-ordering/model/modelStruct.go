package model

// 지금은 화면 처리 안되기 때문에 ID 없음. 구현 안함.
type Orderer struct {
	OrdererID  string `bson:"ordererID"`  //주문자 ID
	OrderName  string `bson:"orderName"`  //주문자 이름
	Address    string `bson:"address"`    //주문자 주소
	Phone      int    `bson:"phone"`      //주문자 폰번호
	OrderCount int    `bson:"orderCount"` //주문 숫자
	// CreateDate primitive.DateTime `bson:"createDate"` //등록일
	// ModifyDate primitive.DateTime `bson:"modifyDate"` //수정일
}

// 지금은 화면 처리 안되기 때문에 ID 없음. 구현 안함.
type Seller struct {
	SellerID   string `bson:"sellerID"`   //판매자 ID
	SellerName string `bson:"sellerName"` //판매자이름
	Address    string `bson:"address"`    //판매자 주소
	Phone      int    `bson:"phone"`      //판매자 전화번호
	SellCount  int    `bson:"sellCount"`  //판매자 판매갯수
	// CreateDate primitive.DateTime `bson:"createDate"` //등록일
	// ModifyDate primitive.DateTime `bson:"modifyDate"` //수정일
}

type Menu struct {
	MenuID     string `bson:"menuID"`     //메뉴 ID
	SellerID   string `bson:"sellerID"`   //판매자 ID
	SellerName string `bson:"sellerName"` //판매자 이름
	MenuName   string `bson:"menuName"`   //메뉴 이름
	Status     string `bson:"status"`     //주문 가능 상태
	MaxCount   int    `bson:"maxCount"`   //판매 가능 갯수
	CountryOf  string `bson:"countryOf"`  //원산지
	Price      int    `bson:"price"`      //가격
	Spicy      string `bson:"spicy"`      //맵기
	Popularity int    `bson:"popularity"` //인기도
	IsDisabled bool   `bson:"isDisabled"` //판매여부
	TodayMenu  bool   `bson:"todayMenu"`  //오늘의 추천메뉴 여부
	Category   string `bson:"category"`   //메뉴 카테고리
	// CreateDate primitive.DateTime `bson:"createDate"` //등록일
	// ModifyDate primitive.DateTime `bson:"modifyDate"` //수정일
}

type OrdererMenuLink struct {
	OrderNo        string `bson:"orderNO"`        //주문번호
	MenuID         string `bson:"menuID"`         //메뉴 ID
	OrdererID      string `bson:"ordererID"`      //주문자ID
	MenuName       string `bson:"menuName"`       //메뉴이름
	OrderStarGrade int    `bson:"orderStarGrade"` //평점
	OrderComment   string `bson:"ordercomment"`   //후기
	OrderStatus    string `bson:"orderStatus"`    //주문상태
	OrdererAddress string `bson:"ordererAddress"` //주문자 주소
	OrdererPhone   int    `bson:"ordererPhone"`   //주문자 폰번호
	// CreateDate     primitive.DateTime `bson:"createDate"`     //등록일
	// ModifyDate     primitive.DateTime `bson:"modifyDate"`     //수정일
}