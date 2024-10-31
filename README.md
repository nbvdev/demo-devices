# Device Management Service (Demo)

## Requirements:
- golang-migrate
- mysql (optional, can be used from container for development purpose)
- docker compose (optional, if necessary to create containers for local development )

## Configuration:
Configuration is stored in infrastructure/.env file
This configuration used for docker images creation as well as for make commands execution

## Preparation steps:
- make prepare (this install golang-migrate)
- make up (this creates container with mysql for local development)

## Migrations:
- make migrate_create ( create new migration )
- make migrate_up ( apply migrations manually )
- make migrate_down ( rollback migration manually )

### Important Note: 
Migrations also will be automatically applied when application starts

### Testing:
- make test (runs all available test for internals)

### Mockery:
- make mockery (generate mocks for testing purpose)