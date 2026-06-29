Docker image creation
command
    PS C:\RAJA-MY-Folder\Learning_Section\GoLang\Udemy\Go-Micro\project>docker-compose up -d
    
docker ps - running container id will get
docker-compose down - stop all
docker stop project-broker-service-1 - pass container id to stop service
docker compose up -d        - to start again

image rebuild
     docker compose up -d --build

docker start <container-id>     - to start specific