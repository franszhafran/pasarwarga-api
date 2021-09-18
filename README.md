# Pasar Warga Backend Test

## Migrate schema
1. Run ```docker-compose up -d```
2. Run ```go build``` to build executable
3. Run ```./pasarwarga-service-api migrate``` to run migration

## Deployment

1. Run ```docker-compose up -d```
2. Run ```go build``` to build executable
3. Run ```./pasarwarga-service-api web``` to run application, will be live on port :8080