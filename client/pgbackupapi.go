package client

import "go-heroku/api"
import "go-heroku/types/addons"

func (client *PgBackupClient) ListBackups() (backups *[]addons.PgBackup, err error) {
	err = client.Do(api.ListBackups(), &backups)
	if err == nil {
		return nil, err
	}
	return backups, nil
}

func (client *PgBackupClient) GetBackup(name string) (backup *addons.PgBackup, err error) {
	err = client.Do(api.GetBackup(name), &backup)
	if err != nil {
		return nil, err
	}
	return backup, nil
}

func (client *PgBackupClient) GetLatestBackup() (backup *addons.PgBackup, err error) {
	err = client.Do(api.GetLatestBackup(), &backup)
	if err != nil {
		return nil, err
	}
	return backup, nil
}
