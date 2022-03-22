package user

import (
	"github.com/gin-gonic/gin"
	"mingin/handler"
	"mingin/pkg/errno"
	"mingin/service"
)

func List(c *gin.Context) {
	var r ListReq
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
	}
	handler.SendResponse(c, nil, ListRep{
		TotalCount: count,
		UserList:   infos,
	})
}
