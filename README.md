# Go blueprint

A new BP for a full stack Go + React program

## Development setup

- [Golang >= 1.21](https://go.dev/)
- [Flox](https://flox.dev/) (optional)
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
- [Mantine UI](https://mantine.dev)
- [Tanstack Query](https://tanstack.com/query/latest)
- [React router](https://reactrouter.com)

## Development environment

If you want an easy development environment with a DB, BE dev server & FE dev server I recommend installing flox. After installing you can activate the environment with:
- `flox activate` if you don't want to run the background services
- `flox activate --start-services` to directly start the backgrond services when entering the env

To see the logs of all the services you can use `flox services logs --follow`

## Useful flows

### Setup

Make sure you have atleast go 1.24 installed
Then run `make setup`

### Adding a new typed query (sqlc)

1) Add your new query to db/queries/{target}.sql
2) run `make query`
3) Enjoy your statically typed query

### Adding a migration

1) make create-migration
> [!NOTE]
> For nix users using the devshell replace `make goose` with `goose -dir ./db/migrations postgres create my_migration_name sql`
2) Edit the newly made migration that can be found in the `db/migrations` folder
3) Update the queries in the `db/queries` accordingly
4) Run `sqlc generate` to generate the new table structs


## Deployment

It is recommended to run the application in an docker container.

If you have healthchecks or something similar running, it is recommended to run the run command with `--init`

It need at least the following additional resources:
- Postgres Database
- Redis
