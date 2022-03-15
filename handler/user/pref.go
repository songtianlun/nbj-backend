package user

import (
	"github.com/gin-gonic/gin"
	"minepin-backend/handler"
	"minepin-backend/model"
	"minepin-backend/pkg/errno"
	"minepin-backend/utils"
	"strconv"
)

func GetPreferences(c *gin.Context) {
	uid, err := utils.GetUint64ByContext(c, "id")
	if err != nil {
		handler.SendResponse(c, errno.ErrUrl, nil)
	}
	up, err := model.GetPrefByUID(uid)
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, up)
}

func SetPreferences(c *gin.Context) {
	var up model.UserPrefModel
	uid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.SendResponse(c, errno.ErrUrl, nil)
	}
	if err := c.Bind(&up); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := up.UpdateUserPref(uid); err != nil {
		handler.SendResponse(c, errno.ErrUpdatePref, nil)
		return
	}
	handler.SendResponse(c, nil, up)

}
