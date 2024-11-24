#!/bin/bash
source .env

sleep 2 && goose -dir ./migrations postgres "${PG_DSN}" up -v