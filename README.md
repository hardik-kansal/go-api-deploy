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

- [X] Implement RESTful HTTP API in Go using Gin
- [X] Load config from file & environment variables in Go with Viper
- [ ] Mock DB for testing HTTP API in Go and achieve 100% coverage
- [X] Implement transfer money API with a custom params validator
- [X] Add users table with unique & foreign key constraints in PostgreSQL
- [ ] How to handle DB errors in Golang correctly
- [X] How to securely store passwords? Hash password in Go with Bcrypt!
- [ ] How to write stronger unit tests with a custom gomock matcher
- [X] Why PASETO is better than JWT for token-based authentication?
- [X] How to create and verify JWT & PASETO token in Golang
- [X] Implement login user API that returns PASETO or JWT access token in Go
- [X] Implement authentication middleware and authorization rules in Golang using Gin

## Section 3: Deploying the application to production [Kubernetes + AWS]

- [X] Build a minimal Golang Docker image with a multistage Dockerfile
- [X] How to use docker network to connect 2 stand-alone containers
- [X] How to write docker-compose file and control service start-up orders with wait-for.sh
- [X] How to create a free tier AWS account
- [X] Auto build & push docker image to AWS ECR with Github Actions
- [X] How to create a production DB on AWS RDS
- [X] Store & retrieve production secrets with AWS secrets manager
- [ ] Kubernetes architecture & How to create an EKS cluster on AWS
- [ ] How to use kubectl & k9s to connect to a kubernetes cluster on AWS EKS
- [ ] How to deploy a web app to Kubernetes cluster on AWS EKS
- [ ] Register a domain name & set up A-record using Route53
- [ ] How to use Ingress to route traffics to different services in Kubernetes
- [ ] Automatic issue TLS certificates in Kubernetes with Let's Encrypt
- [ ] Automatic deploy to Kubernetes with Github Action

#### Honing golang skills each day !



```markdown
# Simple Bank

This project sets up a simple banking application using Docker, PostgreSQL, and a Go backend. The following instructions will guide you through the setup process, including building and running the Docker containers and visualizing the database with TablePlus.

## Prerequisites

- Git
- Docker
- TablePlus (or any database visualization tool)
- Make (if you don't have Make installed, you can manually run the equivalent commands)
```

## Setup Instructions

### Step 1: Clone the Repository

```bash
git clone git@github.com:hardik-kansal/go-api-deploy.git
cd go-api-deploy
```

### Step 2: Create a Docker Network

```bash
docker network create bank-network
```

### Step 3: Set Up PostgreSQL

```bash
make postgres
make createdb
make migrateup
```

### Step 4: Prepare the Dockerfile

Overwrite the `example.dockerfile` to `Dockerfile`.

```bash
mv example.dockerfile Dockerfile
```

### Step 5: Build the Docker Image

```bash
docker build -t simplebank:latest .
```

### Step 6: Run the Docker Container

```bash
docker run --name simplebank -e GIN_MODE=release --network bank-network -p 8000:8000 -e DB_URL="postgresql://root:123@pg16:5432/simplebank?sslmode=disable" simplebank
```

### Step 7: Visualize the Database with TablePlus

1. Open TablePlus.
2. Create a new connection:
   - **DB Name**: simplebank
   - **User**: root
   - **Password**: root
   - **Host**: localhost
   - **Port**: 5432

### Step 8: Send Requests Using Postman

Send requests to your server at `localhost:8000`.

## Notes

- Ensure that PostgreSQL is running on port 5432.
- The server should be accessible on port 8000.
- You can use TablePlus or any other database visualization tool to inspect and manage your database.
- Use Postman to send HTTP requests to your backend service at `http://localhost:8000`.

[Backend Course Link](https://www.youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE)


