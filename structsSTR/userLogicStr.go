package structsSTR

import (
	"math/rand"
	"time"
)

type User struct {
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
	SetRandId()
	SetID(string)
	SetPassword(string)
	SetName(string)
}

func (u *User) GetPrefix() string {
	return u.prefix
}

func (u *User) GetId() string {
	return u.Id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPrefix(prefix string) {
	u.prefix = prefix
}

func (u *User) SetRandId() {
	rand.Seed(time.Now().UnixNano())
	randStr := "qwertyuiopasdfghjklzxcvbnm1234567890"

	strIdrand := ""

	for i := 0; i < 12; i++ {
		strIdrand += string(randStr[rand.Intn(len(randStr))])
	}

	u.Id = strIdrand
}

func (u *User) SetID(id string) {
	u.Id = id
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetName(name string) {
	u.name = name
}

func NewUser(Map map[string]*User, Name, password string) {
	rand.Seed(time.Now().UnixNano())
	randStr := "abcdefgh1234567890"

	strIdrand := ""

	for i := 0; i < 8; i++ {
		strIdrand += string(randStr[rand.Intn(len(randStr))])
	}

	b := &User{Name, strIdrand, password, " "}

	AddUserToMap(b, Map)
}

func AddUserToMap(u *User, m map[string]*User) {
	m[u.name] = u
}
