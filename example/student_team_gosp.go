package example

import (
	q "github.com/athlum/gosp/query"
)

type StudentTeam struct {
	ID        int64 `db:"id, pk, autoincrement"`
	StudentID int64 `db:"student_id"`
	TeamID    int64 `db:"team_id"`
}

type StudentTeamModel struct {
	query *q.Query
}

const _studentTeamTableName = "student_team"

func (s *StudentTeamModel) TableName() string {
	return _studentTeamTableName
}

func (s *StudentTeam) TableName() string {
	return _studentTeamTableName
}

type studentTeamFields struct {
	ID        *q.Field
	StudentID *q.Field
	TeamID    *q.Field
}

func (fs *studentTeamFields) Array() []string {
	return []string{
		fs.ID.String(),
		fs.StudentID.String(),
		fs.TeamID.String(),
	}
}

var _studentTeamFields = &studentTeamFields{
	ID:        q.TableField(_studentTeamTableName, "id"),
	StudentID: q.TableField(_studentTeamTableName, "student_id"),
	TeamID:    q.TableField(_studentTeamTableName, "team_id"),
}

func (s *StudentTeamModel) Fields() *studentTeamFields {
	return _studentTeamFields
}

func (s *StudentTeam) Fields() *studentTeamFields {
	return _studentTeamFields
}

func (s *StudentTeamModel) FieldArray() []string {
	return _studentTeamFields.Array()
}

func (s *StudentTeam) FieldArray() []string {
	return _studentTeamFields.Array()
}

func (s *StudentTeamModel) PK() (*q.Field, interface{}) {
	m, _ := s.Queryset().GetMain()
	return m.Model.PK()
}

func (s *StudentTeam) PK() (*q.Field, interface{}) {
	return s.Fields().ID, s.ID
}

func (s *StudentTeamModel) Queryset() *q.Query {
	if s.query == nil {
		s.query = q.NewQuery(&StudentTeam{})
	}
	return s.query
}