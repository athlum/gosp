package example

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	var (
		student = Table.Student()
		school  = Table.School()
		company = Table.Company()
		guild   = Table.Guild()
		team    = Table.Team()
	)
	qs, err := student.Join(
		school.Where(
			school.Fields().Name.EQ("school_asd"),
		), student.Fields().Expect,
	).Join(
		company.Where(
			company.Fields().Name.EQ("company_asd"),
		),
	).Join(
		guild.Where(
			guild.Fields().Name.EQ("guild_asd"),
		),
	).Join(
		team.Where(
			team.Fields().Name.EQ("team_asd"),
		),
	).Query().Fields(
		student.Fields().Name.String(),
		school.Fields().Name.String(),
		company.Fields().Name.String(),
		guild.Fields().Name.String(),
		team.Fields().Name.String(),
	).WithoutModelFields().Query()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(qs)
}
