package example

import (
	q "github.com/athlum/gosp/query"
)

type SchoolModel struct {
	query *q.Query
}

const _schoolTableName = "school"

func (s *SchoolModel) TableName() string {
	return _schoolTableName
}

func (s *School) TableName() string {
	return _schoolTableName
}

type schoolFields struct {
	ID   *q.Field
	Name *q.Field
}

func (fs *schoolFields) Array() []string {
	return []string{
		fs.ID.String(),
		fs.Name.String(),
	}
}

var _schoolFields = &schoolFields{
	ID:   q.TableField(_schoolTableName, "id"),
	Name: q.TableField(_schoolTableName, "name"),
}

func (s *SchoolModel) Fields() *schoolFields {
	return _schoolFields
}

func (s *School) Fields() *schoolFields {
	return _schoolFields
}

func (s *SchoolModel) FieldArray() []string {
	return _schoolFields.Array()
}

func (s *School) FieldArray() []string {
	return _schoolFields.Array()
}

func (s *SchoolModel) PK() (*q.Field, interface{}) {
	m, _ := s.Query().GetMain()
	return m.Model.PK()
}

func (s *School) PK() (*q.Field, interface{}) {
	return s.Fields().ID, s.ID
}

func (s *SchoolModel) Query() *q.Query {
	if s.query == nil {
		s.query = q.NewQuery(&School{})
	}
	return s.query
}

func (s *SchoolModel) Where(cs ...*q.Condition) *SchoolModel {
	s.Query().Where(cs...)
	return s
}

func (s *SchoolModel) getJs(m q.QueryModel, f *q.Field) (*q.Joinset, error) {
	jm, ok := _schoolRels[m.TableName()]
	if ok {
		if f == nil {
			return jm[_schoolDefaultRels[m.TableName()]], nil
		} else {
			if j, ok := jm[f.String()]; ok {
				return j, nil
			}
		}
	}
	return nil, q.ErrRelNotFound
}

var _schoolRels = map[string]q.JoinsetMap{
}

var _schoolDefaultRels = map[string]string{
}

func (s *SchoolModel) Join(qm q.QueryModel, direction string, fs ...*q.Field) *SchoolModel {
	var (
		qs  = s.Query()
		f   *q.Field
		js  *q.Joinset
		err error
	)
	if len(fs) > 0 {
		f = fs[0]
	}
	iqs := qm.Query()
	var main *q.ModelSet
	main, err = iqs.GetMain()
	if err == nil {
		js, err = s.getJs(qm, f)
		if err == nil {
			js.JoinModelSet(main)
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