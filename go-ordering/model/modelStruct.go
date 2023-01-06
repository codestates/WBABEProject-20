package model

/*
	구현이 되어 있지는 않은 부분이지만 코멘트 드립니다.
	보통은 주문자와 피 주문자를 다른 것으로 나누는 것 보다는
	유저라는 구조에서 판매자인지, 구매자인지를 필드를 통해 구분합니다. 현재의 구조에서는 중복된 필드들이 많이 존재하니까요.
*/
/*
	수정내용
	Order 와 Seller를 삭제하고,
	UserAccount 하나로 수정, UserType으로 판매자/주문자를 구분
*/

// DB는 INSERT하고, 값을 가져오기 위해 사용
type UserAccount struct {
	UserID     string `bson:"userID"`     //주문자 ID
	UserName   string `bson:"userName"`   //주문자 이름
	UserType   string `bson:"userType"`   //판매자, 주문자 nums(판매자, 주문자)
	Address    string `bson:"address"`    //주문자 주소
	Phone      int    `bson:"phone"`      //주문자 폰번호
	OrderCount int    `bson:"orderCount"` //주문 숫자
	SellCount  int    `bson:"sellCount"`  //주문 숫자
}

/*
수정내용
validate 값 min, max를 정의
*/
type Menu struct {
	MenuID     string `bson:"menuID"`                            //메뉴 ID
	SellerID   string `bson:"sellerID"`                          //판매자 ID
	SellerName string `bson:"sellerName"`                        //판매자 이름
	MenuName   string `bson:"menuName"`                          //메뉴 이름
	Status     string `bson:"status"`                            //주문 가능 상태 nums(준비중, 판매중)
	MaxCount   int    `bson:"maxCount" validate:"min=1,max=50"`  //판매 가능 갯수 mininum(1) maxinum(50)
	CountryOf  string `bson:"countryOf"`                         //원산지 Enums(한국, 일본, 중국)
	Price      int    `bson:"price"`                             //가격
	Spicy      string `bson:"spicy"`                             //맵기 Enums(아주매움, 매움, 보통, 순한맛)
	Popularity int    `bson:"popularity" validate:"min=1,max=5"` //인기도 mininum(1) maxinum(5)
	/*
		bool 값들의 경우는 네이밍시에 일반적으로 긍정의 단어를 사용하고, 그 여부는 true, false로 제어합니다.
		즉, 추천드리는 네이밍은 IsAvailable, IsPublic, ForSale 이 되겠습니다.
		disable이라는 부정의 의미보다는 긍정의 의미로 네이밍을 짓고 변수 값으로 여부를 판단하는 것이 더 읽기에 자연스럽습니다.
	*/
	/*
		수정내용
		명칭 변경
		IsDisabled > IsRecommeded
		TodayMenu > IsTdoayMenu
	*/
	IsRecommeded bool `bson:"isRecommeded"` //판매여부 default(true)
	/*
		TodayMenu만 보고서는 bool 값인지 유추하기가 힘들어 보입니다. 현재는 오늘의 메뉴가 무엇인지 하는 String 데이터가 예상이 됩니다.
		따라서 IsRecommeded, IsTdoayMenu 과 같은 네이밍이 적절해 보입니다.
	*/
	IsTdoayMenu bool `bson:"isTdoayMenu"` //오늘의 추천메뉴 여부 default(false)

	/*
		수정내용
		메뉴 카테고리를 별도의 구조체로 생성해서 메뉴에 카테고리를 0~2개 이상의 값을 지니도록 수정
	*/
	Category []string `bson:"category"` //메뉴 카테고리 Enums(한식, 일식, 중식)
}

type OrdererMenuLink struct {
	OrderNo        string `bson:"orderNo"`                               //주문번호
	SellerID       string `bson:"sellerID"`                              //판매자 ID
	MenuID         string `bson:"menuID"`                                //메뉴 ID
	OrdererID      string `bson:"ordererID"`                             //주문자ID
	MenuName       string `bson:"menuName"`                              //메뉴이름
	OrderStarGrade int    `bson:"orderStarGrade" validate:"min=1,max=5"` //평점 mininum(1) maxinum(5)
	OrderComment   string `bson:"ordercomment"`                          //후기
	OrderStatus    string `bson:"orderStatus"`                           //주문상태 Enums(주문확인중 - 조리중 - 배달중 - 배달완료 - 주문취소)
	OrdererAddress string `bson:"ordererAddress"`                        //주문자 주소
	OrdererPhone   int    `bson:"ordererPhone"`                          //주문자 폰번호
}

/*
수정내용
필수에는 binding:"required"를 추가
*/
type BindChangeOrderState struct {
	OrderNo        string `bson:"orderNo" binding:"required"` //주문번호
	OrderStatus    string `bson:"orderStatus"`                //주문상태 Enums(주문확인중 - 조리중 - 배달중 - 배달완료 - 주문취소)
	OrdererAddress string `bson:"ordererAddress"`             //주문자 주소
	OrdererPhone   int    `bson:"ordererPhone"`               //주문자 폰번호
	ChangeOrderCmd string `bson:"changeOrderCmd"`             //주문 변경 Enums(주문추가, 주문취소, 정보변경)
}
