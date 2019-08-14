package example

import (
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
	).Queryset().Fields(
		student.Fields().Name,
		school.Fields().Name,
		company.Fields().Name,
		guild.Fields().Name,
		team.Fields().Name,
	).WithoutModelFields().Query()
	if err != nil {
		t.Error(err)
	}
	if qs != `select student.name,school.name,company.name,guild.name,team.name from student inner join school on student.expect = school.name inner join stc inner join company on stc.company = company.id on student.id = stc.person inner join stg inner join guild on stg.guild_id = guild.id on student.id = stg.student_id inner join student_team inner join team on student_team.team_id = team.id on student.id = student_team.student_id where school.name = 'school_asd' and company.name = 'company_asd' and guild.name = 'guild_asd' and team.name = 'team_asd'` {
		t.Error("wrong query")
	}
}
