# Gin
A RESTful API service template built with gin.

![Integration](https://github.com/DarkbordermanTemplate/gin/workflows/Integration/badge.svg)
![Build](https://github.com/DarkbordermanTemplate/gin/workflows/Build/badge.svg)
![Release](https://github.com/DarkbordermanTemplate/gin/workflows/Release/badge.svg)

## Development

### Prerequisite

| Name | Version |
| --- | --- |
| GNU make | 4.2.1 |
| go | 1.16 |

### Environment setup

0. Initialize environment variable

```
cp sample.env .env
```

1. Apply environment variables

```
make shell
```

2. Download required packages
```
make init
```

3. Start development API service
```
make run
```

4. (Optional) Create binary executable output
```
make build
```

### Formatting

This project use `gofmt` for formatting
```
make format
```

### Testing

This project use `go` for testing
```
make test
```

## Deployment

### Prerequisite

| Name | Version |
| --- | --- |
| Docker | 19.03.6 |
| docker-compose | 1.17.1 |

### Building image

This will build the image with tag `gin:latest`
```
docker-compose build
```

### Deployment

The service is deployed with `docker-compose`.

0. Start containers
```
docker-compose up
```

### Contribution
* Darkborderman/Divik(reastw1234@gmail.com)
