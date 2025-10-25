package repo

import "errors"

func (r *Repo) Login(email string) (string, error) {
	var hash string
	err := r.db.QueryRow(
		"SELECT password_hash FROM users WHERE email=$1", email).Scan(&hash)
	if err != nil || hash == "" {
		return "", errors.New("user not exists")
	}
	return hash, nil
}
