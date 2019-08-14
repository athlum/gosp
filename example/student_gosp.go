package example

import (
	q "github.com/athlum/gosp/query"
)

type StudentModel struct {
	query *q.Query
}

const _studentTableName = "student"

func (s *StudentModel) TableName() string {
	return _studentTableName
}

func (s *Student) TableName() string {
	return _studentTableName
}

type studentFields struct {
	ID       *q.Field
	Name     *q.Field
	SchoolID *q.Field
	Expect   *q.Field
}

func (fs *studentFields) Array() []string {
	return []string{
		fs.ID.String(),
		fs.Name.String(),
		fs.SchoolID.String(),
		fs.Expect.String(),
	}
}

var _studentFields = &studentFields{
	ID:       q.TableField(_studentTableName, "id"),
	Name:     q.TableField(_studentTableName, "name"),
	SchoolID: q.TableField(_studentTableName, "school_id"),
	Expect:   q.TableField(_studentTableName, "expect"),
}

func (s *StudentModel) Fields() *studentFields {
	return _studentFields
}

func (s *Student) Fields() *studentFields {
	return _studentFields
}

func (s *StudentModel) FieldArray() []string {
	return _studentFields.Array()
}

func (s *Student) FieldArray() []string {
	return _studentFields.Array()
}

func (s *StudentModel) PK() (*q.Field, interface{}) {
	m, _ := s.Query().GetMain()
	return m.Model.PK()
}

func (s *Student) PK() (*q.Field, interface{}) {
	return s.Fields().ID, s.ID
}

func (s *StudentModel) Query() *q.Query {
	if s.query == nil {
		s.query = q.NewQuery(&Student{})
	}
	return s.query
}

func (s *StudentModel) Where(cs ...*q.Condition) *StudentModel {
	s.Query().Where(cs...)
	return s
}

func (s *StudentModel) getJs(m q.QueryModel, f *q.Field) (*q.Joinset, error) {
	jm, ok := _studentRels[m.TableName()]
	if ok {
		if f == nil {
			return jm[_studentDefaultRels[m.TableName()]], nil
		} else {
			if j, ok := jm[f.String()]; ok {
				return j, nil
			}
		}
	}
	return nil, q.ErrRelNotFound
}

var _studentRels = map[string]q.JoinsetMap{
	_schoolTableName: q.JoinsetMap{
		_studentFields.SchoolID.String(): q.NewJoinset(
			q.NewModelSet(&School{}, nil, nil),
			_studentFields.SchoolID.EQ(_schoolFields.ID),
		),
		_studentFields.Expect.String(): q.NewJoinset(
			q.NewModelSet(&School{}, nil, nil),
			_studentFields.Expect.EQ(_schoolFields.Name),
		),
	},
	_teamTableName: q.JoinsetMap{
		_studentFields.ID.String(): q.NewJoinset(
			q.NewModelSet(&StudentTeam{}, nil, []*q.Joinset{
				q.NewJoinset(
					q.NewModelSet(&Team{}, nil, nil),
					_studentTeamFields.TeamID.EQ(_teamFields.ID),
				),
			}),
			_studentFields.ID.EQ(_studentTeamFields.StudentID),
		),
	},
	_guildTableName: q.JoinsetMap{
		_studentFields.ID.String(): q.NewJoinset(
			q.NewModelSet(&STG{}, nil, []*q.Joinset{
				q.NewJoinset(
					q.NewModelSet(&Guild{}, nil, nil),
					_stgFields.GuildID.EQ(_guildFields.ID),
				),
			}),
			_studentFields.ID.EQ(_stgFields.StudentID),
		),
	},
	_companyTableName: q.JoinsetMap{
		_studentFields.ID.String(): q.NewJoinset(
			q.NewModelSet(&STC{}, nil, []*q.Joinset{
				q.NewJoinset(
					q.NewModelSet(&Company{}, nil, nil),
					_stcFields.Company.EQ(_companyFields.ID),
				),
			}),
			_studentFields.ID.EQ(_stcFields.Person),
		),
	},
}

var _studentDefaultRels = map[string]string{
	_schoolTableName:  _studentFields.SchoolID.String(),
	_teamTableName:    _studentFields.ID.String(),
	_guildTableName:   _studentFields.ID.String(),
	_companyTableName: _studentFields.ID.String(),
}

func (s *StudentModel) Join(qm q.QueryModel, fs ...*q.Field) *StudentModel {
	return s.JoinDir(qm, "", fs...)
}

func (s *StudentModel) JoinDir(qm q.QueryModel, direction string, fs ...*q.Field) *StudentModel {
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
			err = js.JoinModelSet(main)
		}
		if err != nil {
			qs.SetErr(err)
		} else {
			if direction != "" {
				js.Direction(direction)
			}
			qs.Join(js)
		}
	} else {
		qs.SetErr(err)
	}
	return s
}
