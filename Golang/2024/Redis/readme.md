```
docker run -d --name my-redis \
  -e REDIS_PASSWORD=myStrongPassword \
  -e REDIS_USERNAME=myUser \
  -e REDIS_DB=myDatabase \
  redis
```