# **Buff Backend api**

## **When working with money in this project, everything is saved as cents, so the backend always multiplies the dollars by 100 cents and saves it in the database as BIGINT and then divides it by 100 before display**

This server runs in a docker container, please make sure you have docker and docker-compose installed visit [https://docs.docker.com/compose/install/](https://link) (for docker compose) and [https://docs.docker.com/get-docker/](https://link) (for docker)

## **Create a .env(for production) and a .env.dev(for development) and a .env.test(for testing) with the variables from sampleEnv before starting the app**

After installation of docker and docker compose, simply run

```bash
docker build --target development -t bufftv/user-development:latest ./userService/
docker build --target development -t bufftv/admin-development:latest ./adminService/
docker-compose up
```

run

```bash
docker exec -it auth_service sh
```

to access the container cli to run scripts like migrations etc.
migrations: `npm run db-migrate` to run migrations. make sure you run `docker exec -it auth_service sh` first.

in the root directory of this project

To view documentation for user routes visit:
`
http://localhost:3000/user-api-docs
`
