package course

import (
	"github.com/MuxiKeStack/muxiK-StackBackend/handler"
	"github.com/MuxiKeStack/muxiK-StackBackend/model"
	//	 "github.com/MuxiKeStack/muxiK-StackBackend/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// 删除课程
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	//	id, err := c.Param("id")
	//	if err != nil {
	//		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
	//		return
	//	}

	course := &model.UsingCourseModel{CourseId: id}
	if err := course.GetById(); err != nil {
		log.Info("course.GetById() error.")
		handler.SendError(c, err, nil, err.Error())
		return
	}

	if err := course.Delete(); err != nil {
		log.Info("course.Delete() error.")
		handler.SendError(c, err, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, nil)
}
