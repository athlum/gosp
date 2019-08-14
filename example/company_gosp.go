package example

import (
	q "github.com/athlum/gosp/query"
)

type CompanyModel struct {
	query *q.Query
}

const _companyTableName = "company"

func (s *CompanyModel) TableName() string {
	return _companyTableName
}

func (s *Company) TableName() string {
	return _companyTableName
}

func (s *CompanyModel) Type() int {
	return q.QueryModel_Model
}

type companyFields struct {
	ID   *q.Field
	Name *q.Field
}

func (fs *companyFields) Array() []string {
	return []string{
		fs.ID.String(),
		fs.Name.String(),
	}
}

var _companyFields = &companyFields{
	ID:   q.TableField(_companyTableName, "id"),
	Name: q.TableField(_companyTableName, "name"),
}

func (s *CompanyModel) Fields() *companyFields {
	return _companyFields
}

func (s *Company) Fields() *companyFields {
	return _companyFields
}

func (s *CompanyModel) FieldArray() []string {
	return _companyFields.Array()
}

func (s *Company) FieldArray() []string {
	return _companyFields.Array()
}

func (s *CompanyModel) PK() (*q.Field, interface{}) {
	m, _ := s.Queryset().GetMain()
	return m.Model.PK()
}

func (s *Company) PK() (*q.Field, interface{}) {
	return s.Fields().ID, s.ID
}

func (s *CompanyModel) Queryset() *q.Query {
	if s.query == nil {
		s.query = q.NewQuery(&Company{})
	}
	return s.query
}

func (s *CompanyModel) Where(cs ...*q.Condition) *CompanyModel {
	s.Queryset().Where(cs...)
	return s
}

func (s *CompanyModel) getJs(m q.QueryModel, f *q.Field) (*q.Joinset, error) {
	jm, ok := _companyRels[m.TableName()]
	if ok {
		if f == nil {
			return jm[_companyDefaultRels[m.TableName()]], nil
		} else {
			if j, ok := jm[f.String()]; ok {
				return j, nil
			}
		}
	}
	return nil, q.ErrRelNotFound
}

var _companyRels = map[string]q.JoinsetMap{
}

var _companyDefaultRels = map[string]string{
}

func (s *CompanyModel) Join(qm q.QueryModel, direction string, fs ...*q.Field) *CompanyModel {
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
