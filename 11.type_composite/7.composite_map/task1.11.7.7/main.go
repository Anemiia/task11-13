package main

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	mapa := make(map[string]bool)
	result := make([]User, 0, len(users))

	for _, user := range users {
		if !mapa[user.Nickname] {
			mapa[user.Nickname] = true
			result = append(result, user)
		}
	}

	return result[:len(result):len(result)]
}