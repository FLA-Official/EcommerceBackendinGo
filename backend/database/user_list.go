package database

var users []User

func (u User) Store() User {
	if u.ID != 0 {
		return u
	}
	u.ID = len(users) + 1
	users = append(users, u)
	return u
}

func Find(email, password string) *User {
	for _, u := range users {
		if u.Email == email && u.Password == password {
			return &u
		}
	}

	return nil
}
