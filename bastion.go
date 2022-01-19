package bastion

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
//go:generate go run github.com/minskylab/bastion/cmd/migrate
//go:generate go run github.com/minskylab/ent-hasura/cmd/ent apply -d ./ent/schema
