package types

type CreateTaskReq struct {
	Title   string `json:"title" form:"title" binding:"required,min=2,max=100"`
	Status  int    `json:"status" form:"status"` // 0代办 1已完成
	Content string `json:"content" form:"content" binding:"max=1000"`
}

type ListTaskReq struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" from:"page_size"`
}

type TaskResp struct {
	ID        uint   `json:"id"`
	View      uint64 `json:"view"`
	Title     string `json:"title"`
	Status    int    `json:"status"` // 0代办 1已完成
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type ShowTaskReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

type UpdateTaskReq struct {
	ID      uint   `json:"id" form:"id" binding:"required"`
	Title   string `json:"title" form:"title" binding:"required,min=2,max=100"`
	Status  int    `json:"status" form:"status"` // 0代办 1已完成
	Content string `json:"content" form:"content" binding:"max=1000"`
}

type SearchTaskReq struct {
	Info string `json:"info" form:"info"`
}

type DeleteTaskReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
