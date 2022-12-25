# WBABEProject-20

## 1.프로젝트 정보
프로젝트 제목 : 온라인 주문 시스템
프로젝트 내용 : Golang과 mongodb를 이용하여 온라인 주문 시스템의 Backend API 개발</br>
 ①주문자/피주문자의 역할에 필수적인 기능을 구현</br>
 ②주문자와 피주문자의 입장에서 필요한 기능을 도출하여 주문 서비스 기능을 구현</br>
 ③주문부터 배달까지 주문내역 관리 서비스 기능을 구현</br>
 
## 2.구조
📦go-ordering </br>
 ┣ 📂conf </br>
 ┃ ┣ 📜config.go </br>
 ┃ ┗ 📜config.toml </br>
 ┣ 📂controller </br>
 ┃ ┣ 📜controller.go </br>
 ┃ ┗ 📜controllerConvParams.go </br>
 ┣ 📂docs </br>
 ┃ ┣ 📜docs.go </br>
 ┃ ┣ 📜swagger.json </br>
 ┃ ┗ 📜swagger.yaml </br>
 ┣ 📂logger </br>
 ┃ ┗ 📜logger.go </br>
 ┣ 📂logs </br>
 ┃ ┗ 📜go-loger_2022-12-25.log </br>
 ┣ 📂model </br>
 ┃ ┣ 📜model.go </br>
 ┃ ┣ 📜modelBody.go </br>
 ┃ ┣ 📜modelDataCheck.go </br>
 ┃ ┗ 📜modelStruct.go </br>
 ┣ 📂router </br>
 ┃ ┗ 📜router.go </br>
 ┣ 📜go.mod </br>
 ┣ 📜go.sum </br>
 ┗ 📜main.go </br>
 
 ## 3.프로젝트에 필요한 Go 패키지
 <pre><code>
 #gin
 $ go get "github.com/gin-gonic/gin"
 
 #errorgroup
 $ go mod download golang.org/x/sync
 
 #mongodb 
 $ go get go.mongodb.org/mongo-driver/mongo
 $ go get go.mongodb.org/mongo-driver/mongo/options
 $ go get go.mongodb.org/mongo-driver/bson
 
 #swagger
 $ go get -u github.com/swaggo/swag/cmd/swag
 $ go install github.com/swaggo/swag/cmd/swag@latest
 $ go get -u github.com/swaggo/gin-swagger
 $ go get -u github.com/swaggo/files
 
 #toml
 $ go get "github.com/naoina/toml" 
 
 #log
 $ go get "github.com/natefinch/lumberjack"
 $ go get "go.uber.org/zap"
 $ go get "go.uber.org/zap/zapcore"
 
 #uuid
 $ go get github.com/google/uuid
 </code></pre>
 
 ## 4. API 구현 기능
 ### 피주문자 
 <pre><code>
 /oos/seller/createMenu     // @Description  메뉴 등록 - 피주문자
 /oos/seller/updateMenu     // @Description  메뉴 수정 - 피주문자 (메뉴ID를 기준으로 메뉴 업데이트)
 /oos/seller/deleteMenu     // @Description  메뉴 삭제 - 피주문자 (판매여부 bool 설정변경)
 /oos/order/searchMenu      // @Description  메뉴 검색 - 주문자, 피주문자
 /oos/order/viewMenu        // @Description  메뉴 상세 - 주문자, 피주문자
 /oos/seller/setTodayMenu   // @Description  오늘의 추천메뉴 여부 - 설정 변경 (메뉴ID를 기준으로 메뉴 업데이트)
 /oos/seller/OrderStatus    // @Description  주문 내역 조회 - 피주문자 (판매자ID 기준으로 검색)
 </code></pre>
 
 ### 주문자 
 <pre><code>
 /oos/order/newOrder        // @Description  주문 등록 - 주문자
 /oos/order/viewOrder       // @Description  주문 상세 - 주문자, 피주문자
 /oos/order/searchOrder     // @Description  주문 내역 조회 기능 - 주문자 (주문자ID, 주문상태로 조회)
 /oos/order/searchTodayMenu // @Description  오늘의 추천메뉴 리스트
 /oos/order/changeOrder     // @Description  주문 변경 - 주문자 (주문변경 커멘드 Enums(주문추가, 주문취소, 정보변경))
 /oos/order/createReview    // @Description  리뷰 등록 - 주문자 (주문번호 기준으로 등록
 </code></pre>

 ### Swagger 참고
 ![image](https://user-images.githubusercontent.com/119834304/209469839-0d5d8805-ef48-48ec-b593-53c937deb123.png)

 ## 5. DataBase
 ### Database : go-ready
 ### 메뉴 Collection : tMenu
 ### 주문리스트 Collection : tOrdererMenuLink
 >>tOrdererMenuLink에 MenuID 속성을 추가하여 tMenu와 링크 관리
 ### 속성
 <pre><code>
 type Menu struct {
	MenuID     string `bson:"menuID"`     //메뉴 ID
	SellerID   string `bson:"sellerID"`   //판매자 ID
	SellerName string `bson:"sellerName"` //판매자 이름
	MenuName   string `bson:"menuName"`   //메뉴 이름
	Status     string `bson:"status"`     //주문 가능 상태 nums(준비중, 판매중)
	MaxCount   int    `bson:"maxCount"`   //판매 가능 갯수 mininum(1) maxinum(50)
	CountryOf  string `bson:"countryOf"`  //원산지 Enums(한국, 일본, 중국)
	Price      int    `bson:"price"`      //가격
	Spicy      string `bson:"spicy"`      //맵기 Enums(아주매움, 매움, 보통, 순한맛)
	Popularity int    `bson:"popularity"` //인기도 mininum(1) maxinum(5)
	IsDisabled bool   `bson:"isDisabled"` //판매여부 default(true)
	TodayMenu  bool   `bson:"todayMenu"`  //오늘의 추천메뉴 여부 default(false)
	Category   string `bson:"category"`   //메뉴 카테고리 Enums(한식, 일식, 중식)
}

type OrdererMenuLink struct {
	OrderNo        string `bson:"orderNo"`        //주문번호
	SellerID       string `bson:"sellerID"`       //판매자 ID
	MenuID         string `bson:"menuID"`         //메뉴 ID
	OrdererID      string `bson:"ordererID"`      //주문자ID
	MenuName       string `bson:"menuName"`       //메뉴이름
	OrderStarGrade int    `bson:"orderStarGrade"` //평점 mininum(1) maxinum(5)
	OrderComment   string `bson:"ordercomment"`   //후기
	OrderStatus    string `bson:"orderStatus"`    //주문상태 Enums(주문확인중 - 조리중 - 배달중 - 배달완료 - 주문취소)
	OrdererAddress string `bson:"ordererAddress"` //주문자 주소
	OrdererPhone   int    `bson:"ordererPhone"`   //주문자 폰번호
}

 </code></pre>
