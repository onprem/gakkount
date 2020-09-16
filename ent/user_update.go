// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/prmsrswt/edu-accounts/ent/course"
	"github.com/prmsrswt/edu-accounts/ent/department"
	"github.com/prmsrswt/edu-accounts/ent/predicate"
	"github.com/prmsrswt/edu-accounts/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks      []Hook
	mutation   *UserMutation
	predicates []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmail sets the email field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetHash sets the hash field.
func (uu *UserUpdate) SetHash(s string) *UserUpdate {
	uu.mutation.SetHash(s)
	return uu
}

// SetRole sets the role field.
func (uu *UserUpdate) SetRole(u user.Role) *UserUpdate {
	uu.mutation.SetRole(u)
	return uu
}

// SetNillableRole sets the role field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(u *user.Role) *UserUpdate {
	if u != nil {
		uu.SetRole(*u)
	}
	return uu
}

// SetPhoto sets the photo field.
func (uu *UserUpdate) SetPhoto(s string) *UserUpdate {
	uu.mutation.SetPhoto(s)
	return uu
}

// SetNillablePhoto sets the photo field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhoto(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhoto(*s)
	}
	return uu
}

// ClearPhoto clears the value of photo.
func (uu *UserUpdate) ClearPhoto() *UserUpdate {
	uu.mutation.ClearPhoto()
	return uu
}

// SetAltEmail sets the altEmail field.
func (uu *UserUpdate) SetAltEmail(s string) *UserUpdate {
	uu.mutation.SetAltEmail(s)
	return uu
}

// SetNillableAltEmail sets the altEmail field if the given value is not nil.
func (uu *UserUpdate) SetNillableAltEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetAltEmail(*s)
	}
	return uu
}

// ClearAltEmail clears the value of altEmail.
func (uu *UserUpdate) ClearAltEmail() *UserUpdate {
	uu.mutation.ClearAltEmail()
	return uu
}

// SetPhone sets the phone field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// ClearPhone clears the value of phone.
func (uu *UserUpdate) ClearPhone() *UserUpdate {
	uu.mutation.ClearPhone()
	return uu
}

// SetSalutation sets the salutation field.
func (uu *UserUpdate) SetSalutation(s string) *UserUpdate {
	uu.mutation.SetSalutation(s)
	return uu
}

// SetNillableSalutation sets the salutation field if the given value is not nil.
func (uu *UserUpdate) SetNillableSalutation(s *string) *UserUpdate {
	if s != nil {
		uu.SetSalutation(*s)
	}
	return uu
}

// ClearSalutation clears the value of salutation.
func (uu *UserUpdate) ClearSalutation() *UserUpdate {
	uu.mutation.ClearSalutation()
	return uu
}

// SetLinkedin sets the linkedin field.
func (uu *UserUpdate) SetLinkedin(s string) *UserUpdate {
	uu.mutation.SetLinkedin(s)
	return uu
}

// SetNillableLinkedin sets the linkedin field if the given value is not nil.
func (uu *UserUpdate) SetNillableLinkedin(s *string) *UserUpdate {
	if s != nil {
		uu.SetLinkedin(*s)
	}
	return uu
}

// ClearLinkedin clears the value of linkedin.
func (uu *UserUpdate) ClearLinkedin() *UserUpdate {
	uu.mutation.ClearLinkedin()
	return uu
}

// SetTwitter sets the twitter field.
func (uu *UserUpdate) SetTwitter(s string) *UserUpdate {
	uu.mutation.SetTwitter(s)
	return uu
}

// SetNillableTwitter sets the twitter field if the given value is not nil.
func (uu *UserUpdate) SetNillableTwitter(s *string) *UserUpdate {
	if s != nil {
		uu.SetTwitter(*s)
	}
	return uu
}

// ClearTwitter clears the value of twitter.
func (uu *UserUpdate) ClearTwitter() *UserUpdate {
	uu.mutation.ClearTwitter()
	return uu
}

// SetFacebook sets the facebook field.
func (uu *UserUpdate) SetFacebook(s string) *UserUpdate {
	uu.mutation.SetFacebook(s)
	return uu
}

// SetNillableFacebook sets the facebook field if the given value is not nil.
func (uu *UserUpdate) SetNillableFacebook(s *string) *UserUpdate {
	if s != nil {
		uu.SetFacebook(*s)
	}
	return uu
}

// ClearFacebook clears the value of facebook.
func (uu *UserUpdate) ClearFacebook() *UserUpdate {
	uu.mutation.ClearFacebook()
	return uu
}

// SetGithub sets the github field.
func (uu *UserUpdate) SetGithub(s string) *UserUpdate {
	uu.mutation.SetGithub(s)
	return uu
}

// SetNillableGithub sets the github field if the given value is not nil.
func (uu *UserUpdate) SetNillableGithub(s *string) *UserUpdate {
	if s != nil {
		uu.SetGithub(*s)
	}
	return uu
}

// ClearGithub clears the value of github.
func (uu *UserUpdate) ClearGithub() *UserUpdate {
	uu.mutation.ClearGithub()
	return uu
}

// SetRollNo sets the rollNo field.
func (uu *UserUpdate) SetRollNo(s string) *UserUpdate {
	uu.mutation.SetRollNo(s)
	return uu
}

// SetNillableRollNo sets the rollNo field if the given value is not nil.
func (uu *UserUpdate) SetNillableRollNo(s *string) *UserUpdate {
	if s != nil {
		uu.SetRollNo(*s)
	}
	return uu
}

// ClearRollNo clears the value of rollNo.
func (uu *UserUpdate) ClearRollNo() *UserUpdate {
	uu.mutation.ClearRollNo()
	return uu
}

// SetAdmissionTime sets the admissionTime field.
func (uu *UserUpdate) SetAdmissionTime(t time.Time) *UserUpdate {
	uu.mutation.SetAdmissionTime(t)
	return uu
}

// SetNillableAdmissionTime sets the admissionTime field if the given value is not nil.
func (uu *UserUpdate) SetNillableAdmissionTime(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetAdmissionTime(*t)
	}
	return uu
}

// ClearAdmissionTime clears the value of admissionTime.
func (uu *UserUpdate) ClearAdmissionTime() *UserUpdate {
	uu.mutation.ClearAdmissionTime()
	return uu
}

// SetCourseEndTime sets the courseEndTime field.
func (uu *UserUpdate) SetCourseEndTime(t time.Time) *UserUpdate {
	uu.mutation.SetCourseEndTime(t)
	return uu
}

// SetNillableCourseEndTime sets the courseEndTime field if the given value is not nil.
func (uu *UserUpdate) SetNillableCourseEndTime(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCourseEndTime(*t)
	}
	return uu
}

// ClearCourseEndTime clears the value of courseEndTime.
func (uu *UserUpdate) ClearCourseEndTime() *UserUpdate {
	uu.mutation.ClearCourseEndTime()
	return uu
}

// SetDesignation sets the Designation field.
func (uu *UserUpdate) SetDesignation(s string) *UserUpdate {
	uu.mutation.SetDesignation(s)
	return uu
}

// SetNillableDesignation sets the Designation field if the given value is not nil.
func (uu *UserUpdate) SetNillableDesignation(s *string) *UserUpdate {
	if s != nil {
		uu.SetDesignation(*s)
	}
	return uu
}

// ClearDesignation clears the value of Designation.
func (uu *UserUpdate) ClearDesignation() *UserUpdate {
	uu.mutation.ClearDesignation()
	return uu
}

// SetCourseID sets the course edge to Course by id.
func (uu *UserUpdate) SetCourseID(id int) *UserUpdate {
	uu.mutation.SetCourseID(id)
	return uu
}

// SetNillableCourseID sets the course edge to Course by id if the given value is not nil.
func (uu *UserUpdate) SetNillableCourseID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetCourseID(*id)
	}
	return uu
}

// SetCourse sets the course edge to Course.
func (uu *UserUpdate) SetCourse(c *Course) *UserUpdate {
	return uu.SetCourseID(c.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (uu *UserUpdate) SetDepartmentID(id int) *UserUpdate {
	uu.mutation.SetDepartmentID(id)
	return uu
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (uu *UserUpdate) SetNillableDepartmentID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetDepartmentID(*id)
	}
	return uu
}

// SetDepartment sets the department edge to Department.
func (uu *UserUpdate) SetDepartment(d *Department) *UserUpdate {
	return uu.SetDepartmentID(d.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCourse clears the course edge to Course.
func (uu *UserUpdate) ClearCourse() *UserUpdate {
	uu.mutation.ClearCourse()
	return uu
}

// ClearDepartment clears the department edge to Department.
func (uu *UserUpdate) ClearDepartment() *UserUpdate {
	uu.mutation.ClearDepartment()
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := uu.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return 0, &ValidationError{Name: "role", err: fmt.Errorf("ent: validator failed for field \"role\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldHash,
		})
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldRole,
		})
	}
	if value, ok := uu.mutation.Photo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhoto,
		})
	}
	if uu.mutation.PhotoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPhoto,
		})
	}
	if value, ok := uu.mutation.AltEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAltEmail,
		})
	}
	if uu.mutation.AltEmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAltEmail,
		})
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if uu.mutation.PhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uu.mutation.Salutation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSalutation,
		})
	}
	if uu.mutation.SalutationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldSalutation,
		})
	}
	if value, ok := uu.mutation.Linkedin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLinkedin,
		})
	}
	if uu.mutation.LinkedinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldLinkedin,
		})
	}
	if value, ok := uu.mutation.Twitter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTwitter,
		})
	}
	if uu.mutation.TwitterCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTwitter,
		})
	}
	if value, ok := uu.mutation.Facebook(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFacebook,
		})
	}
	if uu.mutation.FacebookCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldFacebook,
		})
	}
	if value, ok := uu.mutation.Github(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGithub,
		})
	}
	if uu.mutation.GithubCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldGithub,
		})
	}
	if value, ok := uu.mutation.RollNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRollNo,
		})
	}
	if uu.mutation.RollNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldRollNo,
		})
	}
	if value, ok := uu.mutation.AdmissionTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldAdmissionTime,
		})
	}
	if uu.mutation.AdmissionTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldAdmissionTime,
		})
	}
	if value, ok := uu.mutation.CourseEndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCourseEndTime,
		})
	}
	if uu.mutation.CourseEndTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldCourseEndTime,
		})
	}
	if value, ok := uu.mutation.Designation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDesignation,
		})
	}
	if uu.mutation.DesignationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldDesignation,
		})
	}
	if uu.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CourseTable,
			Columns: []string{user.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CourseTable,
			Columns: []string{user.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.DepartmentTable,
			Columns: []string{user.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.DepartmentTable,
			Columns: []string{user.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmail sets the email field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetHash sets the hash field.
func (uuo *UserUpdateOne) SetHash(s string) *UserUpdateOne {
	uuo.mutation.SetHash(s)
	return uuo
}

// SetRole sets the role field.
func (uuo *UserUpdateOne) SetRole(u user.Role) *UserUpdateOne {
	uuo.mutation.SetRole(u)
	return uuo
}

// SetNillableRole sets the role field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(u *user.Role) *UserUpdateOne {
	if u != nil {
		uuo.SetRole(*u)
	}
	return uuo
}

// SetPhoto sets the photo field.
func (uuo *UserUpdateOne) SetPhoto(s string) *UserUpdateOne {
	uuo.mutation.SetPhoto(s)
	return uuo
}

// SetNillablePhoto sets the photo field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhoto(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhoto(*s)
	}
	return uuo
}

// ClearPhoto clears the value of photo.
func (uuo *UserUpdateOne) ClearPhoto() *UserUpdateOne {
	uuo.mutation.ClearPhoto()
	return uuo
}

// SetAltEmail sets the altEmail field.
func (uuo *UserUpdateOne) SetAltEmail(s string) *UserUpdateOne {
	uuo.mutation.SetAltEmail(s)
	return uuo
}

// SetNillableAltEmail sets the altEmail field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAltEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAltEmail(*s)
	}
	return uuo
}

// ClearAltEmail clears the value of altEmail.
func (uuo *UserUpdateOne) ClearAltEmail() *UserUpdateOne {
	uuo.mutation.ClearAltEmail()
	return uuo
}

// SetPhone sets the phone field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// ClearPhone clears the value of phone.
func (uuo *UserUpdateOne) ClearPhone() *UserUpdateOne {
	uuo.mutation.ClearPhone()
	return uuo
}

// SetSalutation sets the salutation field.
func (uuo *UserUpdateOne) SetSalutation(s string) *UserUpdateOne {
	uuo.mutation.SetSalutation(s)
	return uuo
}

// SetNillableSalutation sets the salutation field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSalutation(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSalutation(*s)
	}
	return uuo
}

// ClearSalutation clears the value of salutation.
func (uuo *UserUpdateOne) ClearSalutation() *UserUpdateOne {
	uuo.mutation.ClearSalutation()
	return uuo
}

// SetLinkedin sets the linkedin field.
func (uuo *UserUpdateOne) SetLinkedin(s string) *UserUpdateOne {
	uuo.mutation.SetLinkedin(s)
	return uuo
}

// SetNillableLinkedin sets the linkedin field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLinkedin(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLinkedin(*s)
	}
	return uuo
}

// ClearLinkedin clears the value of linkedin.
func (uuo *UserUpdateOne) ClearLinkedin() *UserUpdateOne {
	uuo.mutation.ClearLinkedin()
	return uuo
}

// SetTwitter sets the twitter field.
func (uuo *UserUpdateOne) SetTwitter(s string) *UserUpdateOne {
	uuo.mutation.SetTwitter(s)
	return uuo
}

// SetNillableTwitter sets the twitter field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTwitter(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTwitter(*s)
	}
	return uuo
}

// ClearTwitter clears the value of twitter.
func (uuo *UserUpdateOne) ClearTwitter() *UserUpdateOne {
	uuo.mutation.ClearTwitter()
	return uuo
}

// SetFacebook sets the facebook field.
func (uuo *UserUpdateOne) SetFacebook(s string) *UserUpdateOne {
	uuo.mutation.SetFacebook(s)
	return uuo
}

// SetNillableFacebook sets the facebook field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFacebook(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFacebook(*s)
	}
	return uuo
}

// ClearFacebook clears the value of facebook.
func (uuo *UserUpdateOne) ClearFacebook() *UserUpdateOne {
	uuo.mutation.ClearFacebook()
	return uuo
}

// SetGithub sets the github field.
func (uuo *UserUpdateOne) SetGithub(s string) *UserUpdateOne {
	uuo.mutation.SetGithub(s)
	return uuo
}

// SetNillableGithub sets the github field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGithub(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetGithub(*s)
	}
	return uuo
}

// ClearGithub clears the value of github.
func (uuo *UserUpdateOne) ClearGithub() *UserUpdateOne {
	uuo.mutation.ClearGithub()
	return uuo
}

// SetRollNo sets the rollNo field.
func (uuo *UserUpdateOne) SetRollNo(s string) *UserUpdateOne {
	uuo.mutation.SetRollNo(s)
	return uuo
}

// SetNillableRollNo sets the rollNo field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRollNo(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRollNo(*s)
	}
	return uuo
}

// ClearRollNo clears the value of rollNo.
func (uuo *UserUpdateOne) ClearRollNo() *UserUpdateOne {
	uuo.mutation.ClearRollNo()
	return uuo
}

// SetAdmissionTime sets the admissionTime field.
func (uuo *UserUpdateOne) SetAdmissionTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetAdmissionTime(t)
	return uuo
}

// SetNillableAdmissionTime sets the admissionTime field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAdmissionTime(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetAdmissionTime(*t)
	}
	return uuo
}

// ClearAdmissionTime clears the value of admissionTime.
func (uuo *UserUpdateOne) ClearAdmissionTime() *UserUpdateOne {
	uuo.mutation.ClearAdmissionTime()
	return uuo
}

// SetCourseEndTime sets the courseEndTime field.
func (uuo *UserUpdateOne) SetCourseEndTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCourseEndTime(t)
	return uuo
}

// SetNillableCourseEndTime sets the courseEndTime field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCourseEndTime(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCourseEndTime(*t)
	}
	return uuo
}

// ClearCourseEndTime clears the value of courseEndTime.
func (uuo *UserUpdateOne) ClearCourseEndTime() *UserUpdateOne {
	uuo.mutation.ClearCourseEndTime()
	return uuo
}

// SetDesignation sets the Designation field.
func (uuo *UserUpdateOne) SetDesignation(s string) *UserUpdateOne {
	uuo.mutation.SetDesignation(s)
	return uuo
}

// SetNillableDesignation sets the Designation field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDesignation(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDesignation(*s)
	}
	return uuo
}

// ClearDesignation clears the value of Designation.
func (uuo *UserUpdateOne) ClearDesignation() *UserUpdateOne {
	uuo.mutation.ClearDesignation()
	return uuo
}

// SetCourseID sets the course edge to Course by id.
func (uuo *UserUpdateOne) SetCourseID(id int) *UserUpdateOne {
	uuo.mutation.SetCourseID(id)
	return uuo
}

// SetNillableCourseID sets the course edge to Course by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCourseID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetCourseID(*id)
	}
	return uuo
}

// SetCourse sets the course edge to Course.
func (uuo *UserUpdateOne) SetCourse(c *Course) *UserUpdateOne {
	return uuo.SetCourseID(c.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (uuo *UserUpdateOne) SetDepartmentID(id int) *UserUpdateOne {
	uuo.mutation.SetDepartmentID(id)
	return uuo
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDepartmentID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetDepartmentID(*id)
	}
	return uuo
}

// SetDepartment sets the department edge to Department.
func (uuo *UserUpdateOne) SetDepartment(d *Department) *UserUpdateOne {
	return uuo.SetDepartmentID(d.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCourse clears the course edge to Course.
func (uuo *UserUpdateOne) ClearCourse() *UserUpdateOne {
	uuo.mutation.ClearCourse()
	return uuo
}

// ClearDepartment clears the department edge to Department.
func (uuo *UserUpdateOne) ClearDepartment() *UserUpdateOne {
	uuo.mutation.ClearDepartment()
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if v, ok := uuo.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return nil, &ValidationError{Name: "role", err: fmt.Errorf("ent: validator failed for field \"role\": %w", err)}
		}
	}

	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldHash,
		})
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldRole,
		})
	}
	if value, ok := uuo.mutation.Photo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhoto,
		})
	}
	if uuo.mutation.PhotoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPhoto,
		})
	}
	if value, ok := uuo.mutation.AltEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAltEmail,
		})
	}
	if uuo.mutation.AltEmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAltEmail,
		})
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if uuo.mutation.PhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uuo.mutation.Salutation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSalutation,
		})
	}
	if uuo.mutation.SalutationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldSalutation,
		})
	}
	if value, ok := uuo.mutation.Linkedin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLinkedin,
		})
	}
	if uuo.mutation.LinkedinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldLinkedin,
		})
	}
	if value, ok := uuo.mutation.Twitter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTwitter,
		})
	}
	if uuo.mutation.TwitterCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTwitter,
		})
	}
	if value, ok := uuo.mutation.Facebook(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFacebook,
		})
	}
	if uuo.mutation.FacebookCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldFacebook,
		})
	}
	if value, ok := uuo.mutation.Github(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGithub,
		})
	}
	if uuo.mutation.GithubCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldGithub,
		})
	}
	if value, ok := uuo.mutation.RollNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRollNo,
		})
	}
	if uuo.mutation.RollNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldRollNo,
		})
	}
	if value, ok := uuo.mutation.AdmissionTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldAdmissionTime,
		})
	}
	if uuo.mutation.AdmissionTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldAdmissionTime,
		})
	}
	if value, ok := uuo.mutation.CourseEndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCourseEndTime,
		})
	}
	if uuo.mutation.CourseEndTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldCourseEndTime,
		})
	}
	if value, ok := uuo.mutation.Designation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDesignation,
		})
	}
	if uuo.mutation.DesignationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldDesignation,
		})
	}
	if uuo.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CourseTable,
			Columns: []string{user.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CourseTable,
			Columns: []string{user.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.DepartmentTable,
			Columns: []string{user.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.DepartmentTable,
			Columns: []string{user.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
