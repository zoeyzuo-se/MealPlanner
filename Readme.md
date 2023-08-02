## Azure resources created
- Azure SQL Server koopa-mealplanner
- Azure SQL Database mealplanner
- Web app for containers fantastic-meal-planner
- Azure Container Registry mealplanner.azurecr.io

### Additional set up for Azure resources
- For Azure SQL Server koopa-mealplanner, create a firewall rule to allow access from your IP address. 

- Also check the box allow azure service to access this server. This can be done in the SQL Server → Networking → Exceptions → Check allow azure services and resources to access this server

## Build and Run

Run without docker
```
go run main.go
```
It will start the app on port 8080. Go to http://localhost:8080/ to see the app.

Run with docker
```bash
docker build --tag mealplanner.azurecr.io/api:local .
docker run -p 8080:8080 mealplanner.azurecr.io/api:local
```

## Deploy to Azure
```bash
az account set --subscription "Visual Studio Enterprise Subscription"
az acr login
docker push mealplanner.azurecr.io/api:local
```