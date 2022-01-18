package cmd

import (
	"context"

	"github.com/google/uuid"
	"github.com/minskylab/bastion/ent"
)

func test() {
	client := ent.NewClient()

	// client.Organization.

	org, _ := uuid.NewUUID()
	mem, _ := uuid.NewUUID()

	client.Organization.UpdateOneID(org).AddDeveloperIDs(mem).SaveX(context.Background())
}
