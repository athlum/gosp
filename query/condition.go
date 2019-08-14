package query

import (
	"fmt"
)

const (
	Connector_AND = "and"
	Connector_OR  = "or"
)

type Conditions []*Condition

func (cs Conditions) Fork() Conditions {
	ncs := make([]*Condition, len(cs))
	for i, c := range cs {
		ncs[i] = c.Fork()
	}
	return ncs
}

func (cs Conditions) String() string {
	s := ""
	for i, c := range cs {
		s = fmt.Sprintf("%s%s", s, c.String(i))
	}
	return s
}

type Condition struct {
	key        *string
	operator   string
	value      interface{}
	connector  string
	conditions Conditions
	rel        bool
}

func Case(key, operator string, value interface{}) *Condition {
	return &Condition{
		key:       &key,
		operator:  operator,
		value:     value,
		connector: Connector_AND,
	}
}

func OrCase(key, operator string, value interface{}, cs ...*Condition) *Condition {
	return &Condition{
		key:        &key,
		operator:   operator,
		value:      value,
		connector:  Connector_OR,
		conditions: cs,
	}
}

func caseCombine(connector string, cs ...*Condition) *Condition {
	if len(cs) == 1 {
		cs[0].connector = connector
		return cs[0]
	}
	return &Condition{
		connector:  connector,
		conditions: cs,
	}
}

func Or(cs ...*Condition) *Condition {
	return caseCombine(Connector_OR, cs...)
}

func And(cs ...*Condition) *Condition {
	return caseCombine(Connector_AND, cs...)
}

func (c *Condition) Fork() *Condition {
	k := *(c.key)
	return &Condition{
		key:        &k,
		operator:   c.operator,
		value:      c.value,
		connector:  c.connector,
		conditions: c.conditions.Fork(),
		rel:        c.rel,
	}
}

func (c *Condition) String(index int) string {
	var (
		v = c.value
		s string
	)
	if c.key != nil {
		_, isInt := c.value.(int)
		_, isField := c.value.(*Field)
		if !(isField || isInt || c.rel || c.value == "null" || c.operator == IN || c.operator == NOTIN) {
			v = fmt.Sprintf("'%v'", c.value)
		}
		s = fmt.Sprintf("%v %v %v", *c.key, c.operator, v)
	}
	if c.conditions != nil && len(c.conditions) > 0 {
		for i, cc := range c.conditions {
			sep := i
			if i == 0 && len(s) > 0 {
				sep = 1
			}
			s = fmt.Sprintf("%v%v", s, cc.String(sep))
		}
		if len(c.conditions) > 1 {
			s = fmt.Sprintf("(%v)", s)
		}
	}
	if index > 0 {
		s = fmt.Sprintf(" %v %v", c.connector, s)
	}
	return s
}

func (c *Condition) Relation() *Condition {
	c.rel = true
	return c
}

func (c *Condition) Sub(cs ...*Condition) *Condition {
	if c.conditions == nil {
		c.conditions = []*Condition{}
	}
	c.conditions = append(c.conditions, cs...)
	return c
}
