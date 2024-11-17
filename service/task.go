package service

import (
	"context"
	"sync"
	"time"

	"todolist/pkg/ctl"
	"todolist/pkg/util"
	"todolist/repository/db/dao"
	"todolist/repository/db/model"
	"todolist/types"
)

var TaskSrvOnce sync.Once
var TaskSrvIns *TaskSrv

type TaskSrv struct{}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (tasksrv *TaskSrv) CreateTask(ctx context.Context, req *types.CreateTaskReq, userId uint) (resp interface{}, err error) {
	user, err := dao.NewUserDao(ctx).FindUserByUserId(userId)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}
	task := &model.Task{
		User:      *user,
		Uid:       user.ID,
		Title:     req.Title,
		Status:    req.Status,
		Content:   req.Content,
		StartTime: time.Now().Unix(),
	}
	err = dao.NewTaskDao(ctx).CreateTask(task)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (taskSrv *TaskSrv) ListTask(ctx context.Context, req *types.ListTaskReq, userId uint) (resp interface{}, err error) {
	tasks, total, err := dao.NewTaskDao(ctx).ListTask(req.Page, req.PageSize, userId)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}
	taskRespList := make([]*types.TaskResp, 0)
	for _, task := range tasks {
		taskRespList = append(taskRespList, &types.TaskResp{
			ID:        task.ID,
			View:      task.GetView(),
			Title:     task.Title,
			Status:    task.Status,
			Content:   task.Content,
			CreatedAt: task.CreatedAt.Unix(),
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
		})
	}
	return ctl.RespList(taskRespList, total), nil
}

func (taskSrv *TaskSrv) ShowTask(ctx context.Context, req *types.ShowTaskReq, userId uint) (resp interface{}, err error) {
	task, err := dao.NewTaskDao(ctx).FindTaskByIdAndUserId(req.ID, userId)
	if err != nil {
		util.LogrusObj.Errorln(err)
		return
	}
	respTask := &types.TaskResp{
		ID:        task.ID,
		View:      task.GetView(),
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		CreatedAt: task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
	task.AddView()
	return ctl.RespSuccessWithData(respTask), nil
}

func (taskSrv *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskReq, userId uint) (resp interface{}, err error) {
	task, err := dao.NewTaskDao(ctx).FindTaskByIdAndUserId(req.ID, userId)
	if err != nil {
		util.LogrusObj.Errorln(err)
		return
	}
	task.Title = req.Title
	task.Content = req.Content
	task.Status = req.Status
	dao.NewTaskDao(ctx).Save(&task)
	return ctl.RespSuccess(), nil
}

func (taskSrv *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq, userId uint) (resp interface{}, err error) {
	tasks, err := dao.NewTaskDao(ctx).SearchTask(req.Info, userId)
	if err != nil {
		util.LogrusObj.Errorln(err)
		return
	}
	taskRespList := make([]*types.TaskResp, 0)
	for _, v := range tasks {
		taskRespList = append(taskRespList, &types.TaskResp{
			ID: v.ID,
			Title: v.Title,
			Content: v.Content,
			Status: v.Status,
			View: v.GetView(),
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime: v.EndTime,
		})
	}
	return ctl.RespList(taskRespList, int64(len(taskRespList))), nil
}

func (taskSrv *TaskSrv) DeleteTask(ctx context.Context, req *types.DeleteTaskReq, userId uint) (resp interface{}, err error) {
	err = dao.NewTaskDao(ctx).DeleteTask(req.ID, userId)
	if err != nil {
		util.LogrusObj.Errorln(err)
		return
	}
	
	return ctl.RespSuccess(), nil
} 
