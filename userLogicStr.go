package structsSTR

import "math/rand"

type user struct {
	name     string
	Id       string
	password string
	prefix   string
}

type userInt interface {
	GetPrefix() string
	GetId() string
	GetName() string
	GetPassword() string

	SetPrefix(string)
	SetRandId() string
	SetID(string) string
}

func (u *user) GetPrefix() string {
	return u.prefix
}

func (u *user) GetId() string {
	return u.Id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetPassword() string {
	return u.password
}

func (u *user) SetPrefix(prefix string) {
	u.prefix = prefix
}

func (u *user) SetRandId() string {
	randStr := "qwertyuiopasdfghjklzxcvbnm1234567890"

	strIdrand := ""

	for i := 0; i == 12; i++ {
		strIdrand += string(randStr[rand.Intn(len(randStr))])
		u.Id += strIdrand
	}
	return u.Id
}

func (u *user) SetID(id string) {
	u.Id = id
}

func NewUser(Map map[string]*user, Name, password string) {
	randStr := "qwertyuiopasdfghjklzxcvbnm1234567890"

	strIdrand := ""

	for i := 0; i < 12; i++ {
		strIdrand += string(randStr[rand.Intn(len(randStr))])
	}

	b := user{Name, strIdrand, password, " "}

	AddUserAMap(b, Map)

}

func AddUserAMap(u user, m map[string]*user) {
	m[u.GetId()] = &u
}
