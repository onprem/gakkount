package api

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prmsrswt/edu-accounts/ent"
)

func (a *API) createStudent(
	ctx context.Context,
	email, name, rollNo, password string,
	admissionTime, courseEndTime time.Time,
	course *ent.Course,
) (*ent.User, error) {
	email = strings.ToLower(email)
	rollNo = strings.ToLower(rollNo)

	if !strings.HasSuffix(email, "@"+a.domain) {
		return nil, errors.New("invalid primary email")
	}
	if admissionTime.After(courseEndTime) {
		return nil, errors.New("admission time must be before the course end time")
	}
	hash, err := generateHash(password)
	if err != nil {
		return nil, err
	}
	return a.store.User.
		Create().
		SetEmail(email).
		SetName(name).
		SetHash(hash).
		SetRole("student").
		SetRollNo(rollNo).
		SetAdmissionTime(admissionTime).
		SetCourseEndTime(courseEndTime).
		SetCourse(course).
		Save(ctx)
}

func (a *API) createFaculty(
	ctx context.Context,
	email, name, salutation, password string,
	department *ent.Department,
) (*ent.User, error) {
	email = strings.ToLower(email)

	if !strings.HasSuffix(email, "@"+a.domain) {
		return nil, errors.New("invalid primary email")
	}
	hash, err := generateHash(password)
	if err != nil {
		return nil, err
	}
	return a.store.User.
		Create().
		SetEmail(email).
		SetName(name).
		SetHash(hash).
		SetRole("faculty").
		SetSalutation(salutation).
		SetDepartment(department).
		Save(ctx)
}

func (a *API) createAdmin(ctx context.Context, email, name, password string) (*ent.User, error) {
	email = strings.ToLower(email)

	if !strings.HasSuffix(email, "@"+a.domain) {
		return nil, errors.New("invalid primary email")
	}
	hash, err := generateHash(password)
	if err != nil {
		return nil, err
	}
	return a.store.User.
		Create().
		SetEmail(email).
		SetName(name).
		SetHash(hash).
		SetRole("admin").
		Save(ctx)
}

func (a *API) createCourse(ctx context.Context, name, code string, semesters int) (*ent.Course, error) {
	return a.store.Course.
		Create().
		SetName(name).
		SetCode(code).
		SetSemesters(semesters).
		Save(ctx)
}

func (a *API) createDepartment(ctx context.Context, name string) (*ent.Department, error) {
	return a.store.Department.
		Create().
		SetName(name).
		Save(ctx)
}

func (a *API) createDummy(c *gin.Context) {
	ctx := context.Background()
	imt, _ := a.createCourse(ctx, "IPG MTech", "IMT", 10)
	img, _ := a.createCourse(ctx, "IPG MBA", "IMG", 10)
	bcs, _ := a.createCourse(ctx, "BTech in Computer Science", "BCS", 8)

	startT, _ := time.Parse(time.RFC3339, "2017-08-13T00:00:00Z")
	oldT := startT.AddDate(-3, 0, 0)
	a.createStudent(ctx, "imt68@iiitm.ac.in", "Prem Kumar", "2017imt-068", "1234", startT, startT.AddDate(5, 0, 0), imt)
	a.createStudent(ctx, "imt73@iiitm.ac.in", "Rohit Kunji", "2017imt-073", "1234", startT, startT.AddDate(5, 0, 0), imt)
	a.createStudent(ctx, "img31@iiitm.ac.in", "Manish Mavi", "2017img-031", "1234", startT, startT.AddDate(5, 0, 0), img)
	a.createStudent(ctx, "bcs08@iiitm.ac.in", "Divyanshu Tripathi", "2017bcs-08", "1234", startT, startT.AddDate(4, 0, 0), bcs)
	a.createStudent(ctx, "imt99@iiitm.ac.in", "Shubham Shukla", "2014imt-068", "1234", oldT, oldT.AddDate(5, 0, 0), imt)

	a.createAdmin(ctx, "admin@iiitm.ac.in", "Admin", "1234")

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "created dummy accounts"})
}
