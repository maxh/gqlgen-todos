nis_name=gqlgen_todos
dev_db_name=gqlgen_todos_dev
test_db_name=gqlgen_todos_test

conn="postgresql://postgres:postgres@localhost:5432"
from="FROM information_schema.tables"
where="WHERE table_schema='public'"

run_for_all_tables () {
  db_name=$1
  statements=`psql $conn/$db_name -c "SELECT $2 $from $where;" | grep ";"`
  echo $statements
  psql $conn/$db_name -c "$statements"
}

# Main functions.

delete_rows_for_db () {
  db_name=$1
  echo "Deleting all rows from all tables in $db_name database..."
  # We have to disable triggers to avoid foreign key constraints.
  run_for_all_tables $db_name "'ALTER TABLE ' || table_name || ' DISABLE TRIGGER ALL;'"
  run_for_all_tables $db_name "'TRUNCATE ' || table_name || ' CASCADE;'"
  run_for_all_tables $db_name "'ALTER TABLE ' || table_name || ' ENABLE TRIGGER ALL;'"
}

ensure_db_live () {
  db_name=$1
  output=$(PGPASSWORD=postgres psql -h localhost -U postgres -d $db_name -c '\echo' 2>&1)
  if [[ $output == *"closed the connection unexpectedly"* ]]; then
    echo "It appears postgresql server is still starting up. Try again in a few seconds..."
    exit 1
  fi
  if [[ $output == *"does not exist"* ]]; then
    echo "The database ${db_name} doesn't exist - creating it..."
    PGPASSWORD=postgres createdb -h localhost --username=postgres $db_name
    echo "Created."
  fi
}

# Run conditions.

ensure_db_live $dev_db_name
ensure_db_live $test_db_name

if [[ $1 == 'delete_tables' ]]; then
  delete_tables_for_db $dev_db_name
  echo
  delete_tables_for_db $test_db_name
fi

if [[ $1 == 'd' ]] || [[ $1 == 'delete' ]] || [[ $1 == 'delete_rows' ]]; then
  delete_rows_for_db $dev_db_name
  echo
  delete_rows_for_db $test_db_name
fi



