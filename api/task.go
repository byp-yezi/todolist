package api

import (
	"net/http"
	"todolist/consts"
	"todolist/pkg/ctl"
	"todolist/pkg/e"
	"todolist/pkg/util"
	"todolist/service"
	"todolist/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
			// if req.Id == 0 {
			// 	ctx.JSON(http.StatusNotFound, gin.H{
			// 		"Status": 400,
			// 		"msg": "缺少任务ID",
			// 	})
			// 	return
			// }
			userId, err := ctl.GetUserId(ctx)
			if err != nil {
				util.LogrusObj.Infoln(err)
				return
			}
			resp, err := service.GetTaskSrv().ShowTask(ctx, &req, userId)
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status": e.ErrorTaskNotFound,
					"msg": e.GetMsg(e.ErrorTaskNotFound),
				})
				return
			}
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UpdateTaskReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Errorln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		} else {
			if req.Status != 0 && req.Status != 1 {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status": e.ErrorTaskStatusParameter,
					"msg": e.GetMsg(e.ErrorTaskStatusParameter),
				})
				return
			}
			userId, err := ctl.GetUserId(ctx)
			if err != nil {
				util.LogrusObj.Errorln(err)
				return
			}
			resp, err := service.GetTaskSrv().UpdateTask(ctx, &req, userId)
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status": e.ErrorTaskNotFoundByUser,
					"msg": e.GetMsg(e.ErrorTaskNotFoundByUser),
				})
				return
			}
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Errorln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		} else {
			userId, err := ctl.GetUserId(ctx)
			if err != nil {
				util.LogrusObj.Errorln(err)
				return
			}
			resp, err := service.GetTaskSrv().SearchTask(ctx, &req, userId)
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status": e.ErrorTaskNotFound,
					"msg": e.GetMsg(e.ErrorTaskNotFound),
				})
				return
			}
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func DeleteTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteTaskReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Errorln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		userId, err := ctl.GetUserId(ctx)
		if err != nil {
			util.LogrusObj.Errorln(err)
			return
		}
		resp, err := service.GetTaskSrv().DeleteTask(ctx, &req, userId)
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status": e.ErrorTaskNotFound,
				"msg": e.GetMsg(e.ErrorTaskNotFound),
			})
			return
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
