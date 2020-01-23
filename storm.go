
	type User struct {
      ID int // primary key
      Group string `storm:"index"` // this field will be indexed
      Email string `storm:"unique"` // this field will be indexed with a unique constraint
      Name string // this field will not be indexed
	}

    db, err := storm.Open("my.db")
	user := User{ID: 10, Group: "staff", Email: "john@provider.com", Name: "John"}
	err := db.Save(&user)

	var user User
	err := db.One("Email", "john@provider.com", &user)
	var users []User
	err := db.Find("Group", "staff", &users)
	err = db.Select(q.Gte("ID", 10), q.Lte("ID", 100)).Find(&users)
