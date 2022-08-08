package v1

type CountJobRequest = ListJobRequest

type ListJobRequest struct {
	Name   string `form:"name" json:"name" binding:"max=100"`
	Status int    `form:"status" json:"status" binding:"oneof= 0 1 2"`
	Group  int    `form:"group" json:"group" binding:"oneof=0 1 2"`
}

type CreateJobRequest struct {
	Name       string `form:"name" json:"name" binding:"required,max=100"`
	Status     int    `form:"status" json:"status" binding:"oneof=-1 0 1"`
	Group      int    `form:"group" json:"group" binding:"oneof= 0 1"`
	Type       int    `form:"type" json:"type" binding:"oneof= 0 1"`
	Strategy   int    `form:"strategy" json:"strategy" binding:"oneof= 0 1"`
	Concurrent int    `form:"concurrent" json:"concurrent" binding:"oneof= 0 1"`
	Cron       string `form:"cron" json:"cron" binding:"required,max=100"`
	Target     string `form:"target" json:"target" binding:"required,max=100"`
	Args       string `form:"args" json:"args" `
}

type UpdateJobRequest struct {
	Id         int    `form:"id" json:"id" binding:"required"`
	Name       string `form:"name" json:"name" binding:"max=100"`
	Status     int    `form:"status" json:"status" binding:"oneof=-1 0 1"`
	Group      int    `form:"group" json:"group" binding:"oneof= 0 1"`
	Type       int    `form:"type" json:"type" binding:"oneof= 0 1"`
	Strategy   int    `form:"strategy" json:"strategy" binding:"oneof= 0 1 2 3"`
	Concurrent int    `form:"concurrent" json:"concurrent" binding:"oneof= 0 1"`
	Cron       string `form:"cron" json:"cron" binding:"max=100"`
	Target     string `form:"target" json:"target" binding:"max=100"`
	Args       string `form:"args" json:"args" binding:"max=100"`
}

type DeleteJobRequest struct {
	Id int `form:"id" json:"id" binding:"required"`
}
type CommonJob struct {
	Id int `form:"id" json:"id" binding:"required"`
}
