#!/bin/bash
set -e

# Connect to the default database and install uuid-ossp extension
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    DO \$\$
    BEGIN
      IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'asad') THEN
        CREATE USER asad;
      END IF;
    END \$\$;
EOSQL

# Check if the product_svc database exists
if ! psql -lqt --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" | cut -d \| -f 1 | grep -wq "product_svc"; then
  psql --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -c "CREATE DATABASE product_svc;"
  psql --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -c "GRANT ALL PRIVILEGES ON DATABASE product_svc TO asad;"
fi

# Install uuid-ossp extension in product_svc database
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "product_svc" <<-EOSQL
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL

# Check if the auth_svc database exists
if ! psql -lqt --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" | cut -d \| -f 1 | grep -wq "auth_svc"; then
  psql --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -c "CREATE DATABASE auth_svc;"
  psql --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -c "GRANT ALL PRIVILEGES ON DATABASE auth_svc TO asad;"
fi

# Install uuid-ossp extension in auth_svc database
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "auth_svc" <<-EOSQL
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL

