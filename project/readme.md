Docker image creation
command
    PS C:\RAJA-MY-Folder\Learning_Section\GoLang\Udemy\Go-Micro\project>docker-compose up -d
    
docker ps - running container id will get
docker-compose down - stop all
docker stop project-broker-service-1 - pass container id to stop service
docker compose up -d        - to start again


clear all image 
        docker system prune -a

image rebuild
     docker compose up -d --build

to check log 
  docker logs -f <container name>

docker start <container-id>     - to start specific



admin connect on mongodb through copmpass app

  URI to:   mongodb://admin:password@localhost:27018/?authSource=admin
