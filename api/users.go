package api

import (
	"easy-search/models"
	r "easy-search/pkgs/dto"
	e "easy-search/pkgs/exception"
	"easy-search/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryUserPhone(c *gin.Context) {

	// request接收
	var getUserPhoneRequest serializers.UserPhoneRequest
	var user models.User
	err := c.Bind(&getUserPhoneRequest)
	if err != nil {
		e.CommonError(e.ErrInputInvalid, c)
		return
	}

	// 具体查询
	user.Phone = getUserPhoneRequest.Phone
	result, err := models.GetUserPhone(user)
	if err != nil {
		e.NotFoundError(err, c)
		return
	}

	// 构造返回体
	c.JSON(http.StatusOK, r.SuccessResponse{
		Status: "success",
		Data: serializers.UserResponse{
			User: serializers.SerializeUser(*result),
		},
	})
}

func QueryUserQq(c *gin.Context) {

	// request接收
	var getUserQqRequest serializers.UserQqRequest
	var user models.User
	err := c.Bind(&getUserQqRequest)
	if err != nil {
		e.CommonError(e.ErrInputInvalid, c)
		return
	}

	// 具体查询
	user.Qq = getUserQqRequest.Qq
	result, err := models.GetUserQq(user)
	if err != nil {
		e.NotFoundError(err, c)
		return
	}

	// 构造返回体
	c.JSON(http.StatusOK, r.SuccessResponse{
		Status: "success",
		Data: serializers.UserResponse{
			User: serializers.SerializeUser(*result),
		},
	})
}
