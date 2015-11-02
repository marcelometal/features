package main

import (
	"fmt"

	"github.com/albertoleal/features"
	"github.com/albertoleal/features/engine"
	"github.com/albertoleal/features/engine/memory"
)

func main() {
	Features := features.New(memory.New())

	feature := engine.FeatureFlag{
		Key:     "Feature X",
		Enabled: false,
		Users:   []*engine.User{&engine.User{Id: "alice@example.org"}},
	}
	Features.Save(feature)

	active, _ := Features.IsActive("Feature X")
	fmt.Printf("Is `Feature X` enabled? %t \n", active)

	deactive, _ := Features.IsInactive("Feature X")
	fmt.Printf("Is `Feature X` disabled? %t \n", deactive)

	Features.With("Feature X", func() {
		fmt.Println("`Feature X` is enabled!")
	})

	Features.Without("Feature X", func() {
		fmt.Println("`Feature X` is disabled!")
	})

	fmt.Printf("Does `alice@example.org` have access to `Feature X`? %t \n", Features.UserHasAccess("Feature X", "alice@example.org"))

}
