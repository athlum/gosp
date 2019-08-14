package example

import (
	q "github.com/athlum/gosp/query"
)

type TeamModel struct {
	query *q.Query
}

const _teamTableName = "team"

func (s *TeamModel) TableName() string {
	return _teamTableName
}

func (s *Team) TableName() string {
	return _teamTableName
}

type teamFields struct {
	ID   *q.Field
	Name *q.Field
}

func (fs *teamFields) Array() []string {
	return []string{
		fs.ID.String(),
		fs.Name.String(),
	}
}

var _teamFields = &teamFields{
	ID:   q.TableField(_teamTableName, "id"),
	Name: q.TableField(_teamTableName, "name"),
}

func (s *TeamModel) Fields() *teamFields {
	return _teamFields
}

func (s *Team) Fields() *teamFields {
	return _teamFields
}

func (s *TeamModel) FieldArray() []string {
	return _teamFields.Array()
}

func (s *Team) FieldArray() []string {
	return _teamFields.Array()
}

func (s *TeamModel) PK() (*q.Field, interface{}) {
	m, _ := s.Queryset().GetMain()
	return m.Model.PK()
}

func (s *Team) PK() (*q.Field, interface{}) {
	return s.Fields().ID, s.ID
}

func (s *TeamModel) Queryset() *q.Query {
	if s.query == nil {
		s.query = q.NewQuery(&Team{})
	}
	return s.query
}

func (s *TeamModel) Where(cs ...*q.Condition) *TeamModel {
	s.Queryset().Where(cs...)
	return s
}

func (s *TeamModel) getJs(m q.QueryModel, f *q.Field) (*q.Joinset, error) {
	jm, ok := _teamRels[m.TableName()]
	if ok {
		if f == nil {
			return jm[_teamDefaultRels[m.TableName()]], nil
		} else {
			if j, ok := jm[f.String()]; ok {
				return j, nil
			}
		}
	}
	return nil, q.ErrRelNotFound
}

var _teamRels = map[string]q.JoinsetMap{
}

var _teamDefaultRels = map[string]string{
}

func (s *TeamModel) Join(qm q.QueryModel, direction string, fs ...*q.Field) *TeamModel {
	var (
		qs  = s.Queryset()
		f   *q.Field
		js  *q.Joinset
		err error
	)
	if len(fs) > 0 {
		f = fs[0]
	}
	iqs := qm.Queryset()
	var main *q.Modelset
	main, err = iqs.GetMain()
	if err == nil {
		js, err = s.getJs(qm, f)
		if err == nil {
			js.JoinModelset(main)
		}
		if direction != "" {
			js.Direction(direction)
		}
		qs.Join(js)
	} else {
		qs.SetErr(err)
	}
	return s
}