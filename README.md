# automation project based on Terraform, github actions, docker and aws 
This project leverages automation for building docker images , pushing it to dockerHub ,running integration test, creating aws infrastructure and deploy kubernetes infra (which is in progress )


### technologies and services used in this project

- GitHub Actions - CI/CD pipeline for deployment automation.
- Docker - Application containerization.
- Docker Compose - Orchestration of containers in local development environments.


### Running Project Locally

This project runs using Docker. 

```sh
docker compose  up 
```

 Browse the project at [http://127.0.0.1:3000]

### Running Tests Locally


```sh
docker compose -f docker-compose.testing.yml up --abort-on-container-exit --remove-orphans

docker compose -f docker-compose.testing.yml down --volumes
```

and create manually S3 bucket for tf-state holding and dynamodb table for tf-state lock 
# Running CI/CD

You need to assign github variables

### GitHub Actions Variables

This section lists the GitHub Actions variables which need to be configured on the GitHub project.

Variables:
- `DOCKERHUB_USER`: Username for [Docker Hub](https://hub.docker.com/) for avoiding Docker Pull rate limit issues.
- `PORT` golang api port 
- `FRONT_URL` frontend container name 
- `BASEURL` api contanier name 
- `DB_HOST` psql container name 
- `DB_PORT` psql port
- `DB_USER` psql username 
- `DB_DBNAME` psql database name 
- `AUTH0_AUDIENCE` your auth0 audience 
- `AUTH0_CLIENT_ID`your auth0 client Id
- `AUTH0_EMAIL` your auth0 email 
- `AUTH0_DIET_ENDPOINT` your auth0 endpoint  
- `AUTH0_DOMAIN` your auth0 domain

Secrets:


- `DOCKERHUB_TOKEN`: Token created in 
- `AUTH0_PASSWORD` your auth0 password  
- `DB_PASSWORD` psql password


# To trigger the CI/CD process, you simply need to merge a Pull Request into the main or prod branch, or commit directly to main/prod.


#### Section Notes and Resources
This project contains golang api and nuxt3 frontend borrowed from my friend  (https://github.com/1ChaLLengeR1/diet_project_frontend_nuxt3).
