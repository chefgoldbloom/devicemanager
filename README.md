# Device Manager
### Device Manager is meant to be an extended version of Axis Device manager with less limitations and more portability.
Planning on Backend in Golang, Frontend in Python and DB in postgres

## Endpoints
|     Method     |    URL Pattern      |            Handler           |                   Action                |
| -------------- | ------------------- | ---------------------------- | --------------------------------------- |
|      GET       |   v1/healthcheck    |      healthcheckHandler      | show application information            |
|      GET       |   v1/cameras        |      listCamerasHandler      | show details of cameras                 |
|      POST      |  /v1/cameras        |      createCameraHandler     | create a new camera                     |
|      GET       |  /v1/cameras/:id    |      showCameraHandler       | show the details of a specific camera   |
|      PUT       |  /v1/cameras/:id    |      editCameraHandler       | update the details of a specific camera |
|     DELETE     |  /v1/cameras/:id    |      deleteCameraHandler     | delete a specific camera                |


## Scanner
The ability to scan a range of IPs to create cameras is a nice to have.


## Info From Cameras
 - MAC Address
 - IP Address
 - Model
 - Firmware
 - Site
 - Camera Name
