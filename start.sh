#!bin/sh


# exit in case of error
set -e

echo "migarte up command"
/app/migrate -path app/migration -database "$DB_URL" --verbose up 

echo "start the app"
# execute any command passed -> CMD ["/app/main"] for entrypoint
exec "$@"