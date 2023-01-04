package router

//router.go : api 전체 인입에 대한 관리 및 구성을 담당하는 파일
import (
	ctl "WBABEProject-20/go-ordering/controller"
	logger "WBABEProject-20/go-ordering/logger"
	"fmt"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"

	"WBABEProject-20/go-ordering/docs" //swagger에 의해 자동 생성된 package
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {

	r := &Router{ct: ctl} //controller 포인터를 ct로 복사, 할당

	return r, nil
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("router.CORS c : ", c)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") //http://localhost:8080
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 임의 인증을 위한 함수
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("router.liteAuth c : ", c)
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		//http 헤더내 "Authorization" 폼의 데이터를 조회
		auth := c.GetHeader("Authorization")
		//실제 인증기능이 올수있다. 단순히 출력기능만 처리 현재는 출력예시
		fmt.Println("Authorization-word ", auth)

		c.Next()
	}
}

// 실제 라우팅
func (p *Router) Idx() *gin.Engine {

	// 컨피그나 상황에 맞게 gin 모드 설정
	//gin.SetMode(gin.ReleaseMode)

	//r := gin.Default() //gin 선언
	r := gin.Default()

	// 기존의 logger, recovery 대신 logger에서 선언한 미들웨어 사용
	//r.Use(gin.Logger())   //gin 내부 log, logger 미들웨어 사용 선언
	//r.Use(gin.Recovery()) //gin 내부 recover, recovery 미들웨어 사용 - 패닉복구
	r.Use(logger.GinLogger())
	r.Use(logger.GinRecovery(true))

	r.Use(CORS()) //crossdomain 미들웨어 사용 등록

	logger.Info("[router Idx] start server...")

	r.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	//r.POST("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	//r.PUT("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))

	docs.SwaggerInfo.Host = "localhost:8080" //swagger 정보 등록

	//피주문자 그룹
	seller := r.Group("oos/seller", liteAuth())
	{

		/*
			해당 부분에 대해서 전반적으로 코멘트를 남깁니다.

			1. REST에 대해서 조금 더 공부가 필요해 보입니다.
				자원을 가져오는 것 : GET
				자원을 생성하는 것 : POST
				자원을 업데이트 하는 것 : PUT, PATCH
				자원을 삭제 하는 것 : DELETE
				기본적으로 위의 규칙을 지켜야 REST API라고 할 수 있습니다. 자세한 내용은 'REST API 성숙도 모델' 의 키워드로 검색해보시면 좋을 것 같습니다.
			2. 엔드포인트 네이밍으로 create, update, delete, search와 같은 것이 들어갈 필요가 없습니다.
				HTTP URI의 이름에 위와 같은 단어가 들어가는 것이 아닌,
				/order/menu 라는 URI로 GET을 한다면 메뉴에 대한 정보를 가져오는 것, POST 라면 메뉴를 생성하는 것, PATCH 라면 메뉴 정보를 업데이트하는 것과 같은 패턴으로 구성을 하셔야 합니다.
		*/

		seller.POST("/menu", p.ct.CreateMenu)
		/*
			메뉴에 대한 정보를 업데이트하는 것은 POST가 아니라 PUT, 혹은 PATCH가 되어야 합니다.
		*/
		/*
			수정내용
			PUT 변경
		*/
		seller.PUT("/menu/:menuID", p.ct.UpdateMenu)
		/*
			메뉴를 삭제하는 API인데 PUT이 아닌 DELETE가 되어야 할 것 같습니다.
		*/
		/*
			수정내용
			DELETE로 변경
		*/
		seller.DELETE("/menu", p.ct.DeleteMenu)
		/*
			메뉴에 대한 정보를 가져오는 것이므로, GET이 되어야 합니다.
		*/
		/*
			수정내용
			GET으로 변경
		*/
		seller.GET("/menu", p.ct.SearchMenu)

		seller.GET("/orderStatus", p.ct.OrderStatus)
		seller.PUT("/setTodayMenu", p.ct.SetTodayMenu)

	}

	//주문자 그룹
	order := r.Group("oos/order", liteAuth())
	{

		order.GET("/viewMenu", p.ct.ViewMenu)

		order.POST("/newOrder", p.ct.NewOrder)
		order.PUT("/changeOrder", p.ct.ChangeOrder)
		order.GET("/searchOrder", p.ct.SearchOrder)
		order.GET("/viewOrder", p.ct.ViewOrder)

		order.POST("/createReview", p.ct.CreateReview)
		order.GET("/searchTodayMenu", p.ct.SearchTodayMenu)

	}

	return r
}
