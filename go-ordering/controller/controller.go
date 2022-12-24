package controller

//controller.go : 실제 비지니스 로직 및 프로세스가 처리후 결과 전송
import (
	"WBABEProject-20/go-ordering/logger"
	"WBABEProject-20/go-ordering/model"
	"net/http"

	"github.com/gin-gonic/gin"
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
// @Param MenuName path string true "메뉴이름"
// @Param Price path int true "가격"
// @Param CountryOf path string true "원산지" Enums(한국, 일본, 중국)
// @Param Category path string true "메뉴 카테고리" Enums(한식, 일식, 중식)
// @Param Status path string true "주문 가능 상태" Enums(준비중, 판매중)
// @Param MaxCount path int true "판매 가능 갯수" mininum(1) maxinum(50)
// @Param Spicy path string true "맵기" Enums(아주매움, 매움, 보통, 순한맛)
// @Param IsDisabled path bool true "판매여부" default(true)
// @Param TodayMenu path bool false "오늘의 추천메뉴 여부" default(false)
// @Router /oos/seller/createMenu [post]
// @Success 200 {object} Menu
func (p *Controller) CreateMenu(c *gin.Context) {
	logger.Info("[controller.CreateMenu] start...")

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.CreateMenu(params))
	} else {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// UpdateMenu godoc
// @Summary call UpdateMenu, return ok by json.
// @Description UpdateMenu 메뉴 수정 - 피주문자
// @name UpdateMenu
// @Accept  json
// @Produce  json
// @Param MenuName path string true "메뉴이름"
// @Param Price path int true "가격"
// @Param CountryOf path string true "원산지" Enums(한국, 일본, 중국)
// @Param Category path string true "메뉴 카테고리" Enums(한식, 일식, 중식)
// @Param Status path string true "주문 가능 상태" Enums(준비중, 판매중)
// @Param MaxCount path int true "판매 가능 갯수" mininum(1) maxinum(50)
// @Param Spicy path string true "맵기" Enums(아주매움, 매움, 보통, 순한맛)
// @Param IsDisabled path bool true "판매여부"
// @Param TodayMenu path bool false "오늘의 추천메뉴 여부"
// @Router /oos/seller/updateMenu [post]
// @Success 200 {object} Menu
func (p *Controller) UpdateMenu(c *gin.Context) {
	logger.Info("[controller.UpdateMenu] start...")

	params := setParamMenu(c)
	c.JSON(http.StatusOK, p.md.UpdateMenu(params))

}

// DeleteMenu godoc
// @Summary call DeleteMenu, return ok by json.
// @Description DeleteMenu 메뉴 삭제 - 피주문자
// @name DeleteMenu
// @Accept  json
// @Produce  json
// @Param MenuName path string true "메뉴이름"
// @Param IsDisabled path bool true "판매여부"
// @Router /oos/seller/deleteMenu [put]
// @Success 200 {object} Menu
func (p *Controller) DeleteMenu(c *gin.Context) {
	logger.Info("[controller.DeleteMenu] start...")

	menuName := c.PostForm("MenuName")
	isDisabled := c.PostForm("IsDisabled")

	logger.Info("[controller.DeleteMenu Param]", menuName, isDisabled)
	c.JSON(http.StatusOK, gin.H{"Persons": p.md.DeleteMenu(menuName, isDisabled)})
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
// @Success 200 {object} []Menu
func (p *Controller) SearchMenu(c *gin.Context) {
	logger.Info("[controller.SearchMenu] start...")

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.SearchMenu(params))
	} else {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// ViewMenu godoc
// @Summary call ViewMenu, return ok by json.
// @Description ViewMenu 메뉴 상세 - 주문자, 피주문자
// @name ViewMenu
// @Accept  json
// @Produce  json
// @Param MenuName path string true "메뉴이름"
// @Router /oos/order/viewMenu [post]
// @Success 200 {object} Menu
func (p *Controller) ViewMenu(c *gin.Context) {
	logger.Info("[controller.ViewMenu] start...")

	menuName := c.PostForm("MenuName")
	c.JSON(http.StatusOK, gin.H{"Persons": p.md.ViewMenu(menuName)})
}

// SetTodayMenu godoc
// @Summary call SetTodayMenu, return ok by json.
// @Description SetTodayMenu 오늘의 추천메뉴 설정 변경
// @name SetTodayMenu
// @Accept  json
// @Produce  json
// @Param SellerID path string true "판매자 ID"
// @Param MenuName path string true "메뉴이름"
// @Param TodayMenu path bool true "오늘의 추천메뉴 여부"
// @Router /oos/seller/setTodayMenu [post]
// @Success 200 {object} Menu
func (p *Controller) SetTodayMenu(c *gin.Context) {
	logger.Info("[controller.SetTodayMenu] start...")

	// var params model.Menu
	// params.SellerID = c.PostForm("SellerID")
	// TodayMenuboolValue, err := strconv.ParseBool(c.PostForm("TodayMenu"))
	// if err != nil {
	// 	logger.Error(err)
	// }
	// params.TodayMenu = TodayMenuboolValue
	// params.IsDisabled = true
	// params.MenuName = c.PostForm("MenuName")

	params := setParamMenu(c)
	logger.Info("[controller.SetTodayMenu Param]", params)

	c.JSON(http.StatusOK, p.md.UpdateMenu(params))

}

// SearchTodayMenu godoc
// @Summary call SearchTodayMenu, return ok by json.
// @Description SearchTodayMenu 오늘의 추천메뉴 리스트
// @name SearchTodayMenu
// @Accept  json
// @Produce  json
// @Param SellerID path string true "판매자 ID"
// @Router /oos/order/searchTodayMenu [post]
// @Success 200 {object} []Menu
func (p *Controller) SearchTodayMenu(c *gin.Context) {
	logger.Info("[controller.SearchMenu] start...")

	var params model.Menu
	params.SellerID = c.PostForm("SellerID")
	params.TodayMenu = true
	params.IsDisabled = true

	logger.Info("[controller.SetTodayMenu Param]", params)

	c.JSON(http.StatusOK, p.md.SearchMenu(params))
}

// NewOrder godoc
// @Summary call NewOrder, return ok by json.
// @Description NewOrder 주문 등록 - 주문자
// @name NewOrder
// @Accept  json
// @Produce  json
// @Param MenuName path string true "메뉴이름"
// @Param OrdererID path string true "주문자 ID"
// @Param OrderStatus path string true "주문 상태" Enums(준비중, 주문취소, 배달중, 배달완료)
// @Param OrdererAddress path string false "주문자 주소"
// @Param OrdererPhone path int false "주문자 폰번호"
// @Router /oos/order/newOrder [post]
// @Success 200 {object} OrdererMenuLink
func (p *Controller) NewOrder(c *gin.Context) {
	logger.Info("[controller.NewOrder] start...")

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
// @Description OrderStatus 주문 내역 조회 - 피주문자
// @name OrderStatus
// @Accept  json
// @Produce  json
// @Param MenuName path string true "메뉴이름"
// @Param OrderStatus path string true "주문 상태" Enums(준비중, 주문취소, 배달중, 배달완료)
// @Router /oos/seller/OrderStatus [post]
// @Success 200 {object} []OrdererMenuLink
func (p *Controller) OrderStatus(c *gin.Context) {
	logger.Info("[controller.OrderList] start...")

	menuName := c.PostForm("MenuName")
	orderStatus := c.PostForm("OrderStatus")

	logger.Info("[controller.OrderStatus Param]", menuName, orderStatus)
	c.JSON(http.StatusOK, gin.H{"Persons": p.md.OrderStatus(menuName, orderStatus)})

}

// ChangeOrder godoc
// @Summary call ChangeOrder, return ok by json.
// @Description ChangeOrder 주문 변경 - 주문자 (수정/취소)
// @name ChangeOrder
// @Accept  json
// @Produce  json
// @Param MenuName path string true "메뉴이름"
// @Param OrdererID path string true "주문자 ID"
// @Param OrderStatus path string true "주문 상태" Enums(준비중, 주문취소, 배달중, 배달완료)
// @Param OrdererAddress path string false "주문자 주소"
// @Param OrdererPhone path int false "주문자 폰번호"
// @Router /oos/order/changeOrder [post]
// @Success 200 {object} OrdererMenuLink
func (p *Controller) ChangeOrder(c *gin.Context) {
	logger.Info("[controller.ChangeOrder] start...")

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
// @Param MenuName path string true "메뉴이름"
// @Param OrderStatus path string true "주문 상태" Enums(준비중, 주문취소, 배달중, 배달완료)
// @Router /oos/order/searchOrder [post]
// @Success 200 {object} []OrdererMenuLink
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
// @Description CreateReview 리뷰 등록 - 주문자
// @name CreateReview
// @Accept  json
// @Produce  json
// @Param MenuName path string true "메뉴이름"
// @Param OrdererID path string true "주문자 ID"
// @Param OrderComment path string true "후기"
// @Param OrderStarGrade path string true "평점"
// @Router /oos/order/createReview [post]
// @Success 200 {object} OrdererMenuLink
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
