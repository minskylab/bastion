package bastion

import (
	"context"
	"fmt"
	"time"

	"github.com/minskylab/bastion/prisma/db"
)

func test() {
	client := db.NewClient()

	organizationID := "organizationID"

	project, err := client.Project.CreateOne(
		db.Project.Name.Set("test-project"),
		db.Project.EstimatedDuration.Set(int(5*30*24*time.Hour)),
		db.Project.Organization.Link(db.Organization.ID.Equals(organizationID)),
	).Exec(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("project created id: %s\n", project.ID)
}
