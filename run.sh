if docker images bookserver | grep -q bookserver; then
    docker ps -a --filter "ancestor=bookserver" -q | xargs -r docker rm -f
    docker rmi bookserver
else
    echo "Image 'bookserver' does not exist."
fi  
go build -o BookServer ./main.go
docker build -f deployments/dockerfile -t bookserver .
docker run -p 8080:8080 bookserver init -z 8080