## Section 1: Working with Database [Postgres]

- [X] Design DB schema and generate SQL code with dbdiagram.io 
- [X] Install & use Docker + Postgres + TablePlus to create DB schema
- [X] How to write & run database migration in Golang
- [X] Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc
- [X] Write unit tests for database CRUD with random data in Golang
- [X] A clean way to implement database transaction in Golang
- [X] DB transaction lock & How to handle deadlock in Golang
- [X] How to avoid deadlock in DB transaction? Queries order matters!
- [X] Deeply understand transaction isolation levels & read phenomena in MySQL & PostgreSQL
- [X] Setup Github Actions for Golang + Postgres to run automated tests

## Section 2: Building RESTful HTTP JSON API [Gin]

- [ ] Implement RESTful HTTP API in Go using Gin
- [ ] Load config from file & environment variables in Go with Viper
- [ ] Mock DB for testing HTTP API in Go and achieve 100% coverage
- [ ] Implement transfer money API with a custom params validator
- [ ] Add users table with unique & foreign key constraints in PostgreSQL
- [ ] How to handle DB errors in Golang correctly
- [ ] How to securely store passwords? Hash password in Go with Bcrypt!
- [ ] How to write stronger unit tests with a custom gomock matcher
- [X] Why PASETO is better than JWT for token-based authentication?
- [ ] How to create and verify JWT & PASETO token in Golang
- [ ] Implement login user API that returns PASETO or JWT access token in Go
- [ ] Implement authentication middleware and authorization rules in Golang using Gin

## Section 3: Deploying the application to production [Kubernetes + AWS]

- [X] Build a minimal Golang Docker image with a multistage Dockerfile
- [ ] How to use docker network to connect 2 stand-alone containers
- [ ] How to write docker-compose file and control service start-up orders with wait-for.sh
- [X] How to create a free tier AWS account
- [X] Auto build & push docker image to AWS ECR with Github Actions
- [X] How to create a production DB on AWS RDS
- [ ] Store & retrieve production secrets with AWS secrets manager
- [ ] Kubernetes architecture & How to create an EKS cluster on AWS
- [ ] How to use kubectl & k9s to connect to a kubernetes cluster on AWS EKS
- [ ] How to deploy a web app to Kubernetes cluster on AWS EKS
- [ ] Register a domain name & set up A-record using Route53
- [ ] How to use Ingress to route traffics to different services in Kubernetes
- [ ] Automatic issue TLS certificates in Kubernetes with Let's Encrypt
- [ ] Automatic deploy to Kubernetes with Github Action

#### Honing golang skills each day !
[Backend Course Link](https://www.youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE)


`docker run --name simplebank -e GIN_MODE=release --network bank-network -p 8000:8000 -e DB_URL="postgresql://root:123@pg16:5432/simplebank?sslmode=disable" simplebank`
