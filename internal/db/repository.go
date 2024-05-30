package db

func (d *DB) RepositoryExists(host, owner, name string) bool {
	if err := d.Where("host = ? AND owner = ? AND name = ?", host, owner, name).First(&Repository{}).Error; err != nil {
		return false
	}

	return true
}

func (d *DB) CreateRepository(host, owner, name string) (*Repository, error) {
	repo := &Repository{
		Host:  host,
		Owner: owner,
		Name:  name,
	}

	if err := d.Create(repo).Error; err != nil {
		return nil, err
	}

	return repo, nil
}
