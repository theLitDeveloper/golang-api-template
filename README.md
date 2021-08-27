# GOLANG API TEMPLATE
A starter for a faster project setup and predictable results when creating an RESTful API with Go and a RDBMS. It provides support for MySQL/MariaDB and Postgres, and thus also for CockroachDB or Amazon Aurora.  

It also provides endpoints for monitoring and version info:
- **/version** Version info  
- **/metrics** Prometheus  
- **/health** Health check

## What's inside?
- **Echo Framework** [Learn more](https://echo.labstack.com/)
- **GORM** [Learn more](https://gorm.io/index.html)
- **zap** [Learn more](https://github.com/uber-go/zap)  

## Development
1. MySQL/MariaDB is enabled by default. If want to use Postgres, edit [pkg/service/datastore.go](pkg/service/datastore.go)
2. Set required env vars
```
export VERSION=$(git describe --tags --abbrev=0)

export DB_PASS=<db user password>
export DB_USER=<db user name>
export DB_NAME=<db name>
export DB_HOST=<host>
export DB_PORT=<port>
```
and replace all placeholder values with yours.  
