package controller

import (
	"WBABEProject-20/go-ordering/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setParamMenu(c *gin.Context) model.Menu {

	var params model.Menu

	if c.PostForm("MenuID") != "" {
		params.MenuID = c.PostForm("MenuID")
	}
	if c.PostForm("SellerID") != "" {
		params.SellerID = c.PostForm("SellerID")
	}
	if c.PostForm("SellerName") != "" {
		params.SellerName = c.PostForm("SellerName")
	}
	if c.PostForm("MenuName") != "" {
		params.MenuName = c.PostForm("MenuName")
	}
	if c.PostForm("Status") != "" {
		params.Status = c.PostForm("Status")
	}
	if c.PostForm("MaxCount") != "" {
		if v, err := strconv.Atoi(c.PostForm("MaxCount")); err == nil {
			params.MaxCount = v
		}
	}
	if c.PostForm("CountryOf") != "" {
		params.CountryOf = c.PostForm("CountryOf")
	}
	if c.PostForm("Price") != "" {
		if v, err := strconv.Atoi(c.PostForm("Price")); err == nil {
			params.Price = v
		}
	}
	if c.PostForm("Spicy") != "" {
		params.Spicy = c.PostForm("Spicy")
	}
	if c.PostForm("Popularity") != "" {
		if v, err := strconv.Atoi(c.PostForm("Popularity")); err == nil {
			params.Popularity = v
		}
	}
	if c.PostForm("IsDisabled") != "" {
		if v, err := strconv.ParseBool(c.PostForm("IsDisabled")); err == nil {
			params.IsDisabled = v
		}
	}
	if c.PostForm("TodayMenu") != "" {
		if v, err := strconv.ParseBool(c.PostForm("TodayMenu")); err == nil {
			params.TodayMenu = v
		}
	}
	if c.PostForm("Category") != "" {
		params.Category = c.PostForm("Category")
	}

	return params
}

func setParamOrdererMenuLink(c *gin.Context) model.OrdererMenuLink {

	var params model.OrdererMenuLink

	if c.PostForm("OrderNo") != "" {
		params.OrderNo = c.PostForm("OrderNo")
	}
	if c.PostForm("MenuID") != "" {
		params.MenuID = c.PostForm("MenuID")
	}
	if c.PostForm("OrdererID") != "" {
		params.OrdererID = c.PostForm("OrdererID")
	}
	if c.PostForm("MenuName") != "" {
		params.MenuName = c.PostForm("MenuName")
	}
	if c.PostForm("OrderStarGrade") != "" {
		if v, err := strconv.Atoi(c.PostForm("OrderStarGrade")); err == nil {
			params.OrderStarGrade = v
		}
	}
	if c.PostForm("OrderComment") != "" {
		params.OrderComment = c.PostForm("OrderComment")
	}
	if c.PostForm("OrderStatus") != "" {
		params.OrderStatus = c.PostForm("OrderStatus")
	}
	if c.PostForm("OrdererAddress") != "" {
		params.OrdererAddress = c.PostForm("OrdererAddress")
	}
	if c.PostForm("OrdererPhone") != "" {
		if v, err := strconv.Atoi(c.PostForm("OrdererPhone")); err == nil {
			params.OrdererPhone = v
		}
	}
	return params
}

func setParamSeller(c *gin.Context) model.Seller {

	var params model.Seller

	if c.PostForm("SellerID") != "" {
		params.SellerID = c.PostForm("SellerID")
	}
	if c.PostForm("SellerName") != "" {
		params.SellerName = c.PostForm("SellerName")
	}
	if c.PostForm("Address") != "" {
		params.Address = c.PostForm("Address")
	}
	if c.PostForm("Phone") != "" {
		if v, err := strconv.Atoi(c.PostForm("Phone")); err == nil {
			params.Phone = v
		}
	}
	if c.PostForm("SellCount") != "" {
		if v, err := strconv.Atoi(c.PostForm("SellCount")); err == nil {
			params.SellCount = v
		}
	}

	return params
}

func setParamOrderer(c *gin.Context) model.Orderer {

	var params model.Orderer

	if c.PostForm("OrdererID") != "" {
		params.OrdererID = c.PostForm("OrdererID")
	}
	if c.PostForm("OrderName") != "" {
		params.OrderName = c.PostForm("OrderName")
	}
	if c.PostForm("Address") != "" {
		params.Address = c.PostForm("Address")
	}
	if c.PostForm("Phone") != "" {
		if v, err := strconv.Atoi(c.PostForm("Phone")); err == nil {
			params.Phone = v
		}
	}
	if c.PostForm("OrderCount") != "" {
		if v, err := strconv.Atoi(c.PostForm("OrderCount")); err == nil {
			params.OrderCount = v
		}
	}
	return params
}
