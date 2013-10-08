package client

import "encoding/json"
import "fmt"
import "net/http"
import "io/ioutil"
import "go-heroku/api"
import "go-heroku/types"

type ApiError struct {
	Id      string
	Message string `json:"error"`
}

func (err *ApiError) Error() string {
	return fmt.Sprintf("ApiError: %+v", err)
}

type UnknownError struct {
	Message string
}

func (err *UnknownError) Error() string {
	return fmt.Sprintf("UnknownError: %+v", err)
}

func NewApiError(res *http.Response) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	apiError := &ApiError{}
	err = json.Unmarshal(body, apiError)
	if err != nil {
		unknownError := &UnknownError{string(body)}
		return unknownError
	}

	return apiError
}

func decodeApiResponse(res *http.Response, data interface{}) error {
	err := json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		fmt.Println("JsonDecoderError:", err)
		return NewApiError(res)
	}

	return nil
}

/* Applications */

func (client *HerokuClient) ListApplications() (apps []types.Application, err error) {
	err = client.Do(api.ListApplications, &apps)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (client *HerokuClient) CreateApplication(application *types.Application) (app *types.Application, err error) {
	err = client.Do(api.CreateApplication(application), &app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (client *HerokuClient) RenameApplication(application *types.Application, newName string) (app *types.Application, err error) {
	err = client.Do(api.RenameApplication(application, newName), &app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (client *HerokuClient) TransferApplication(app *types.Application, collaborator *types.Collaborator) error {
	return client.Do(api.TransferApplication(app, collaborator), nil)
}

func (client *HerokuClient) SetApplicationMaintenance(app *types.Application, active bool) error {
	return client.Do(api.SetApplicationMaintenance(app, active), nil)
}

func (client *HerokuClient) DeleteApplication(app *types.Application) error {
	return client.Do(api.DeleteApplication(app), nil)
}

/* Collaborators */

func (client *HerokuClient) ListCollaborators(app *types.Application) (collaborators *[]types.Collaborator, err error) {
	err = client.Do(api.ListCollaborators(app), &collaborators)
	if err != nil {
		return nil, err
	}

	return collaborators, nil
}

func (client *HerokuClient) AddCollaborator(app *types.Application, collaborator *types.Collaborator) error {
	return client.Do(api.AddCollaborator(app, collaborator), nil)
}

func (client *HerokuClient) RemoveCollaborator(app *types.Application, collaborator *types.Collaborator) error {
	return client.Do(api.RemoveCollaborator(app, collaborator), nil)
}

/* Config */

func (client *HerokuClient) ListConfig(app *types.Application) (*[]types.Config, error) {
	configMap := make(map[string]interface{})

	err := client.Do(api.ListConfig(app), &configMap)
	if err != nil {
		return nil, err
	}

	configs := make([]types.Config, len(configMap))
	for k, v := range configMap {
		configs = append(configs, types.Config{k, fmt.Sprint(v)})
	}

	return &configs, nil
}

func (client *HerokuClient) AddConfig(app *types.Application, configVars *[]types.Config) (configs *[]types.Config, err error) {
	err = client.Do(api.AddConfig(app, configVars), &configs)
	if err != nil {
		return nil, err
	}

	return configs, nil
}

func (client *HerokuClient) DeleteConfig(app *types.Application, config *types.Config) error {
	return client.Do(api.DeleteConfig(app, config), nil)
}

/* Domains */

func (client *HerokuClient) ListDomains(app *types.Application) (domains *[]types.Domain, err error) {
	err = client.Do(api.ListDomains(app), &domains)
	if err != nil {
		return nil, err
	}

	return domains, nil
}

func (client *HerokuClient) AddDomain(app *types.Application, domain *types.Domain) error {
	return client.Do(api.AddDomain(app, domain), nil)
}

func (client *HerokuClient) DeleteDomain(app *types.Application, domain *types.Domain) error {
	return client.Do(api.DeleteDomain(app, domain), nil)
}

/* Keys */

func (client *HerokuClient) ListKeys() (keys *[]types.SshKey, err error) {
	err = client.Do(api.ListKeys(), &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (client *HerokuClient) AddKey(key *types.SshKey) error {
	return client.Do(api.AddKey(key), nil)
}

func (client *HerokuClient) DeleteKey(key *types.SshKey) error {
	return client.Do(api.DeleteKey(key), nil)
}

func (client *HerokuClient) DeleteAllKeys() error {
	return client.Do(api.DeleteAllKeys(), nil)
}

/* Processes */

func (client *HerokuClient) ListProcesses(app *types.Application) (ps *[]types.Dyno, err error) {
	err = client.Do(api.ListProcesses(app), &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (client *HerokuClient) RunProcess(app *types.Application, command string) error {
	return client.Do(api.RunProcess(app, command), nil)
}

func (client *HerokuClient) RestartProcess(app *types.Application, processName string, processType string) error {
	return client.Do(api.RestartProcess(app, processName, processType), nil)
}

func (client *HerokuClient) StopProcess(app *types.Application, processName string, processType string) error {
	return client.Do(api.StopProcess(app, processName, processType), nil)
}

func (client *HerokuClient) ScaleProcess(app *types.Application, processType string, quantity int) error {
	return client.Do(api.ScaleProcess(app, processType, quantity), nil)
}

/* Releases */

func (client *HerokuClient) ListReleases(app *types.Application) (releases *[]types.Release, err error) {
	err = client.Do(api.ListReleases(app), &releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (client *HerokuClient) GetRelease(app *types.Application, release *types.Release) (rel *types.Release, err error) {
	err = client.Do(api.GetRelease(app, release), &rel)
	if err != nil {
		return nil, err
	}

	return rel, nil
}

func (client *HerokuClient) RollbackToRelease(app *types.Application, release *types.Release) error {
	return client.Do(api.RollbackToRelease(app, release), nil)
}

/* Stacks */

func (client *HerokuClient) ListStacks(app *types.Application) (stacks *[]types.Stack, err error) {
	err = client.Do(api.ListStacks(app), &stacks)
	if err != nil {
		return nil, err
	}

	return stacks, nil
}

func (client *HerokuClient) MigrateStack(app *types.Application, stack *types.Stack) error {
	return client.Do(api.MigrateStack(app, stack), nil)
}

/* Addons */

func (client *HerokuClient) ListAllAddons() (addons *[]types.Addon, err error) {
	err = client.Do(api.ListAllAddons(), &addons)
	if err != nil {
		return nil, err
	}

	return addons, nil
}

func (client *HerokuClient) ListAddons(app *types.Application) (addons *[]types.Addon, err error) {
	err = client.Do(api.ListAddons(app), &addons)
	if err != nil {
		return nil, err
	}

	return addons, nil
}

func (client *HerokuClient) InstallAddon(app *types.Application, addon *types.Addon) error {
	return client.Do(api.InstallAddon(app, addon), nil)
}

func (client *HerokuClient) UpgradeAddon(app *types.Application, addon *types.Addon) error {
	return client.Do(api.UpgradeAddon(app, addon), nil)
}

func (client *HerokuClient) UninstallAddon(app *types.Application, addon *types.Addon) error {
	return client.Do(api.UninstallAddon(app, addon), nil)
}
