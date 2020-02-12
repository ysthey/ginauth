package ginauth

type SimpleAuthUser struct {
	// Auth
	ID       string
	Email    string
	Password string
}

// PutPID into user
func (u *SimpleAuthUser) PutPID(pid string) { u.ID = pid }

// PutPassword into user
func (u *SimpleAuthUser) PutPassword(password string) { u.Password = password }

// PutEmail into user
func (u *SimpleAuthUser) PutEmail(email string) { u.Email = email }

// GetPID from user
func (u SimpleAuthUser) GetPID() string { return u.ID }

// GetPassword from user
func (u SimpleAuthUser) GetPassword() string { return u.Password }

// GetEmail from user
func (u SimpleAuthUser) GetEmail() string { return u.Email }
