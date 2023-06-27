package users

import (
	"github.com/AlejandroCarballo/ExamplesGO/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Ale", time.Now(), true)
}
