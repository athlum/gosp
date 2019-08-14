package example

type tables struct{}

func (t *tables) Student() *StudentModel {
	return &StudentModel{}
}
func (t *tables) School() *SchoolModel {
	return &SchoolModel{}
}
func (t *tables) Company() *CompanyModel {
	return &CompanyModel{}
}
func (t *tables) Guild() *GuildModel {
	return &GuildModel{}
}
func (t *tables) Team() *TeamModel {
	return &TeamModel{}
}
func (t *tables) StudentTeam() *StudentTeamModel {
	return &StudentTeamModel{}
}
func (t *tables) STG() *STGModel {
	return &STGModel{}
}
func (t *tables) STC() *STCModel {
	return &STCModel{}
}

var Table tables
