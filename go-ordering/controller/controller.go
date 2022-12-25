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
// @Param SellerID path string true "판매자 ID"
// @Param SellerName path string true "판매자 이름"
// @Param MenuName path string true "메뉴이름"
// @Param Status path string true "주문 가능 상태" Enums(준비중, 판매중)
// @Param MaxCount path int true "판매 가능 갯수" mininum(1) maxinum(50)
// @Param CountryOf path string true "원산지" Enums(한국, 일본, 중국)
// @Param Price path int true "가격"
// @Param Spicy path string true "맵기" Enums(아주매움, 매움, 보통, 순한맛)
// @Param IsDisabled path bool true "판매여부" default(true)
// @Param TodayMenu path bool false "오늘의 추천메뉴 여부" default(false)
// @Param Category path string true "메뉴 카테고리" Enums(한식, 일식, 중식)
// @Router /oos/seller/createMenu [post]
// @Success 200 {object} model.Menu
func (p *Controller) CreateMenu(c *gin.Context) {
	logger.Info("[controller.CreateMenu] start...")

	//메뉴 등록시 판매자 로그인 필수.
	errChk, errMsg := checkCreateMenu(c)
	if errChk {
		logger.Error(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
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
// @Param MenuID path string true "메뉴 ID"
// @Param SellerID path string true "판매자 ID"
// @Param SellerName path string true "판매자 이름"
// @Param MenuName path string true "메뉴이름"
// @Param Status path string true "주문 가능 상태" Enums(준비중, 판매중)
// @Param MaxCount path int true "판매 가능 갯수" mininum(1) maxinum(50)
// @Param CountryOf path string true "원산지" Enums(한국, 일본, 중국)
// @Param Price path int true "가격"
// @Param Spicy path string true "맵기" Enums(아주매움, 매움, 보통, 순한맛)
// @Param IsDisabled path bool true "판매여부" default(true)
// @Param TodayMenu path bool false "오늘의 추천메뉴 여부" default(false)
// @Param Category path string true "메뉴 카테고리" Enums(한식, 일식, 중식)
// @Router /oos/seller/updateMenu [post]
// @Success 200 {object} model.Menu
func (p *Controller) UpdateMenu(c *gin.Context) {
	logger.Info("[controller.UpdateMenu] start...")

	menu, updateFilter := UpdateMenuAppendQuery(c)
	c.JSON(http.StatusOK, p.md.UpdateMenu(menu, updateFilter))
}

// DeleteMenu godoc
// @Summary call DeleteMenu, return ok by json.
// @Description DeleteMenu 메뉴 삭제 - 피주문자 (판매여부 bool 설정변경)
// @name DeleteMenu
// @Accept  json
// @Produce  json
// @Param MenuID path string true "메뉴 ID"
// @Param IsDisabled path bool true "판매여부"
// @Router /oos/seller/deleteMenu [put]
// @Success 200 {object} model.Menu
func (p *Controller) DeleteMenu(c *gin.Context) {
	logger.Info("[controller.DeleteMenu] start...")

	menuID := c.PostForm("MenuID")
	isDisabled := c.PostForm("IsDisabled")

	fmt.Println("[controller.DeleteMenu Param]", menuID, isDisabled)
	c.JSON(http.StatusOK, gin.H{"판매불가 설정되었습니다.": p.md.DeleteMenu(menuID, isDisabled)})
	// 	c.JSON(http.StatusOK, p.md.DeleteMenu(params))
}

// SearchMenu godoc
// @Summary call SearchMenu, return ok by json.
// @Description SearchMenu 메뉴 검색 - 주문자, 피주문자
// @name SearchMenu
// @Accept  json
// @Produce  json
// @Param MenuName path string true "메뉴이름"
// @Param Price path int true "가격"
// @Param CountryOf path string true "원산지" Enums(한국, 일본, 중국)
// @Param Category path string true "메뉴 카테고리" Enums(한식, 일식, 중식)
// @Param Status path string true "주문 가능 상태" Enums(준비중, 판매중)
// @Param Spicy path string true "맵기" Enums(아주매움, 매움, 보통, 순한맛)
// @Param TodayMenu path bool false "오늘의 추천메뉴 여부"
// @Router /oos/order/searchMenu [post]
// @Success 200 {object} []model.Menu
func (p *Controller) SearchMenu(c *gin.Context) {
	logger.Info("[controller.SearchMenu] start...")

	var filter bson.D
	_, filter = SearchMenuAppendQuery(c, filter) //검색 조건 쿼리 추가
	fmt.Println("[controller.SearchMenu] filter : ", filter)
	c.JSON(http.StatusOK, p.md.SearchMenu(filter))

	// var params model.Menu
	// if err := c.ShouldBind(&params); err == nil {
	// 	c.JSON(http.StatusOK, p.md.SearchMenu(params))
	// } else {
	// 	logger.Error(err)
	// 	c.JSON(http.StatusBadRequest, "ERROR")
	//}
}

// ViewMenu godoc
// @Summary call ViewMenu, return ok by json.
// @Description ViewMenu 메뉴 상세 - 주문자, 피주문자
// @name ViewMenu
// @Accept  json
// @Produce  json
// @Param MenuID path string true "메뉴 ID"
// @Router /oos/order/viewMenu [post]
// @Success 200 {object} model.Menu
func (p *Controller) ViewMenu(c *gin.Context) {
	logger.Info("[controller.ViewMenu] start...")

	menuId := c.PostForm("MenuId")
	c.JSON(http.StatusOK, p.md.ViewMenu(menuId))
}

// SetTodayMenu godoc
// @Summary call SetTodayMenu, return ok by json.
// @Description SetTodayMenu 오늘의 추천메뉴 여부 - 설정 변경 (메뉴ID를 기준으로 메뉴 업데이트)
// @name SetTodayMenu
// @Accept  json
// @Produce  json
// @Param MenuID path string true "메뉴 ID"
// @Param TodayMenu path bool true "오늘의 추천메뉴 여부"
// @Router /oos/seller/setTodayMenu [post]
// @Success 200 {object} model.Menu
func (p *Controller) SetTodayMenu(c *gin.Context) {
	logger.Info("[controller.SetTodayMenu] start...")
	menu, updateFilter := UpdateMenuAppendQuery(c)
	fmt.Println("[controller.SetTodayMenu Param] menu : ", menu)
	fmt.Println("[controller.SetTodayMenu Param] updateFilter : ", updateFilter)
	c.JSON(http.StatusOK, p.md.UpdateMenu(menu, updateFilter))
}

// SearchTodayMenu godoc
// @Summary call SearchTodayMenu, return ok by json.
// @Description SearchTodayMenu 오늘의 추천메뉴 리스트
// @name SearchTodayMenu
// @Accept  json
// @Produce  json
// @Param SellerID path string true "판매자 ID"
// @Router /oos/order/searchTodayMenu [post]
// @Success 200 {object} []model.Menu
func (p *Controller) SearchTodayMenu(c *gin.Context) {
	logger.Info("[controller.SearchMenu] start...")

	var filter bson.D
	_, filter = SearchMenuAppendQuery(c, filter)

	//오늘의 추천메뉴 true 조회
	filter = append(filter, bson.E{"todayMenu", true})

	fmt.Println("[controller.SetTodayMenu filter]", filter)
	c.JSON(http.StatusOK, p.md.SearchMenu(filter))
}

// NewOrder godoc
// @Summary call NewOrder, return ok by json.
// @Description NewOrder 주문 등록 - 주문자
// @name NewOrder
// @Accept  json
// @Produce  json
// @Param MenuID path string true "메뉴 ID"
// @Param MenuName path string true "메뉴이름"
// @Param OrdererID path string true "주문자 ID"
// @Param OrdererAddress path string false "주문자 주소"
// @Param OrdererPhone path int false "주문자 폰번호"
// @Router /oos/order/newOrder [post]
// @Success 200 {object} model.OrdererMenuLink
func (p *Controller) NewOrder(c *gin.Context) {
	logger.Info("[controller.NewOrder] start...")

	menuID := c.PostForm("MenuID")

	//주문 상태를 확인해서 취소 가능 상태 제어
	menu := p.md.ViewMenu(menuID)
	if menu.IsDisabled {
		c.JSON(http.StatusBadRequest, "주문할 수 없는 메뉴 입니다.")
		return
	} else if menu.Status == "오늘판매완료" {
		c.JSON(http.StatusBadRequest, "완판되었습니다.")
		return
	}

	var params model.OrdererMenuLink
	if err := c.ShouldBind(&params); err == nil {
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
// @Param SellerID path string true "판매자 ID"
// @Param OrderStatus path string true "주문 상태" Enums(주문확인중 - 조리중 - 배달중 - 배달완료 - 주문취소)
// @Router /oos/seller/OrderStatus [post]
// @Success 200 {object} []model.OrdererMenuLink
func (p *Controller) OrderStatus(c *gin.Context) {
	logger.Info("[controller.OrderList] start...")

	orderStatus := c.PostForm("OrderStatus")
	sellerID := c.PostForm("SellerID")

	fmt.Println("[controller.OrderStatus Param]", orderStatus, sellerID)
	c.JSON(http.StatusOK, gin.H{"주문내역 리스트 ": p.md.OrderStatus(orderStatus, sellerID)})

}

// ViewOrder godoc
// @Summary call ViewOrder, return ok by json.
// @Description ViewOrder 주문 상세 - 주문자, 피주문자
// @name ViewOrder
// @Accept  json
// @Produce  json
// @Param OrderNo path string true "주문번호"
// @Router /oos/order/viewOrder [post]
// @Success 200 {object} model.OrdererMenuLink
func (p *Controller) ViewOrder(c *gin.Context) {
	logger.Info("[controller.ViewOrder] start...")

	orderNO := c.PostForm("OrderNO")
	fmt.Println("[controller.ViewOrder Param]", orderNO)
	c.JSON(http.StatusOK, p.md.ViewOrder(orderNO))
}

// ChangeOrder godoc
// @Summary call ChangeOrder, return ok by json.
// @Description ChangeOrder 주문 변경 - 주문자 (주문변경 커멘드 Enums(주문추가, 주문취소, 정보변경))
// @name ChangeOrder
// @Accept  json
// @Produce  json
// @Param ChangeOrderCmd path string true "주문변경 커멘드" Enums(주문추가, 주문취소, 정보변경)
// @Param OrderNo path string true "주문번호"
// @Param OrdererAddress path string false "주문자 주소"
// @Param OrdererPhone path int false "주문자 폰번호"
// @Router /oos/order/changeOrder [post]
// @Success 200 {object} model.OrdererMenuLink
func (p *Controller) ChangeOrder(c *gin.Context) {
	logger.Info("[controller.ChangeOrder] start...")

	orderNO := c.PostForm("OrderNO")
	changeOrderCmd := c.PostForm("ChangeOrderCmd")

	fmt.Println("changeOrderCmd", changeOrderCmd)
	OrdererMenuLink := p.md.ViewOrder(orderNO)
	fmt.Println("OrdererMenuLink", OrdererMenuLink.OrderStatus)

	//주문 상태를 확인해서 취소 가능 상태 제어 (미구현)
	// if changeOrderCmd == "주문취소" {
	// 	if OrdererMenuLink.OrderStatus == "조리중" ||
	// 		OrdererMenuLink.OrderStatus == "배달중" ||
	// 		OrdererMenuLink.OrderStatus == "배달완료" {
	// 		c.JSON(http.StatusBadRequest, "주문을 취소할 수 없는 상태입니다.")
	// 		return
	// 	}
	// } else if changeOrderCmd == "주문추가" {
	// 	if OrdererMenuLink.OrderStatus == "배달중" ||
	// 		OrdererMenuLink.OrderStatus == "배달완료" {
	// 		c.JSON(http.StatusBadRequest, "배달 중이어서 주문을 추가할 수 없는 상태입니다.")
	// 		return
	// 	}
	// }

	var params model.OrdererMenuLink
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.ChangeOrder(params))
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
// @Param OrdererID path string true "주문자 ID"
// @Param MenuName path string true "메뉴이름"
// @Param OrderStatus path string true "주문 상태" Enums(주문확인중 - 조리중 - 배달중 - 배달완료 - 주문취소)
// @Router /oos/order/searchOrder [post]
// @Success 200 {object} []model.OrdererMenuLink
func (p *Controller) SearchOrder(c *gin.Context) {
	logger.Info("[controller.SearchOrder] start...")

	var params model.OrdererMenuLink
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.SearchOrder(params))
	} else {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// CreateReview godoc
// @Summary call CreateReview, return ok by json.
// @Description CreateReview 리뷰 등록 - 주문자 (주문번호 기준으로 등록)
// @name CreateReview
// @Accept  json
// @Produce  json
// @Param OrderNo path string true "주문번호"
// @Param OrderComment path string true "후기"
// @Param OrderStarGrade path int true "평점" mininum(1) maxinum(5)
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
