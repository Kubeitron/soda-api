# SODA API
TBD

# Enviroment Variables
- `PORT` used to determine the port bound to by the API at startup
- `MONGO_INITDB_ROOT_USERNAME` root user for mongodb
- `MONGO_INITDB_ROOT_PASSWORD` root password for mongodb
- `MONGO_INITDB_DATABASE` initial database
- `MONGO_CLUSTER_ADDRESS` address to connect to

# How to develop
- Build and bring up the local dev container: `docker-compose up -d api`
- Attach to the container using VScode's extension: [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
- Develop!

- `Note` you need to re-build the image if you need to install any go modules via `go get`

# Misc
- For the OpenAPI spec, refer to http://localhost:5000/swagger/index.html while running the API
- To update the swagger spec, run `swag i cmd/api/` from the project root
