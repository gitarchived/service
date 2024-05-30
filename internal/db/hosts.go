package db

func (d *DB) IsValidHost(host string) bool {
	if err := d.Where("prefix = ?", host).First(&Host{}).Error; err != nil {
		return false
	}

	return true
}
