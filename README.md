# atlas-drop-information
Mushroom Game - Drop Information Service

## Overview
A RESTful resource which provides drop information for monsters. This is backed by a Postgres database, but is seeded by a JSON file. Currently, this is based on GMS v83 drop data provided by HeavenMS.
## Environment
- JAEGER_HOST - Jaeger [host]:[port]
- JSON_FILE_PATH - File path to JSON data
- LOG_LEVEL - Logging level - Panic / Fatal / Error / Warn / Info / Debug / Trace
- DB_USER - Postgres user name
- DB_PASSWORD - Postgres user password
- DB_HOST - Postgres Database host
- DB_PORT - Postgres Database port
- DB_NAME - Postgres Database name

## API

Retrieval all drops for a given monster.
- /api/dis/monsters/{monsterId}/drops