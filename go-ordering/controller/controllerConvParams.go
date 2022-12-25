package controller

import (
	"WBABEProject-20/go-ordering/logger"
	"WBABEProject-20/go-ordering/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Description Menu 구조체 파라메터 매핑
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
		fmt.Println("fmt.Println(c.PostForm(Price))", c.PostForm("Price"))
		if v, err := strconv.Atoi(c.PostForm("Price")); err == nil {
			fmt.Println("v", v)
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

// @Description OrdererMenuLink 구조체 파라메터 매핑
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

// @Description 메뉴 등록시 판매자 ID 체크 (로그인 여부)
func checkCreateMenu(c *gin.Context) (bool, string) {

	var errChk = false
	var errMsg = ""

	//메뉴 등록시 판매자 로그인 필수.
	if c.PostForm("SellerID") == "" {
		errMsg = "로그인해주세요. (판매자 ID (SellerID)는 필수 입니다.)"
		errChk = true
	}

	return errChk, errMsg
}

// @Description 메뉴 검색 - 속성 값 체크해서 검색 조건 리턴
func SearchMenuAppendQuery(c *gin.Context, filter bson.D) (model.Menu, bson.D) {

	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {

		if c.PostForm("MenuID") != "" {
			filter = append(filter, bson.E{"menuID", params.MenuID})
		}
		if c.PostForm("SellerID") != "" {
			filter = append(filter, bson.E{"sellerID", params.SellerID})
		}
		if c.PostForm("SellerName") != "" {
			filter = append(filter, bson.E{"sellerName", params.SellerName})
		}
		if c.PostForm("MenuName") != "" {
			filter = append(filter, bson.E{"menuName", params.MenuName})
		}
		if c.PostForm("Status") != "" {
			filter = append(filter, bson.E{"status", params.Status})
		}
		if c.PostForm("MaxCount") != "" {
			filter = append(filter, bson.E{"maxCount", params.MaxCount})
		}
		if c.PostForm("CountryOf") != "" {
			filter = append(filter, bson.E{"countryOf", params.CountryOf})
		}
		if c.PostForm("Price") != "" {
			filter = append(filter, bson.E{"price", params.Price})

		}
		if c.PostForm("Spicy") != "" {
			filter = append(filter, bson.E{"spicy", params.Spicy})
		}
		if c.PostForm("Popularity") != "" {
			filter = append(filter, bson.E{"popularity", params.Popularity})

		}
		if c.PostForm("IsDisabled") != "" {
			filter = append(filter, bson.E{"isDisabled", params.IsDisabled})

		}
		if c.PostForm("TodayMenu") != "" {
			filter = append(filter, bson.E{"todayMenu", params.TodayMenu})

		}
		if c.PostForm("Category") != "" {
			filter = append(filter, bson.E{"category", params.Category})
		}
	} else {
		logger.Error(err)
	}
	return params, filter
}

// @Description 메뉴 검색 - 속성 값 체크해서 검색 조건 리턴
func UpdateMenuAppendQuery(c *gin.Context) (model.Menu, bson.M) {

	filter := bson.M{
		"$set": bson.M{}}
	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {

		//ID는 업데이트 하지 않음
		// if c.PostForm("MenuID") != "" {
		// 	filter["$set"].(bson.M)["menuID"] = params.MenuID
		// }
		// if c.PostForm("SellerID") != "" {
		// 	filter["$set"].(bson.M)["sellerID"] = params.SellerID
		// }
		if c.PostForm("SellerName") != "" {
			filter["$set"].(bson.M)["sellerName"] = params.SellerName
		}
		if c.PostForm("MenuName") != "" {
			filter["$set"].(bson.M)["menuName"] = params.MenuName
		}
		if c.PostForm("Status") != "" {
			filter["$set"].(bson.M)["status"] = params.Status
		}
		if c.PostForm("MaxCount") != "" {
			filter["$set"].(bson.M)["maxCount"] = params.MaxCount
		}
		if c.PostForm("CountryOf") != "" {
			filter["$set"].(bson.M)["countryOf"] = params.CountryOf
		}
		if c.PostForm("Price") != "" {
			filter["$set"].(bson.M)["price"] = params.Price

		}
		if c.PostForm("Spicy") != "" {
			filter["$set"].(bson.M)["spicy"] = params.Spicy
		}
		if c.PostForm("Popularity") != "" {
			filter["$set"].(bson.M)["popularity"] = params.Popularity

		}
		if c.PostForm("IsDisabled") != "" {
			filter["$set"].(bson.M)["isDisabled"] = params.IsDisabled

		}
		if c.PostForm("TodayMenu") != "" {
			filter["$set"].(bson.M)["todayMenu"] = params.TodayMenu

		}
		if c.PostForm("Category") != "" {
			filter["$set"].(bson.M)["category"] = params.Category
		}
	} else {
		logger.Error(err)
	}
	return params, filter
}

// // @Description Seller 구조체 파라메터 매핑
// func setParamSeller(c *gin.Context) model.Seller {

// 	var params model.Seller

// 	if c.PostForm("SellerID") != "" {
// 		params.SellerID = c.PostForm("SellerID")
// 	}
// 	if c.PostForm("SellerName") != "" {
// 		params.SellerName = c.PostForm("SellerName")
// 	}
// 	if c.PostForm("Address") != "" {
// 		params.Address = c.PostForm("Address")
// 	}
// 	if c.PostForm("Phone") != "" {
// 		if v, err := strconv.Atoi(c.PostForm("Phone")); err == nil {
// 			params.Phone = v
// 		}
// 	}
// 	if c.PostForm("SellCount") != "" {
// 		if v, err := strconv.Atoi(c.PostForm("SellCount")); err == nil {
// 			params.SellCount = v
// 		}
// 	}

// 	return params
// }

// // @Description Order 구조체 파라메터 매핑
// func setParamOrderer(c *gin.Context) model.Orderer {

// 	var params model.Orderer

// 	if c.PostForm("OrdererID") != "" {
// 		params.OrdererID = c.PostForm("OrdererID")
// 	}
// 	if c.PostForm("OrderName") != "" {
// 		params.OrderName = c.PostForm("OrderName")
// 	}
// 	if c.PostForm("Address") != "" {
// 		params.Address = c.PostForm("Address")
// 	}
// 	if c.PostForm("Phone") != "" {
// 		if v, err := strconv.Atoi(c.PostForm("Phone")); err == nil {
// 			params.Phone = v
// 		}
// 	}
// 	if c.PostForm("OrderCount") != "" {
// 		if v, err := strconv.Atoi(c.PostForm("OrderCount")); err == nil {
// 			params.OrderCount = v
// 		}
// 	}
// 	return params
// }
