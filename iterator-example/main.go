package main

import (
	"fmt"
	"iterator-example/adapter"
)

func main() {
	fmt.Println("=== Iterator Pattern ===")

	user1 := &adapter.User{Name: "a", Age: 30}
	user2 := &adapter.User{Name: "b", Age: 20}

	userCollection := &adapter.UserCollection{
		Users: []*adapter.User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext().(*adapter.User)
		fmt.Printf("User is %s\n", user.Name)
	}
}

