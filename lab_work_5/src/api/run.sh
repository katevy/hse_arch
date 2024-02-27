docker build . -t api
docker run -i -t -p 8080:8080 api