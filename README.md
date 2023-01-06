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
 
 #validator
 $ get github.com/go-playground/validator
 </code></pre>
 
 ## 4. DataBase
 ### Database : go-ready
 ### 유저 Collection : tUserAccount
 ### 메뉴 Collection : tMenu
 ### 주문리스트 Collection : tOrdererMenuLink
 >>tOrdererMenuLink에 MenuID 속성을 추가하여 tMenu와 링크 관리
 
 ### 초기값 설정 : 유저 설정을 위해 DB에 유저값을 INSERT한다. (유저등록은 구현안함)
 <pre><code>
 db.tUserAccount.insertMany([{userID:"LEE",userName:"이철수",userType:"판매자"}
,{userID:"KIM",userName:"김영희",userType:"주문자"}])
 </code></pre>
 
 ### 속성
 <pre><code>
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
	IsRecommeded bool `bson:"isRecommeded"` 		     //판매여부 default(true)
	IsTdoayMenu bool `bson:"isTdoayMenu"` 			     //오늘의 추천메뉴 여부 default(false)
	Category []string `bson:"category"`			     //메뉴 카테고리 Enums(한식, 일식, 중식)
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

type UserAccount struct {
	UserID     string `bson:"userID"`     //주문자 ID
	UserName   string `bson:"userName"`   //주문자 이름
	UserType   string `bson:"userType"`   //판매자, 주문자 nums(판매자, 주문자)
	Address    string `bson:"address"`    //주문자 주소
	Phone      int    `bson:"phone"`      //주문자 폰번호
	OrderCount int    `bson:"orderCount"` //주문 숫자
	SellCount  int    `bson:"sellCount"`  //주문 숫자
}
 </code></pre>

 ## 5. API 구현 기능
 ### 피주문자 
 <pre><code>
 /oos/seller/menu [POST]         // @Description  메뉴 등록 - 피주문자
 /oos/seller/menu{menuID} [PUT]  // @Description  메뉴 수정 - 피주문자 (메뉴ID를 기준으로 메뉴 업데이트)
 /oos/seller/menu [DELETE]       // @Description  메뉴 삭제 - 피주문자 (판매여부 bool 설정변경)
 /oos/seller/menu [GET]          // @Description  메뉴 검색 - 주문자, 피주문자
 /oos/seller/orderStatus         // @Description  주문 내역 조회 - 피주문자 (판매자ID 기준으로 검색)
 /oos/seller/setTodayMenu        // @Description  오늘의 추천메뉴 여부 - 설정 변경 (메뉴ID를 기준으로 메뉴 업데이트)
 </code></pre>
 
 ### 주문자 
 <pre><code>
 /oos/order/viewMenu        // @Description  주문 상세 - 주문자, 피주문자
 /oos/order/newOrder        // @Description  주문 등록 - 주문자
 /oos/order/changeOrder     // @Description  주문 변경 - 주문자 (주문변경 커멘드 Enums(주문추가, 주문취소, 정보변경))
 /oos/order/searchOrder     // @Description  주문 내역 조회 기능 - 주문자 (주문자ID, 주문상태로 조회)
 /oos/order/viewOrder       // @Description  주문 상세 - 주문자, 피주문자
 /oos/order/createReview    // @Description  리뷰 등록 - 주문자 (주문번호 기준으로 등록
 /oos/order/searchTodayMenu // @Description  오늘의 추천메뉴 리스트
 </code></pre>

 ### Swagger 참고
 ![image](https://user-images.githubusercontent.com/119834304/210493375-f3c12e9a-b0f3-4355-ae01-7d13ddfc9507.png)

 #### /oos/seller/menu [POST]     // @Description  메뉴 등록 - 피주문자
 <pre><code>
  {
  "category":  ["한식","중식"],
  "countryOf": "대한민국",
  "isRecommeded": true,
  "maxCount": 20,
  "menuName": "마라탕면",
  "price": 9000,
  "sellerID": "LEE",
  "sellerName": "리반점",
  "spicy": "보통",
  "status": "판매중",
  "isTdoayMenu": true
 }
 </code></pre>
 ![image](https://user-images.githubusercontent.com/119834304/210969409-6cacebe4-ceee-42bd-9487-40289c2926e3.png)
 ![image](https://user-images.githubusercontent.com/119834304/210969477-7f6bd7e9-21e4-498e-a2a8-0a0d5376e93d.png)

 #### /oos/seller/menu/{menuID} [PUT]     // @Description  메뉴 수정 - 피주문자 (메뉴ID를 기준으로 메뉴 업데이트)
 <pre><code>
 "menuID": "ee097877-2878-43da-b312-f72f4a233089"
  {
  "category":  ["일식","중식"],
  "countryOf": "중국",
  "popularity" : 4,
  "maxCount" : 20,
  "isRecommeded": true
 }
 </code></pre>
 ![image](https://user-images.githubusercontent.com/119834304/210969628-550c8431-a9ff-4872-9315-a253ff24b3b7.png)
 ![image](https://user-images.githubusercontent.com/119834304/210969747-11cd23fd-6446-498d-8fed-75254faf160e.png)

 #### /oos/seller/menu [DELETE]     // @Description  메뉴 삭제 - 피주문자 (판매여부 bool 설정변경)
 <pre><code>
   "menuID": "ee097877-2878-43da-b312-f72f4a233089"
   "isRecommeded": true
 </code></pre>
 ![image](https://user-images.githubusercontent.com/119834304/210489805-2e61bf73-aa0a-43ea-85e3-1c108ca19082.png)
 ![image](https://user-images.githubusercontent.com/119834304/210489832-042bdb5d-8c7d-45f4-8fb0-c33f0ca827ec.png)

 #### /oos/seller/menu [GET]     // @Description  메뉴 검색 - 주문자, 피주문자
 ![image](https://user-images.githubusercontent.com/119834304/210970535-2506291d-4679-448b-a8fc-01999e06463d.png)
 ![image](https://user-images.githubusercontent.com/119834304/210970597-31b55498-9ecc-4966-a7a3-a3cd975a49c7.png)
 
 #### /oos/seller/orderStatus    // @Description  주문 내역 조회 - 피주문자 (판매자ID 기준으로 검색)
 ![image](https://user-images.githubusercontent.com/119834304/209762474-060f077c-846f-49b1-80c9-116b49031f48.png)
 ![image](https://user-images.githubusercontent.com/119834304/209762492-c24d3be0-ad40-4ccd-8e85-0d4d3da995be.png)

 #### /oos/seller/setTodayMenu   // @Description  오늘의 추천메뉴 여부 - 설정 변경 (메뉴ID를 기준으로 메뉴 업데이트)
 <pre><code>
 {
  "todayMenu": true,
  "menuID": "ee097877-2878-43da-b312-f72f4a233089"
 }
 </code></pre>
 ![image](https://user-images.githubusercontent.com/119834304/209763087-ffb6dba2-05e5-4f04-938f-728d2bb88fe8.png)
 ![image](https://user-images.githubusercontent.com/119834304/209763124-135a0ff8-6633-431b-9378-59740b6703cf.png)

 #### /oos/order/newOrder        // @Description  주문 등록 - 주문자
 <pre><code>
 {
  "menuID": "ee097877-2878-43da-b312-f72f4a233089",
   "ordererID": "KIM",
   "ordererAddress": "서울시 광진구",
   "ordererPhone": 1012345678
 }
 </code></pre>
 ![image](https://user-images.githubusercontent.com/119834304/209763399-747450cb-e385-438a-8d3a-6dbd35c97124.png)
 ![image](https://user-images.githubusercontent.com/119834304/209763425-90992dc2-ed49-442e-ba2a-3d76e6390122.png)
 
 #### /oos/order/viewMenu        // @Description  주문 상세 - 주문자, 피주문자
 ![image](https://user-images.githubusercontent.com/119834304/209763487-8cc40104-bb36-483a-81c1-6e5e3cb07eea.png)
 ![image](https://user-images.githubusercontent.com/119834304/209763512-4bcc932c-f456-47d2-95e1-23534fd26403.png)
 
 #### /oos/order/changeOrder     // @Description  주문 변경 - 주문자 (주문변경 커멘드 Enums(주문추가, 주문취소, 정보변경))
 <pre><code>
 {
  "changeOrderCmd": "정보변경",
  "orderNo": "5ca65c501bc8409ca92a2a9496170943",
  "ordererAddress": "서울시 서초구",
  "ordererPhone": 1043719999
 }
 </code></pre>
 ![image](https://user-images.githubusercontent.com/119834304/209763671-a646dcd3-5ad8-4440-914c-11c7573ffa07.png)
 ![image](https://user-images.githubusercontent.com/119834304/209763693-51cf1c96-347c-4e45-82b9-a381675765a6.png)

 #### /oos/order/searchOrder     // @Description  주문 내역 조회 기능 - 주문자 (주문자ID, 주문상태로 조회)
 ![image](https://user-images.githubusercontent.com/119834304/209763733-7166a334-27e5-4772-8cf3-d3dec9d4a099.png)
 ![image](https://user-images.githubusercontent.com/119834304/209763762-f12e907c-1a01-46e1-8b40-8e63cf350d62.png)

 #### /oos/order/viewOrder       // @Description  주문 상세 - 주문자, 피주문자
 ![image](https://user-images.githubusercontent.com/119834304/209765846-880edb3b-9dde-4c92-a080-b3e229ae5dbd.png)
 ![image](https://user-images.githubusercontent.com/119834304/209765866-23720efd-0fc5-4938-a82d-179eddc2cd7f.png)
 
 #### /oos/order/createReview    // @Description  리뷰 등록 - 주문자 (주문번호 기준으로 등록
 <pre><code>
 {
  "orderComment": "맛있습니다.",
  "orderNo": "5ca65c501bc8409ca92a2a9496170943",
  "orderStarGrade": 5,
  "ordererID" : "KIM"
 }
 </code></pre>
 ![image](https://user-images.githubusercontent.com/119834304/211040540-7496e871-7923-4dfc-a913-7960417dae4d.png)
 ![image](https://user-images.githubusercontent.com/119834304/211040606-0160c73f-91a2-43f6-a34f-ed2143180dfe.png)
 
 #### /oos/order/searchTodayMenu // @Description  오늘의 추천메뉴 리스트
 ![image](https://user-images.githubusercontent.com/119834304/209766156-e8cfd9ea-c3e6-4cc4-94c4-cfa108d107a4.png)
 ![image](https://user-images.githubusercontent.com/119834304/209766185-d9807d38-e6ab-4f16-9446-084343209083.png)
