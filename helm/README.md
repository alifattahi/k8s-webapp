# how to use

## deploy
```bash
cd k8s-webapp
helm install <release-name> . --namespace <namespace>
#helm install webapp-release-1 . --namespace webapp --create-namespace
#helm install webapp-release-1 . --namespace webapp

```

## upgrade
```bash
helm upgrade <release-name> . --namespace <namespace>
helm upgrade --install webapp-release-1 . --namespace webapp --create-namespace
#
```