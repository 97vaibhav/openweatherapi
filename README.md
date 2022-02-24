# openweatherapi

## To use this app you have to first create a file name .apiConfig after that put your api key into this file as shown below 
```
touch .apiConfig

### put the below text with your api key inside blank


{
  "ApiKey": "__"
}

```
### To get api key visit
- https://openweathermap.org/

### Run commands

```
go build 
./openweatherapi.exe

```

## Endpoints are 
- http://localhost:8080/welcome
- http://localhost:8080/weather/{city_name}
