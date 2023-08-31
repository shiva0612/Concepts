docker build -t shiva-golang-docker --target production:1.0 .
docker run --name shiva-go shiva-golang-docker:1.0

docker tag shiva-golang-docker:1.0 docker.io/shiva0612/shiva-golang-docker:1.0
docker push docker.io/shiva0612/shiva-golang-docker:1.0

docker rmi shiva-golang-docker:1.0
docker run --name test docker.io/shiva0612/shiva-golang-docker:1.0