# Meal Planner
Meal planner is a service that allows you to store your recipes and plan your meals for the week.

The API can be accessed at https://fantastic-meal-planner.azurewebsites.net/

## Features
- [x] Add recipe
- [x] View recipes
- [x] Update recipe by name
- [x] Delete recipe by name
- [x] Get recipe by name
- [ ] Add recipes to meal plan
- [ ] View meal plan for the week
- [ ] Add ingredients to recipe
- [ ] Update ingredients in recipe
- [ ] Delete ingredients from recipe
- [ ] Making suggestion on what to cook based on ingredients - OpenAI
- [ ] Making suggestion on what to cook based on recipes - OpenAI
- [ ] Add ingredients to shopping list
- [ ] View shopping list
- [ ] Update shopping list
- [ ] View ingredients


## Tech Stack
- Go
- Azure SQL
- Azure Web App for Containers
- Azure Container Registry
- Docker


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