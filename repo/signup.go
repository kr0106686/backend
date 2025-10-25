package repo

func (r *Repo) Signup(name, email, pass string) error {
	_, err := r.db.Exec(`INSERT INTO users(name, email,password_hash) VALUES($1,$2,$3)`, name, email, pass)
	if err != nil {
		return err
	}
	return nil
}
