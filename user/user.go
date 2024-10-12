package user

type User struct {
	ID   string
	Name string
}

var users []User = []User{
	{
		ID:   "a1b2c3d4e5",
		Name: "John Doe",
	},
	{
		ID:   "f6g7h8i9j0",
		Name: "Jane Smith",
	},
	{
		ID:   "k1l2m3n4o5",
		Name: "Mike Johnson",
	},
	{
		ID:   "p6q7r8s9t0",
		Name: "Emily Davis",
	},
	{
		ID:   "u1v2w3x4y5",
		Name: "Chris Lee",
	},
}

func ListUsers() []User {
	return users
}

func AddUser(user User) {
	users = append(users, user)
}
