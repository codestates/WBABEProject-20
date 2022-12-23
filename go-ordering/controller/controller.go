package controller

//controller.go : 실제 비지니스 로직 및 프로세스가 처리후 결과 전송
import (
	"WBABEProject-20/go-ordering/logger"
	"WBABEProject-20/go-ordering/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	fmt.Println("Controller.NewCTL r : ", r)
	fmt.Println("Controller.NewCTL rep : ", rep)
	return r, nil
}

// CreateMenu godoc
// @Summary call CreateMenu, return ok by json.
// @Description CreateMenu 메뉴 등록 - 피주문자
// @name CreateMenu
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/seller/createMenu [post]
// @Success 200 {object} Controller
func (p *Controller) CreateMenu(c *gin.Context) {
	logger.Info("controller.CreateMenu start...")
	fmt.Println("controller.CreateMenu start2...")

	// var params1 model.BodyMenu
	// var params2 model.BodyMenu
	// fmt.Println("ShouldBind : ", c.ShouldBind(&params1))
	// fmt.Println("ShouldBindJSON : ", c.ShouldBindJSON(&params2))

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.CreateMenu(params))
	} else {
		fmt.Println("[controller.CreateMenu]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// UpdateMenu godoc
// @Summary call UpdateMenu, return ok by json.
// @Description UpdateMenu 메뉴 수정 - 피주문자
// @name UpdateMenu
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/seller/updateMenu [post]
// @Success 200 {object} Controller
func (p *Controller) UpdateMenu(c *gin.Context) {
	logger.Info("[controller.UpdateMenu] start...")
	fmt.Println("[controller.UpdateMenu] start...")

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.UpdateMenu(params))
	} else {
		logger.Error("controller.UpdateMenu start...")
		fmt.Println("[controller.UpdateMenu]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// DeleteMenu godoc
// @Summary call DeleteMenu, return ok by json.
// @Description DeleteMenu 메뉴 삭제 - 피주문자
// @name DeleteMenu
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/seller/deleteMenu [post]
// @Success 200 {object} Controller
func (p *Controller) DeleteMenu(c *gin.Context) {
	logger.Info("[controller.DeleteMenu] start...")
	fmt.Println("[controller.DeleteMenu] start...")

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.DeleteMenu(params))
	} else {
		logger.Error("controller.DeleteMenu start...")
		fmt.Println("[controller.DeleteMenu]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// OrderStates godoc
// @Summary call OrderStates, return ok by json.
// @Description OrderStates 주문 내역 조회 - 피주문자
// @name OrderStates
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/seller/orderStates [post]
// @Success 200 {object} Controller
func (p *Controller) OrderStates(c *gin.Context) {
	logger.Info("[controller.OrderList] start...")
	fmt.Println("[controller.OrderList] start...")

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.OrderStates(params))
	} else {
		logger.Error("controller.OrderList start...")
		fmt.Println("[controller.OrderList]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// SearchMenu godoc
// @Summary call SearchMenu, return ok by json.
// @Description SearchMenu 메뉴 검색 - 주문자, 피주문자
// @name SearchMenu
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/order/searchMenu [post]
// @Success 200 {object} Controller
func (p *Controller) SearchMenu(c *gin.Context) {
	logger.Info("[controller.SearchMenu] start...")
	fmt.Println("[controller.SearchMenu] start...")

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.SearchMenu(params))
	} else {
		logger.Error("controller.SearchMenu start...")
		fmt.Println("[controller.SearchMenu]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// ViewMenu godoc
// @Summary call ViewMenu, return ok by json.
// @Description ViewMenu 메뉴 상세 - 주문자, 피주문자
// @name ViewMenu
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/order/viewMenu [post]
// @Success 200 {object} Controller
func (p *Controller) ViewMenu(c *gin.Context) {
	logger.Info("[controller.ViewMenu] start...")
	fmt.Println("[controller.ViewMenu] start...")

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.ViewMenu(params))
	} else {
		logger.Error("controller.ViewMenu start...")
		fmt.Println("[controller.ViewMenu]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// CreateReview godoc
// @Summary call CreateReview, return ok by json.
// @Description CreateReview 리뷰 등록 - 주문자
// @name CreateReview
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/order/createReview [post]
// @Success 200 {object} Controller
func (p *Controller) CreateReview(c *gin.Context) {
	logger.Info("[controller.CreateReview] start...")
	fmt.Println("[controller.CreateReview] start...")

	var params model.OrdererMenuLink
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.CreateReview(params))
	} else {
		logger.Error("controller.CreateReview start...")
		fmt.Println("[controller.CreateReview]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// NewOrder godoc
// @Summary call NewOrder, return ok by json.
// @Description NewOrder 주문 등록 - 주문자
// @name NewOrder
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/order/newOrder [post]
// @Success 200 {object} Controller
func (p *Controller) NewOrder(c *gin.Context) {
	logger.Info("[controller.NewOrder] start...")
	fmt.Println("[controller.NewOrder] start...")

	var params model.OrdererMenuLink
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.NewOrder(params))
	} else {
		logger.Error("controller.NewOrder start...")
		fmt.Println("[controller.NewOrder]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// ChangeOrder godoc
// @Summary call ChangeOrder, return ok by json.
// @Description ChangeOrder 주문 변경 - 주문자
// @name ChangeOrder
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/order/changeOrder [post]
// @Success 200 {object} Controller
func (p *Controller) ChangeOrder(c *gin.Context) {
	logger.Info("[controller.ChangeOrder] start...")
	fmt.Println("[controller.ChangeOrder] start...")

	var params model.OrdererMenuLink
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.ChangeOrder(params))
	} else {
		logger.Error("controller.ChangeOrder start...")
		fmt.Println("[controller.ChangeOrder]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}

// SearchOrder godoc
// @Summary call SearchOrder, return ok by json.
// @Description SearchOrder 주문 내역 조회 기능 - 주문자
// @name SearchOrder
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /oos/order/searchOrder [post]
// @Success 200 {object} Controller
func (p *Controller) SearchOrder(c *gin.Context) {
	logger.Info("[controller.SearchOrder] start...")
	fmt.Println("[controller.SearchOrder] start...")

	var params model.OrdererMenuLink
	if err := c.ShouldBind(&params); err == nil {
		c.JSON(http.StatusOK, p.md.SearchOrder(params))
	} else {
		logger.Error("controller.SearchOrder start...")
		fmt.Println("[controller.SearchOrder]", err)
		c.JSON(http.StatusBadRequest, "ERROR")
	}
}
