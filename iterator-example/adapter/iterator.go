package adapter

import "iterator-example/domain"

type User struct {
	Name string
	Age  int
}

type UserCollection struct {
	Users []*User
}

func (u *UserCollection) CreateIterator() domain.Iterator {
	return &UserIterator{
		users: u.Users,
	}
}

type UserIterator struct {
	users []*User
	index int
}

func (u *UserIterator) HasNext() bool {
	return u.index < len(u.users)
}

func (u *UserIterator) GetNext() interface{} {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
