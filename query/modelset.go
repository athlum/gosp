package query

import (
	"fmt"
	"strings"
)

type ModelSet struct {
	Model      Model
	Conditions Conditions
	Joins      Joinsets
}

func NewModelSet(m Model, cs Conditions, joins Joinsets) *ModelSet {
	return &ModelSet{
		Model:      m,
		Conditions: cs,
		Joins:      joins,
	}
}

func (ms *ModelSet) Fork() *ModelSet {
	return &ModelSet{
		Model:      ms.Model,
		Conditions: ms.Conditions.Fork(),
		Joins:      ms.Joins.Fork(),
	}
}

func (ms *ModelSet) SameModel(m *ModelSet) bool {
	return ms.Model.TableName() == m.Model.TableName()
}

func (ms *ModelSet) JoinReplace(m *ModelSet) (*ModelSet, error) {
	if ms.Model == nil {
		return nil, ErrNilModel
	}
	if ms.SameModel(m) {
		cm := m.Fork()
		cm.Where([]*Condition(ms.Conditions)...)
		cm.Join([]*Joinset(ms.Joins)...)
		return cm, nil
	} else {
		for _, j := range ms.Joins {
			if err := j.JoinModelSet(m); err != nil {
				return nil, err
			}
		}
	}
	return ms, nil
}

func (ms *ModelSet) FieldArray() ([]string, error) {
	if ms.Model == nil {
		return nil, ErrNilModel
	}
	fs := ms.Model.FieldArray()
	return fs, nil
}

func (ms *ModelSet) PK() (*Field, interface{}, error) {
	if ms.Model == nil {
		return nil, nil, ErrNilModel
	}
	f, v := ms.Model.PK()
	return f, v, nil
}

func (ms *ModelSet) Where(cs ...*Condition) *ModelSet {
	ms.Conditions = append(ms.Conditions, cs...)
	return ms
}

func (ms *ModelSet) Join(rs ...*Joinset) *ModelSet {
	ms.Joins = append(ms.Joins, rs...)
	return ms
}

func (ms *ModelSet) Query(fields ...string) (string, error) {
	if ms.Model == nil {
		return "", ErrNilModel
	}
	tn, cs := ms.Tuple()
	qs := fmt.Sprintf("select %s from %s where %s", strings.Join(fields, ","), tn, Conditions(cs).String())
	return qs, nil
}

func (ms *ModelSet) Tuple() (string, Conditions) {
	var (
		cs = []*Condition(ms.Conditions)
		tn = ms.Model.TableName()
	)

	for _, j := range ms.Joins {
		jtn, jcs := j.Tuple()
		tn = fmt.Sprintf("%s %s", tn, jtn)
		cs = append(cs, jcs...)
	}
	return tn, cs
}
