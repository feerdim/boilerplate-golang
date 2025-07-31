# DATABASE MIGRATION

In this project, we use [golang-migrate](https://github.com/golang-migrate/migrate) to manage database migration.

## Create Migration

```bash
make migrate.create name=${example}
```

use prefix name:
1. `create_` for create table
2. `alter_` for alter table like add/drop column, change column type, etc
3. `drop_` for drop table

## Run Migration

```bash
make migrate.up
```

## Rollback Migration

```bash
make migrate.down
```
