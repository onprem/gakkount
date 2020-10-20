package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prmsrswt/edu-accounts/ent"
	"github.com/prmsrswt/edu-accounts/ent/department"
)

func (a *API) queryAllDepartments(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	query := a.store.Department.Query().
		Order(ent.Asc(department.FieldName)).
		Offset(page * limit)

	if limit != 0 {
		query = query.Limit(limit)
	}

	departments, err := query.All(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: queryall departments: %w", err), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "departments": departments})
}

func (a *API) handleNewDepartment(c *gin.Context) {
	var i struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&i); err != nil {
		respondError(http.StatusBadRequest, err.Error(), c)
		return
	}

	department, err := a.createDepartment(context.TODO(), i.Name)
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: create department: %w", err), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "New department created", "department": department})
}

func (a *API) handleUpdateDepartment(c *gin.Context) {
	var i struct {
		ID   int    `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&i); err != nil {
		respondError(http.StatusBadRequest, err.Error(), c)
		return
	}

	department, err := a.store.Department.UpdateOneID(i.ID).SetName(i.Name).Save(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: update department: %w", err), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Department updated", "department": department})
}
