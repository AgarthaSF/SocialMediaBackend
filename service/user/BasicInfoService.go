package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"terminal/models"
	"terminal/response"
)

func GetBasicInfo(c *gin.Context) (*response.UserBasicInfoResp, error) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return nil, err
	}

	user := new(models.User)
	user.ID = id
	if err := models.DB.First(&user).Error; err != nil {
		return nil, err
	}
	resp := response.UserBasicInfoResp{
		ID:       user.ID,
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
	}
	return &resp, nil
}
