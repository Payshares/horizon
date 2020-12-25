#! /usr/bin/env bash

set -e

SCENARIO=$1
CORE_SQL=./src/github.com/payshares/horizon/test/scenarios/$SCENARIO-core.sql
HORIZON_SQL=./src/github.com/payshares/horizon/test/scenarios/$SCENARIO-horizon.sql

echo "psql $PAYSHARES_CORE_DATABASE_URL < $CORE_SQL" 
psql $PAYSHARES_CORE_DATABASE_URL < $CORE_SQL 
echo "psql $DATABASE_URL < $HORIZON_SQL"
psql $DATABASE_URL < $HORIZON_SQL 
