package controller

//controller.go : 실제 비지니스 로직 및 프로세스가 처리후 결과 전송
import (
	"WBABEProject-20/go-ordering/logger"
	"WBABEProject-20/go-ordering/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

// CreateMenu godoc
// @Summary call CreateMenu, return ok by json.
// @Description CreateMenu 메뉴 등록 - 피주문자
// @name CreateMenu
// @Accept  json
// @Produce  json
// @Param	Menu	body	model.Menu	true	"메뉴"
// @Router /oos/seller/menu [post]
// @Success 200 {object} model.Menu
func (p *Controller) CreateMenu(c *gin.Context) {
	logger.Info("[controller.CreateMenu] start...")

	req := c.Request
	req.ParseForm()
	r := req.Form
	for k, v := range r {
		fmt.Println(fmt.Sprintf("%s=%v", k, v))
	}

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
		fmt.Println("[controller.CreateMenu] params :", params.Category)
		sellerID := params.SellerID
		//메뉴 등록시 판매자 로그인 필수.
		/*
			수정내용
			로그인 유저를 확인하기 위해, UserAccount를 가져온다.
		*/
		user := p.md.GetUserAccount(sellerID)

		errChk, errMsg := checkCreateMenu(params, user)
		if errChk {
			logger.Error(errMsg)
			c.JSON(http.StatusBadRequest, errMsg)
			return
		}

		filter := bson.M{"sellerID": params.SellerID}
		filter["menuName"] = params.MenuName

		menus := p.md.SearchMenu(filter)
		for _, menu := range menus {
			logger.Info("[controller.CreateMenu] MenuName...", menu.MenuName)
			errMsg = "같은 이름의 메뉴가 존재합니다. 수정해주세요."
			logger.Error(errMsg)
			c.JSON(http.StatusBadRequest, errMsg)
			return
		}

		c.JSON(http.StatusOK, gin.H{"등록되었습니다.": p.md.CreateMenu(params)})

	} else {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, "등록에 실패했습니다.")
	}
}

// UpdateMenu godoc
// @Summary call UpdateMenu, return ok by json.
// @Description UpdateMenu 메뉴 수정 - 피주문자 (메뉴ID를 기준으로 메뉴 업데이트)
// @name UpdateMenu
// @Accept  json
// @Produce  json
// @Param  menuID path string true "menuID"
// @Param request	body	model.Menu	true	"변경할 메뉴"
// @Router /oos/seller/menu/{menuID} [PUT]
// @Success 200 {object} model.Menu
func (p *Controller) UpdateMenu(c *gin.Context) {
	logger.Info("[controller.UpdateMenu] start...")

	/*
		수정내용
		메뉴 삭제를 post에서 patch로 변경하면서 Param을 수정
	*/
	menuID := c.Param("menuID")
	if menuID == "" {
		errMsg := "메뉴ID가 입력되지 않았습니다."
		logger.Error(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	menu, updateFilter := UpdateMenuAppendQuery(c)

	fmt.Println("[controller.UpdateMenu] menu...", menu)
	fmt.Println("[controller.UpdateMenu] updateFilter...", updateFilter)

	resp := validationCheck(menu)
	fmt.Println("[controller.UpdateMenu] resp...", resp)

	if resp != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
	} else {
		c.JSON(http.StatusOK, p.md.UpdateMenu(menuID, menu, updateFilter))
	}
}

// DeleteMenu godoc
// @Summary call DeleteMenu, return ok by json.
// @Description DeleteMenu 메뉴 삭제 - 피주문자 (판매여부 bool 설정변경)
// @name DeleteMenu
// @Accept  json
// @Produce  json
// @Router /oos/seller/menu [delete]
// @Param menuID query string true "menuID"
// @Param isRecommeded query string true "isRecommeded"
// @Success 200 {object} model.Menu
func (p *Controller) DeleteMenu(c *gin.Context) {
	logger.Info("[controller.DeleteMenu] start...")

	/*
		수정내용
		메뉴 삭제를 put에서 delete로 변경하면서 Param을 수정
	*/
	menuId := c.Query("menuID")
	isRecommededstr := c.Query("isRecommeded")

	//menu, _ := UpdateMenuAppendQuery(c)

	fmt.Println("[controller.DeleteMenu menuId]", menuId)
	fmt.Println("[controller.DeleteMenu isRecommeded]", isRecommededstr)

	c.JSON(http.StatusOK, gin.H{"판매불가 설정되었습니다.": p.md.DeleteMenu(menuId, isRecommededstr)})
	// 	c.JSON(http.StatusOK, p.md.DeleteMenu(params))
}

// SearchMenu godoc
// @Summary call SearchMenu, return ok by json.
// @Description SearchMenu 메뉴 검색 - 주문자, 피주문자
// @name SearchMenu
// @Accept  json
// @Produce  json
// @Param SellerID query  string false "판매자 ID"
// @Param MenuName query  string false "메뉴이름"
// @Param Price query  int false "가격"
// @Param CountryOf query  string false "원산지" Enums(한국, 일본, 중국)
// @Param Category query  string false "메뉴 카테고리" Enums(한식, 일식, 중식)
// @Param Status query  string false "주문 가능 상태" Enums(준비중, 판매중, 판매완료)
// @Param Spicy query  string false "맵기" Enums(아주매움, 매움, 보통, 순한맛)
// @Param TodayMenu query  bool false "오늘의 추천메뉴 여부"
// @Router /oos/seller/menu [get]
// @Success 200 {object} Controller
func (p *Controller) SearchMenu(c *gin.Context) {
	logger.Info("[controller.SearchMenu] start...")

	/*
		수정내용
		카테고리 배열 변경으로 검색 조건 변경
	*/
	_, filter := SearchMenuAppendQuery(c) //검색 조건 쿼리 추가

	fmt.Println("[controller.SearchMenu] filter : ", filter)
	c.JSON(http.StatusOK, p.md.SearchMenu(filter))

}

// ViewMenu godoc
// @Summary call ViewMenu, return ok by json.
// @Description ViewMenu 메뉴 상세 - 주문자, 피주문자
// @name ViewMenu
// @Accept  json
// @Produce  json
// @Param MenuID query string true "메뉴 ID"
// @Router /oos/order/viewMenu [get]
// @Success 200 {object} model.Menu
func (p *Controller) ViewMenu(c *gin.Context) {
	logger.Info("[controller.ViewMenu] start...")

	menuId := c.Query("MenuID")
	c.JSON(http.StatusOK, p.md.ViewMenu(menuId))
}

// SetTodayMenu godoc
// @Summary call SetTodayMenu, return ok by json.
// @Description SetTodayMenu 오늘의 추천메뉴 여부 - 설정 변경 (메뉴ID를 기준으로 메뉴 업데이트)
// @name SetTodayMenu
// @Accept  json
// @Produce  json
// @Param Menu		body	model.Menu	true	"메뉴"
// @Router /oos/seller/setTodayMenu [put]
// @Success 200 {object} model.Menu
func (p *Controller) SetTodayMenu(c *gin.Context) {
	logger.Info("[controller.SetTodayMenu] start...")

	menuID := c.PostForm("MenuID")
	if menuID == "" {
		errMsg := "메뉴ID가 입력되지 않았습니다."
		logger.Error(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	menu, _ := UpdateMenuAppendQuery(c)
	fmt.Println("[controller.SetTodayMenu Param] menu : ", menu)

	updateFilter := bson.M{
		"$set": bson.M{
			"isTdoayMenu": menu.IsTdoayMenu,
		},
	}

	c.JSON(http.StatusOK, p.md.UpdateMenu(menuID, menu, updateFilter))
}

// SearchTodayMenu godoc
// @Summary call SearchTodayMenu, return ok by json.
// @Description SearchTodayMenu 오늘의 추천메뉴 리스트
// @name SearchTodayMenu
// @Accept  json
// @Produce  json
// @Param SellerID query  string false "판매자 ID"
// @Param MenuName query  string false "메뉴이름"
// @Param Price query  int false "가격"
// @Param CountryOf query  string false "원산지" Enums(한국, 일본, 중국)
// @Param Category query  string false "메뉴 카테고리" Enums(한식, 일식, 중식)
// @Param Status query  string false "주문 가능 상태" Enums(준비중, 판매중)
// @Param Spicy query  string false "맵기" Enums(아주매움, 매움, 보통, 순한맛)
// @Router /oos/order/searchTodayMenu [get]
// @Success 200 {object} Controller
func (p *Controller) SearchTodayMenu(c *gin.Context) {
	logger.Info("[controller.SearchMenu] start...")

	/*
		수정내용
		카테고리 배열 변경으로 검색 조건 변경
	*/
	_, filter := SearchMenuAppendQuery(c)

	//오늘의 추천메뉴 true 조회
	filter["isTdoayMenu"] = true

	fmt.Println("[controller.SetTodayMenu filter]", filter)
	c.JSON(http.StatusOK, p.md.SearchMenu(filter))
}

// NewOrder godoc
// @Summary call NewOrder, return ok by json.
// @Description NewOrder 주문 등록 - 주문자
// @name NewOrder
// @Accept  json
// @Produce  json
// @Param	OrdererMenuLink	body	model.OrdererMenuLink	true	"오더"
// @Router /oos/order/newOrder [post]
// @Success 200 {object} model.OrdererMenuLink
func (p *Controller) NewOrder(c *gin.Context) {
	logger.Info("[controller.NewOrder] start...")

	var params model.OrdererMenuLink
	if err := c.ShouldBind(&params); err == nil {

		menuID := params.MenuID
		if menuID == "" { //메뉴ID체크
			logger.Info("[controller.NewOrder] menuID...", menuID)
			errMsg := "등록된 메뉴 ID를 입력해주세요."
			logger.Error(errMsg)
			c.JSON(http.StatusBadRequest, errMsg)
			return
		}
		//주문 상태를 확인해서 취소 가능 상태 제어
		menu := p.md.ViewMenu(menuID)
		if menu.IsRecommeded {
			c.JSON(http.StatusBadRequest, "주문할 수 없는 메뉴 입니다.")
			return
		} else if menu.Status == "판매완료" {
			c.JSON(http.StatusBadRequest, "완판되었습니다.")
			return
		}

		params.MenuName = menu.MenuName
		params.SellerID = menu.SellerID

		c.JSON(http.StatusOK, p.md.NewOrder(params))

	} else {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// OrderStatus godoc
// @Summary call OrderStatus, return ok by json.
// @Description OrderStatus 주문 내역 조회 - 피주문자 (판매자ID 기준으로 검색)
// @name OrderStatus
// @Accept  json
// @Produce  json
// @Param SellerID query string true "판매자 ID"
// @Param OrderStatus query string false "주문 상태" Enums(주문확인중,조리중,배달중,배달완료,주문취소)
// @Router /oos/seller/orderStatus [get]
// @Success 200 {object} Controller
func (p *Controller) OrderStatus(c *gin.Context) {
	logger.Info("[controller.OrderList] start...")

	orderStatus := c.Query("OrderStatus")
	sellerID := c.Query("SellerID")

	fmt.Println("[controller.OrderStatus Param]", orderStatus, sellerID)
	c.JSON(http.StatusOK, gin.H{"주문내역 리스트 ": p.md.OrderStatus(orderStatus, sellerID)})

}

// ViewOrder godoc
// @Summary call ViewOrder, return ok by json.
// @Description ViewOrder 주문 상세 - 주문자, 피주문자
// @name ViewOrder
// @Accept  json
// @Produce  json
// @Param OrderNo query string true "주문번호"
// @Router /oos/order/viewOrder [get]
// @Success 200 {object} Controller
func (p *Controller) ViewOrder(c *gin.Context) {
	logger.Info("[controller.ViewOrder] start...")

	orderNO := c.Query("OrderNo")
	fmt.Println("[controller.ViewOrder Param]", orderNO)
	c.JSON(http.StatusOK, p.md.ViewOrder(orderNO))
}

// ChangeOrder godoc
// @Summary call ChangeOrder, return ok by json.
// @Description ChangeOrder 주문 변경 - 주문자 (주문변경 커멘드 Enums(주문추가, 주문취소, 정보변경))
// @name ChangeOrder
// @Accept  json
// @Produce  json
// @Param	OrdererMenuLink	body	model.OrdererMenuLink	true	"오더"
// @Router /oos/order/changeOrder [put]
// @Success 200 {object} Controller
func (p *Controller) ChangeOrder(c *gin.Context) {
	logger.Info("[controller.ChangeOrder] start...")

	var params model.BindChangeOrderState

	if err := c.ShouldBind(&params); err == nil {

		orderNO := params.OrderNo
		changeOrderCmd := params.ChangeOrderCmd

		fmt.Println("changeOrderCmd", changeOrderCmd)
		omLink := p.md.ViewOrder(orderNO)
		fmt.Println("OrdererMenuLink", omLink.OrderStatus)

		//주문 상태를 확인해서 취소 가능 상태 제어 (미구현)
		if changeOrderCmd == "주문취소" {
			if omLink.OrderStatus == "조리중" ||
				omLink.OrderStatus == "배달중" ||
				omLink.OrderStatus == "배달완료" {
				c.JSON(http.StatusBadRequest, "주문을 취소할 수 없는 상태입니다.")
				return
			}

			omLink.OrderStatus = "주문취소"

		} else if changeOrderCmd == "주문추가" {
			if omLink.OrderStatus == "배달중" ||
				omLink.OrderStatus == "배달완료" {
				c.JSON(http.StatusBadRequest, "배달 중이어서 주문을 추가할 수 없는 상태입니다.")
				return
			}
			omLink.OrderStatus = "주문확인중"
		}
		fmt.Println("OrdererMenuLink", omLink.OrderStatus)

		omLink.OrdererAddress = params.OrdererAddress
		omLink.OrdererPhone = params.OrdererPhone

		c.JSON(http.StatusOK, p.md.ChangeOrder(omLink))

	} else {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// SearchOrder godoc
// @Summary call SearchOrder, return ok by json.
// @Description SearchOrder 주문 내역 조회 기능 - 주문자
// @name SearchOrder
// @Accept  json
// @Produce  json
// @Param OrdererID query string true "주문자 ID"
// @Param MenuName query string false "메뉴이름"
// @Param OrderStatus query string false "주문 상태" Enums( 주문확인중,조리중,배달중,배달완료,주문취소)
// @Router /oos/order/searchOrder [get]
// @Success 200 {object} Controller
func (p *Controller) SearchOrder(c *gin.Context) {
	logger.Info("[controller.SearchOrder] start...", c.PostForm("OrdererID"))

	var params model.OrdererMenuLink

	params.OrdererID = c.Query("OrdererID")
	params.MenuName = c.Query("MenuName")
	params.OrderStatus = c.Query("OrderStatus")

	fmt.Println("[model.SearchOrder params] ", params)

	c.JSON(http.StatusOK, p.md.SearchOrder(params))

	// if err := c.ShouldBind(&params); err == nil {
	// 	c.JSON(http.StatusOK, p.md.SearchOrder(params))
	// } else {
	// 	logger.Error(err)
	// 	c.JSON(http.StatusBadRequest, "ERROR")
	// }
}

// CreateReview godoc
// @Summary call CreateReview, return ok by json.
// @Description CreateReview 리뷰 등록 - 주문자 (주문번호 기준으로 등록)
// @name CreateReview
// @Accept  json
// @Produce  json
// @Param	OrdererMenuLink	body	model.OrdererMenuLink	true	"오더"
// @Router /oos/order/createReview [post]
// @Success 200 {object} model.OrdererMenuLink
func (p *Controller) CreateReview(c *gin.Context) {
	logger.Info("[controller.CreateReview] start...")

	var params model.OrdererMenuLink
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.CreateReview(params))
	} else {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}
