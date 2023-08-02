FROM --platform=linux/amd64 golang:1.19.3-alpine3.17 

# Define current working directory
WORKDIR /app

# Download modules to local cache so we can skip re-
# downloading on consecutive docker build commands
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the entire project directory into the container
COPY . .

RUN go build -o /meal-planner

EXPOSE 8080

CMD ["/meal-planner"]