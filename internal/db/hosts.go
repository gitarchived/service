package db

func (d *DB) IsValidHost(host string) bool {
	if err := d.Where("name = ?", host).First(&Host{}).Error; err != nil {
		return false
	}

	return true
}

func (d *DB) GetHostByName(name string) (Host, error) {
	var host Host

	if err := d.Where("name = ?", name).First(&host).Error; err != nil {
		return host, err
	}

	return host, nil
}
