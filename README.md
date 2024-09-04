## Docs
here is a simple web application to show multistage building and running in kubernetes with liveness and readiness probes


### Endpoints
```
# show hostname of container
/ 

# get readiness status
/ready

# get liveness status
/healthy
```

### local test only
```bash
docker compose up -d
curl -v localhost:8080/
curl -v localhost:8080/ready
curl -v localhost:8080/healthy
```

### docker image
```bash
docker pull alifattahi/k8s-webapp:latest
```
####
Ali Fattahi