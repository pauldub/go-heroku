package addons

import "heroku/types"

var (
	AcmeAddon = &types.Addon{Name: "example:basic", Description: "Example Basic", Url: "http://devcenter.heroku.com/articles/example-basic", State: "public", Beta: false}
)
