# Go blueprint

A new BP for a full stack Go + React program

## Development setup

- [Golang >= 1.21](https://go.dev/)
- [Air](https://github.com/air-verse/air/tree/master?tab=readme-ov-file#installation)
- [Goose -> Migrations](https://github.com/pressly/goose?tab=readme-ov-file#install)
- [Sqlc -> typed queries](https://docs.sqlc.dev/en/latest/overview/install.html)
- [Golangci-lint](https://golangci-lint.run/welcome/install/#local-installation)
- [Node >= 18](https://nodejs.org/en/download/package-manager)
- [Pnpm](https://pnpm.io/installation)
- make (Optional)

## Used golang libraries
- Fiber
- Viper
- Sqlc

## Used React libraries
- [React hook form](https://react-hook-form.com)
- [Shadcn UI](https://ui.shadcn.com)
- [Tanstack Query](https://tanstack.com/query/latest)
- [React router](https://reactrouter.com)
- [Sonner](https://sonner.emilkowal.ski/)

## Useful flows

### Adding a new typed query (sqlc)

1) Add your new query to db/queries/{target}.sql
2) run `sqlc generate`
3) Enjoy your statically typed query

### Adding a migration

1) make create-migration
> [!NOTE]
> For nix users using the devshell replace `make goose` with `goose -dir ./db/migrations postgres create my_migration_name sql`
2) Edit the newly made migration that can be found in the `db/migrations` folder
3) Update the queries in the `db/queries` accordingly
4) Run `sqlc generate` to generate the new table structs


## Deployment

It is recommended to run the application in an docker container

It need at least the following additional resources:
- Postgres Database
- Redis
