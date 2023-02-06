migrate-up() {
    user=root
    password=root
    host=127.0.0.1
    port=3306
    mysql --user=$user --password=$password --host=$host --port=$port -e "DROP DATABASE IF EXISTS vessel_management; CREATE DATABASE vessel_management; SET GLOBAL FOREIGN_KEY_CHECKS=0;"
    SERVICE_DB_USERNAME=$user SERVICE_DB_PASSWORD=$password SERVICE_DB_HOST=$host SERVICE_DB_PORT=$port SERVICE_DB_DATABASE=vessel_management make migrate-up
}
