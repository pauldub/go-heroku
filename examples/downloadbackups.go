package main

import (
	"fmt"
	heroku "go-heroku/client"
	"go-heroku/types"
	addontypes "go-heroku/types/addons"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

type BackupDownload struct {
	App    *types.Application
	Backup *addontypes.PgBackup
	Addon  *types.Addon
	Url    string
}

func fetchLastBackup(dlClient *http.Client, client *heroku.HerokuClient, app *types.Application, addon *types.Addon) *BackupDownload {
	pgBackupClient, err := heroku.NewPgBackupClientFromHerokuClient(client, app, addon)
	if err != nil {
		fmt.Println("Error", app.Name, err)
		return nil
	}

	backup, err := pgBackupClient.GetLatestBackup()
	if err != nil {
		fmt.Println("failed to fetch backup for", app.Name)
		return nil
	}

	if backup.PublicUrl == "" {
		fmt.Println("no public url for", app.Name)
		return nil
	}

	res, err := dlClient.Get(backup.PublicUrl)
	if err != nil {
		fmt.Println(app.Name, res)
		return nil
	}

	file, err := os.Create(app.Name + ".dump")
	if err != nil {
		fmt.Println(app.Name, err)
		return nil
	}

	defer file.Close()

	fmt.Println("Downloading latest backup for", app.Name, backup.PublicUrl)
	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Println("copy error", err)
		return nil
	}
	return &BackupDownload{Backup: backup, App: app, Addon: addon, Url: backup.PublicUrl}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)
	client := heroku.NewHerokuClient(os.Genv("HEROKU_API_KEY"))

	results := make(chan *BackupDownload)
	completion := make(chan int, 100)
	wg := sync.WaitGroup{}
	dlClient := &http.Client{}

	apps, err := client.ListApplications()
	if err != nil {
		fmt.Println("An error occured:", err)
	} else if len(apps) == 0 {
		fmt.Println("No applications.")
	} else {
		for _, app := range apps {
			wg.Add(1)
			go func(app types.Application) {
				defer wg.Done()
				addons, err := client.ListAddons(&app)
				if err != nil {
					fmt.Println("Error fetching addonds for", app.Name)
					panic(err)
				}
				for _, addon := range *addons {
					if strings.HasPrefix(addon.Name, "pgbackups:") {
						fmt.Println("Application : ", app.Name)
						completion <- 1
						wg.Add(1)
						go func(app types.Application, addon types.Addon) {
							defer wg.Done()
							results <- fetchLastBackup(dlClient, client, &app, &addon)
							<-completion
						}(app, addon)
					}
				}
			}(app)
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		for {
			backup, ok := <-results
			if !ok {
				return
			}
			if backup != nil {
				fmt.Println("completed:", backup.App.Name)
			}
		}

	}
}
