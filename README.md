# About the project
this is a simple rest api to create,update,delete and fetch financial data from mongoDb. 
also there is a code to read the csv_file and enter its data to mongodb using workers in concurrent manner.

# How to run project
- install docker and docker-compose 
- clone the project
- run by `docker-compose up -d` which would starts the http server on port 8000
- it takes some time to populate data to database, and after that the list api
  would return data
- for running test you can user `docker-compose exec api go test ./...` 

# APIs
for testing API use postman collection provided in project.

# Structure
- the cmd directory contains codes that could be compiled to executable binaries 
- business logic is stored in internal directory
- financial package is responsible for crud related to financial data

# Points
- tests are written in financial and queue package.
- we use real database for intergration tests

# Extra libraries used
- gin for routing
- viper for configs
- zap for logging
- testify for tests


