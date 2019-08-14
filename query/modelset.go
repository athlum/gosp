package query

import (
	"fmt"
	"strings"
)

type Modelset struct {
	Model      Model
	Conditions Conditions
	Joins      Joinsets
}

func NewModelset(m Model, cs Conditions, joins Joinsets) *Modelset {
	return &Modelset{
		Model:      m,
		Conditions: cs,
		Joins:      joins,
	}
}

func (ms *Modelset) Fork() *Modelset {
	return &Modelset{
		Model:      ms.Model,
		Conditions: ms.Conditions.Fork(),
		Joins:      ms.Joins.Fork(),
	}
}

func (ms *Modelset) SameModel(m *Modelset) bool {
	return ms.Model.TableName() == m.Model.TableName()
}

func (ms *Modelset) JoinReplace(m *Modelset) (*Modelset, error) {
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
			if err := j.JoinModelset(m); err != nil {
				return nil, err
			}
		}
	}
	return ms, nil
}

func (ms *Modelset) FieldArray() ([]string, error) {
	if ms.Model == nil {
		return nil, ErrNilModel
	}
	fs := ms.Model.FieldArray()
	return fs, nil
}

func (ms *Modelset) PK() (*Field, interface{}, error) {
	if ms.Model == nil {
		return nil, nil, ErrNilModel
	}
	f, v := ms.Model.PK()
	return f, v, nil
}

func (ms *Modelset) Where(cs ...*Condition) *Modelset {
	ms.Conditions = append(ms.Conditions, cs...)
	return ms
}

func (ms *Modelset) Join(rs ...*Joinset) *Modelset {
	ms.Joins = append(ms.Joins, rs...)
	return ms
}

func (ms *Modelset) Query(fields ...string) (string, error) {
	if ms.Model == nil {
		return "", ErrNilModel
	}
	tn, cs := ms.Tuple()
	qs := fmt.Sprintf("select %s from %s where %s", strings.Join(fields, ","), tn, Conditions(cs).String())
	return qs, nil
}

func (ms *Modelset) Tuple() (string, Conditions) {
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
