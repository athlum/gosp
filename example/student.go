package example

//@Table(student)
//@Rel(School)=Student.SchoolID
//@Rel(School)=Student.Expect,School.Name
//@Mul(Team)
//@Mul(Guild, STG)
//@Mul(Company, STC)=STC.Person,STC.Company
type Student struct {
	ID       int64  `db:"id, pk, autoincrement"`
	Name     string `db:"name, length:255"`
	SchoolID int64  `db:"school_id, default:0"`
	Expect   string `db:"expect, length:255, default:''"`
}