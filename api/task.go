package api

import (
	"net/http"
	"todolist/consts"
	"todolist/pkg/ctl"
	"todolist/pkg/util"
	"todolist/service"
	"todolist/types"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateTaskReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		} else {
			userId, err := ctl.GetUserId(ctx)
			if err != nil {
				util.LogrusObj.Infoln(err)
				return
			}
			resp, err := service.GetTaskSrv().CreateTask(ctx, &req, userId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func ListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListTaskReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		} else {
			userId, err := ctl.GetUserId(ctx)
			if err != nil {
				util.LogrusObj.Infoln(err)
				return
			}
			// 参数校验
			if req.PageSize == 0 {
				req.PageSize = consts.BasePageSize
			}
			resp, err := service.GetTaskSrv().ListTask(ctx, &req, userId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp) 
		}
	}
}

func ShowTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ShowTaskReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		} else {
			if req.Id == 0 {
				ctx.JSON(http.StatusNotFound, gin.H{
					"Status": 400,
					"msg": "缺少任务ID",
				})
				return
			}
			userId, err := ctl.GetUserId(ctx)
			if err != nil {
				util.LogrusObj.Infoln(err)
				return
			}
			resp, err := service.GetTaskSrv().ShowTask(ctx, &req, userId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		}
	}
}