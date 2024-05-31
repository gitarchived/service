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

func (d *DB) GetRepository(host, owner, name string) (*Repository, error) {
	repo := &Repository{}

	if err := d.Where("host = ? AND owner = ? AND name = ?", host, owner, name).First(repo).Error; err != nil {
		return nil, err
	}

	return repo, nil
}

func (d *DB) SearchRepositories(query string, index int) ([]*Repository, bool, error) {
	var repos []*Repository

	if err := d.Where("name LIKE ? OR owner LIKE ?", "%"+query+"%", "%"+query+"%").Offset((index - 1) * 10).Limit(10).Find(&repos).Error; err != nil {
		return nil, false, err
	}

	// Pagination
	if len(repos) > 10 {
		return repos[(index-1)*10 : index*10], true, nil
	}

	return repos, false, nil
}

func (d *DB) RepositoriesByOwner(host string, owner string) ([]*Repository, error) {
	var repos []*Repository

	if err := d.Where("host = ? AND owner = ?", host, owner).Find(&repos).Error; err != nil {
		return nil, err
	}

	return repos, nil
}
