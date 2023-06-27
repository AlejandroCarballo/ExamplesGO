package modelos

import "time"

type User struct {
	Id        int
	Name      string
	CreatedAd time.Time
	Status    bool
}

func (u *User) AddUser(id int, name string, createdAd time.Time, status bool) {
	u.Id = id
	u.Name = name
	u.CreatedAd = createdAd
	u.Status = status

}
