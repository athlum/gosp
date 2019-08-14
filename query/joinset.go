package query

import (
	"fmt"
)

const (
	Join_Inner = "inner"
	Join_Left  = "left"
	Join_Right = "right"
)

type Joinsets []*Joinset

func (js Joinsets) Fork() Joinsets {
	njs := make([]*Joinset, len(js))
	for i, j := range js {
		njs[i] = j.Fork()
	}
	return njs
}

type JoinsetMap map[string]*Joinset

type Joinset struct {
	model     *Modelset
	direction string
	on        *Condition
}

func NewJoinset(model *Modelset, c *Condition) *Joinset {
	return &Joinset{
		model:     model,
		direction: Join_Inner,
		on:        c,
	}
}

func (js *Joinset) SetModelset(m *Modelset) *Joinset {
	js.model =m
	return js
}

func (js *Joinset) JoinModelset(m *Modelset) error {
	nm, err := js.model.JoinReplace(m)
	if err != nil {
		return err
	}
	js.model = nm
	return nil
}

func (js *Joinset) Fork() *Joinset {
	return &Joinset{
		model:     js.model.Fork(),
		direction: js.direction,
		on:        js.on.Fork(),
	}
}

func (js *Joinset) Direction(d string) *Joinset {
	js.direction = d
	return js
}

func (js *Joinset) Left() *Joinset {
	return js.Direction(Join_Left)
}

func (js *Joinset) Right() *Joinset {
	return js.Direction(Join_Right)
}

func (js *Joinset) Tuple() (string, []*Condition) {
	cs := []*Condition{}
	mtn, mcs := js.model.Tuple()
	cs = append(cs, mcs...)
	tn := fmt.Sprintf("%s join %s on %s", js.direction, mtn, js.on.String(0))
	return tn, cs
}
