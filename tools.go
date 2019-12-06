// +build tools

package main

import (
	// go run github.com/vektah/dataloaden AccountLoader string '*github.com/oreqizer/goiler/graphql/schemas/account.Account'
	_ "github.com/vektah/dataloaden"
)

// Other generators:

// go run cmd/schemagen/main.go -struct Account -models 'github.com/oreqizer/goiler/models'
