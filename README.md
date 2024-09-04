## Docs
here is a simple web application to show multistage building and running in kubernetes with liveness and readiness probes

### Port
```
3000
```
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
curl -v localhost:3000/
curl -v localhost:3000/ready
curl -v localhost:3000/healthy
```

####
Ali Fattahi