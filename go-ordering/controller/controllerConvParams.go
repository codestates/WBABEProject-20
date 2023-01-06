package controller

import (
	"WBABEProject-20/go-ordering/logger"
	"WBABEProject-20/go-ordering/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
)

// @Description 메뉴 등록시 판매자 ID 체크 (로그인 여부)
func checkCreateMenu(param model.Menu, user model.UserAccount) (bool, string) {
	logger.Info("[controllerConvParams checkCreateMenu] SellerID : ", param.SellerID)
	logger.Info("[controllerConvParams checkCreateMenu] MenuName : ", param.MenuName)
	errChk := false
	errMsg := ""

	/*
		판매자 아이디를 체크한다고 되어있는데, 로직을 살펴보면 빈 스트링 값만 아니면 에러가 발생하지 않고 통과되네요.
		예를들면 SellerId 값으로 'GSODEIGNIOGNRO' 같은 값을 넣어도 판매자라고 인식을 하게 되네요.
		권한 관리의 경우 미들웨어를 통해서 유저의 타입이 판매자인지, 구매자인지를 체크하는 로직들로 제어를 하는 것이 올바르다고 생각합니다.

		gin 권한관리, permission middleware와 같은 키워드로 검색해보시기 바랍니다.
	*/
	/*
		수정내용
		Seller와 Order를 User로 합치고, UserType으로 판매자 구매자를 구분
	*/
	if user.UserType != "판매자" {
		errMsg = "로그인 유저가 판매자가 아닌 경우 메뉴를 만들 수 없습니다."
		errChk = true
	}
	//메뉴 등록시 판매자 로그인 필수.
	if param.SellerID == "" {
		errMsg = "로그인해주세요. (판매자 ID (SellerID)는 필수 입니다.)"
		errChk = true
	}
	if param.MenuName == "" {
		errMsg = "메뉴 이름은 필수 입니다.)"
		errChk = true
	}

	return errChk, errMsg
}

/*
	수정내용
	카테고리 배열 변경으로 검색 조건 변경
*/
// @Description 메뉴 검색 - 속성 값 체크해서 검색 조건 리턴
func SearchMenuAppendQuery(c *gin.Context) (model.Menu, bson.M) {

	fmt.Println("[SearchMenuAppendQuery] atart ")

	filter := bson.M{}
	var params model.Menu
	if err := c.ShouldBind(&params); err == nil {

		if c.Query("MenuID") != "" {
			filter["menuID"] = params.MenuID
		}
		if c.Query("SellerID") != "" {
			filter["sellerID"] = params.SellerID
		}
		if c.Query("SellerName") != "" {
			filter["sellerName"] = params.SellerName
		}
		if c.Query("MenuName") != "" {
			filter["menuName"] = params.MenuName
		}
		if c.Query("Status") != "" {
			filter["status"] = params.Status
		}
		if c.Query("MaxCount") != "" {
			filter["maxCount"] = params.MaxCount
		}
		if c.Query("CountryOf") != "" {
			filter["countryOf"] = params.CountryOf
		}
		if c.Query("Price") != "" {
			filter["price"] = params.Price

		}
		if c.Query("Spicy") != "" {
			filter["spicy"] = params.Spicy
		}
		if c.Query("Popularity") != "" {
			filter["popularity"] = params.Popularity
		}
		if c.Query("IsRecommeded") != "" {
			filter["isRecommeded"] = params.IsRecommeded

		}
		if c.Query("IsTdoayMenu") != "" {
			filter["isTdoayMenu"] = params.IsTdoayMenu

		}
		if c.Query("Category") != "" {
			filter["category"] = bson.M{"$all": params.Category}
		}
	} else {
		logger.Error(err)
	}

	fmt.Println("[SearchMenuAppendQuery] params : ", params)
	fmt.Println("[SearchMenuAppendQuery] filter : ", filter)

	return params, filter
}

/*
수정내용
validator를 사용해 파라메터의 min, max 체크(controller에서 호출)
*/
func validationCheck(menu model.Menu) map[string]interface{} {
	v := validator.New()
	if err := v.Struct(menu); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		return resp
	}
	return nil
}

// @Description 메뉴 검색 - 속성 값 체크해서 검색 조건 리턴
func UpdateMenuAppendQuery(c *gin.Context) (model.Menu, bson.M) {

	/*
		1. Validator를 통해 Input 값을 제어하시면 좋을 것 같습니다.
			min, max 값 혹은 들어오는 데이터에 대해서 검증하려면 Gin에서 제공하는 validtor 기능을 이용하시면 좋을 것 같습니다.
			지금의 구조 같은 경우 판매 가능 갯수가 1~50이지만 그 외의 숫자가 들어와도 모두 가능하도록 되어 있습니다.

			아래의 링크를 참고해보시기 바랍니다.
			https://gin-gonic.com/docs/examples/custom-validators/
		2. 카테고리의 경우 보통은 0개 ~ 2개 이상의 값을 지닙니다.
			해당 구조에서 여러 카테고리에 속하는 경우는 어떻게 입력받을 수 있나요? 이 부분에 대한 고려가 필요해 보입니다.
	*/

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
		if params.SellerName != "" {
			filter["$set"].(bson.M)["sellerName"] = params.SellerName
		}
		if params.MenuName != "" {
			filter["$set"].(bson.M)["menuName"] = params.MenuName
		}
		if params.Status != "" {
			filter["$set"].(bson.M)["status"] = params.Status
		}
		if params.MaxCount > 0 {
			filter["$set"].(bson.M)["maxCount"] = params.MaxCount
		}
		if params.CountryOf != "" {
			filter["$set"].(bson.M)["countryOf"] = params.CountryOf
		}
		if params.Price > 0 {
			filter["$set"].(bson.M)["price"] = params.Price

		}
		if params.Spicy != "" {
			filter["$set"].(bson.M)["spicy"] = params.Spicy
		}
		if params.Popularity > 0 {
			filter["$set"].(bson.M)["popularity"] = params.Popularity

		}
		//별도 업데이트
		// if params.IsRecommeded != "" {
		// 	filter["$set"].(bson.M)["isRecommeded"] = params.IsRecommeded

		// }
		// if params.TodayMenu != "" {
		// 	filter["$set"].(bson.M)["todayMenu"] = params.TodayMenu

		// }
		/*
			수정내용
			메뉴 카테고리는 배열로 변경하여 0개 이상의 경우 수정
		*/
		if len(params.Category) > 0 {
			filter["$set"].(bson.M)["category"] = params.Category
		}
	} else {
		logger.Error(err)
	}

	return params, filter
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
