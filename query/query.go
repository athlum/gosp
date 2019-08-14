package query

import (
	"fmt"
	"github.com/athlum/gosp/types"
	"github.com/athlum/gosp/utils"
	"github.com/juju/errors"
	"strings"
)

const (
	ErrNilSqlExecutor types.Error = "nil sqlExecutor"
	ErrNilModel       types.Error = "nil model"
	ErrRelNotFound    types.Error = "found no rel table"
	ErrNilFields      types.Error = "Need fields when without model"
)

const (
	EQ    = "="
	NE    = "<>"
	GT    = ">"
	GTE   = ">="
	LT    = "<"
	LTE   = "<="
	LIKE  = "like"
	IN    = "in"
	NOTIN = "not in"
	IS    = "is"
	ISNOT = "is not"
)

const (
	ASC  = "asc"
	DESC = "desc"
)

type Query struct {
	main               *Modelset
	fields             []string
	orderBy            []string
	groupBy            []string
	withoutModelFields bool
	limit              int
	offset             int
	err                error
}

func NewQuery(m Model) *Query {
	return &Query{
		main: NewModelset(m, nil, nil),
	}
}

func (q *Query) Fork() *Query {
	return &Query{
		main:               q.main.Fork(),
		fields:             utils.CopyStringArray(q.fields),
		orderBy:            utils.CopyStringArray(q.orderBy),
		groupBy:            utils.CopyStringArray(q.groupBy),
		withoutModelFields: q.withoutModelFields,
		limit:              q.limit,
		offset:             q.offset,
		err:                q.err,
	}
}

func (q *Query) SetErr(err error) *Query {
	q.err = err
	return q
}

func (q *Query) Pagination(page, size int) *Query {
	q.Offset((page - 1) * size).Limit(size)
	return q
}

func (q *Query) WithoutModelFields() *Query {
	q.withoutModelFields = true
	return q
}

func (q *Query) OrderBy(field interface{}, typ string) *Query {
	if q.orderBy == nil {
		q.orderBy = []string{}
	}
	q.orderBy = append(q.orderBy, fmt.Sprintf("%v %v", escape(field), typ))
	return q
}

func (q *Query) GroupBy(fs ...interface{}) *Query {
	if q.groupBy == nil {
		q.groupBy = []string{}
	}
	for _, f := range fs {
		q.groupBy = append(q.groupBy, escape(f))
	}
	return q
}

func (q *Query) Limit(l int) *Query {
	q.limit = l
	return q
}

func (q *Query) Offset(o int) *Query {
	q.offset = o
	return q
}

func (q *Query) Fields(fs ...interface{}) *Query {
	if q.fields == nil {
		q.fields = []string{}
	}
	for _, f := range fs {
		q.fields = append(q.fields, escape(f))
	}
	return q
}

func (q *Query) GetMain() (*Modelset, error) {
	if q.main == nil {
		return nil, ErrNilModel
	}
	return q.main, nil
}

func (q *Query) Where(cs ...*Condition) *Query {
	main, err := q.GetMain()
	if err != nil {
		q.SetErr(err)
	} else {
		main.Where(cs...)
	}
	return q
}

func (q *Query) Join(js ...*Joinset) *Query {
	main, err := q.GetMain()
	if err != nil {
		q.SetErr(err)
	} else {
		main.Join(js...)
	}
	return q
}

func ifAdd(query string, v []string, s string) string {
	if v != nil && len(v) > 0 {
		query = fmt.Sprintf("%v %v %v", query, s, strings.Join(v, ","))
	}
	return query
}

func (q *Query) QueryError(err error, qs string) error {
	return errors.Annotatef(err, "query: %v", qs)
}

func (q *Query) fieldQuery(fields ...string) (string, error) {
	if q.err != nil {
		return "", q.err
	}
	if q.main == nil {
		return "", ErrNilModel
	}
	s, err := q.main.Query(fields...)
	if err != nil {
		return "", err
	}
	s = ifAdd(s, q.groupBy, "group by")
	s = ifAdd(s, q.orderBy, "order by")
	if q.limit > 0 {
		s = fmt.Sprintf("%v limit %v", s, q.limit)
		if q.offset > 0 {
			s = fmt.Sprintf("%v offset %v", s, q.offset)
		}
	}
	return s, nil
}

func (q *Query) Query() (string, error) {
	var fields []string
	if !q.withoutModelFields {
		if q.main == nil {
			return "", ErrNilModel
		}
		fs, err := q.main.FieldArray()
		if err != nil {
			return "", err
		}
		fields = fs
	}
	if q.fields != nil && len(q.fields) > 0 {
		fields = append(fields, q.fields...)
	} else if q.withoutModelFields {
		return "", ErrNilFields
	}
	return q.fieldQuery(fields...)
}

func (q *Query) CountQuery(fields ...string) (string, error) {
	if q.main == nil {
		return "", ErrNilModel
	}
	pk, _, err := q.main.PK()
	if err != nil {
		return "", err
	}
	if len(fields) > 0 {
		pk = TableField("", fields[0])
	}
	return q.fieldQuery(pk.Count().String())
}

func escape(v interface{}) string {
	if f, ok := v.(*Field); ok {
		return f.String()
	}
	return v.(string)
}
