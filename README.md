# Microservices-Calc

It's a web-calculator app based on microservice architecture. (Project creadted for educational purposes)

## Microservice communication pattern

```
           addsub
            /
web -- gateway
            \
           muldiv
```

## Features:
- addition and subtraction served by golang service
- multiplication and divisoon served by python service
- ready to build docker images
- ready to run in docker-compose mode
