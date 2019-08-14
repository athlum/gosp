package example

import (
	q "github.com/athlum/gosp/query"
)

type GuildModel struct {
	query *q.Query
}

const _guildTableName = "guild"

func (s *GuildModel) TableName() string {
	return _guildTableName
}

func (s *Guild) TableName() string {
	return _guildTableName
}

type guildFields struct {
	ID   *q.Field
	Name *q.Field
}

func (fs *guildFields) Array() []string {
	return []string{
		fs.ID.String(),
		fs.Name.String(),
	}
}

var _guildFields = &guildFields{
	ID:   q.TableField(_guildTableName, "id"),
	Name: q.TableField(_guildTableName, "name"),
}

func (s *GuildModel) Fields() *guildFields {
	return _guildFields
}

func (s *Guild) Fields() *guildFields {
	return _guildFields
}

func (s *GuildModel) FieldArray() []string {
	return _guildFields.Array()
}

func (s *Guild) FieldArray() []string {
	return _guildFields.Array()
}

func (s *GuildModel) PK() (*q.Field, interface{}) {
	m, _ := s.Queryset().GetMain()
	return m.Model.PK()
}

func (s *Guild) PK() (*q.Field, interface{}) {
	return s.Fields().ID, s.ID
}

func (s *GuildModel) Queryset() *q.Query {
	if s.query == nil {
		s.query = q.NewQuery(&Guild{})
	}
	return s.query
}

func (s *GuildModel) Where(cs ...*q.Condition) *GuildModel {
	s.Queryset().Where(cs...)
	return s
}

func (s *GuildModel) getJs(m q.QueryModel, f *q.Field) (*q.Joinset, error) {
	jm, ok := _guildRels[m.TableName()]
	if ok {
		if f == nil {
			return jm[_guildDefaultRels[m.TableName()]], nil
		} else {
			if j, ok := jm[f.String()]; ok {
				return j, nil
			}
		}
	}
	return nil, q.ErrRelNotFound
}

var _guildRels = map[string]q.JoinsetMap{
}

var _guildDefaultRels = map[string]string{
}

func (s *GuildModel) Join(qm q.QueryModel, direction string, fs ...*q.Field) *GuildModel {
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