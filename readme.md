## Youtube API Assignment

## To run the server 
* All Dependencies are already included
* Simply `go run main.go`

## Prerequisite for running task
* Create a .env file
* Setting Up Env File
```javascript
API_KEYS =
MONGODB_URL =
DB_NAME =
COLLECTION_NAME = 
```

## Task Details
1. GoRoutines is added so that after every 10 minutes, youtube fetches videos related to `cricket` and push it in mongodb.
2. Get Request API - `https://localhost:3000/api/getVideos`
**Response** - Added in file `mockResponse.json`

