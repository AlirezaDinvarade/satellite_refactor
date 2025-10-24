#!/bin/bash

set -e
set -u

function create_user_and_database() {
    local database=$1
    local user=$2
    local password=$3
    
    echo "Creating user and database '$database'"
    
    psql -U postgres --dbname "postgres" <<-EOFSQL
        CREATE USER $user WITH PASSWORD '$password';
        CREATE DATABASE $database;
        GRANT ALL PRIVILEGES ON DATABASE $database TO $user;
        
        -- Connect to the new database and grant privileges on the public schema
        \c $database;
        GRANT ALL PRIVILEGES ON SCHEMA public TO $user;
        GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO $user;
        GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO $user;
        
        -- Create PostGIS extension if it doesn't exist
        CREATE EXTENSION IF NOT EXISTS postgis;

        -- Create uuid-ossp extension for uuid_generate_v4()
        CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOFSQL
}

if [ -n "$POSTGRES_MULTIPLE_DATABASES" ]; then
    echo "Multiple database creation requested: $POSTGRES_MULTIPLE_DATABASES"
    
    IFS=',' read -ra DB_CONFIG <<< "$POSTGRES_MULTIPLE_DATABASES"
    
    for config in "${DB_CONFIG[@]}"; do
        IFS=':' read db user pass <<< "$config"
        
        create_user_and_database "$db" "$user" "$pass"
        
        echo "Database created: $db with user: $user and password: ***hidden***"
    done
fi