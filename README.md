# Monitoring-Docker-containers


## How to run the project

- Clone the repository
- Run the bash script `./start.sh` to start the project. This script will build the images and start the containers.
- Open the browser and go to `http://localhost:3000` to see the frontend.

## How to stop the project
- Run the bash script `./stop.sh` to stop the project. This script will stop the containers and remove them.

## How to use the project
- The frontend displays the data from the database in a table. The data is updated every 5 seconds.

## Services
- ### Backend
  - The backend service receives data via the "api/ping/add" endpoint, and then updates the data in the database. All pings from the database are sent to the "api/pings" endpoint.

- ### Frontend
  - The frontend service displays the data from the database in a table. The data is updated every 5 seconds. There is also a server that redirects the data.

- ### Database
  - The database service is a PostgreSQL database that stores the pings.

- ### Pinger
  - The pinger service pings all Docker containers every 30 seconds and sends the data.

