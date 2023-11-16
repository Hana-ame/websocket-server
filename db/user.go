package db

type User struct {
	Name     string `json:"name"`
	Account  string `json:"account" gorm:"primaryKey"`
	Password string `json:"password"`
}

// CRUD

func CreateUser(o *User) error {
	tx := db.Create(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadUser(o *User) error {
	tx := db.Take(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateUser(o *User) error {
	tx := db.Save(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteUser(o *User) error {
	tx := db.Delete(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
