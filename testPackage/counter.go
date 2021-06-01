package testPackage

type mycounter int

type Outcount int

func New() mycounter {
	return mycounter(100)
}

type TestUser struct {
	Name  string //公开的
	email string //未公开的
}

type otherUser struct {
	Name  string	//公开的
	Email string
}

type DUser struct {
	otherUser	//未公开的
	Level int
}

func (d *DUser) Setname(str string) {
	d.Name = str
}