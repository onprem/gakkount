package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prmsrswt/edu-accounts/ent"
	"github.com/prmsrswt/edu-accounts/ent/course"
)

func (a *API) queryAllCourses(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	query := a.store.Course.Query().
		Order(ent.Asc(course.FieldCode)).
		Offset(page * limit)

	if limit != 0 {
		query = query.Limit(limit)
	}

	courses, err := query.All(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: queryall courses: %w", err), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "courses": courses})
}

func (a *API) handleNewCourse(c *gin.Context) {
	var i struct {
		Name      string `json:"name" binding:"required"`
		Code      string `json:"code" binding:"required"`
		Semesters int    `json:"semesters" binding:"required"`
	}

	if err := c.ShouldBindJSON(&i); err != nil {
		respondError(http.StatusBadRequest, err.Error(), c)
		return
	}

	course, err := a.createCourse(context.TODO(), i.Name, i.Code, i.Semesters)
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: create course: %w", err), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "New course created", "course": course})
}
