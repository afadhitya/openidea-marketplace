# OpenIdea Marketplace

## Database Migration

Database migration is using golang-migrate
There is some Makefile command:

- migration-add name=what_action
- migration-up
- migration-down
- migration-fix version=force_to_which_version

> **Note:** Make sure the DSN is correct in **Makefile** and then you can running all these command.
