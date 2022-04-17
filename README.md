## Description

App is built using [fiber](https://github.com/gofiber/fiber) and [gorm](https://gorm.io/).

Implemented:
- Overview of all projects
- Creating new projects
- Updating projects

## Build
```docker build -f ./deploy/Dockerfile_app -t project_app .```

## Run 
```docker-compose -f ./deploy/docker-compose.yml up -d```

By default app is host on `8080` port, you can change it in `deploy/docker-compose.yml`.

## Api

### Get projects
```
GET /api/project
```

**Response:**
```javascript
[
    {
        "ID": 1,
        "Name": "project",
        "State": "planned",
        "Progress": 0,
        "Owner": {
            "Id": "<employee_id>",
            "FirstName": "Name",
            "LastName": "Last name",
            "Department": "marketing"
        },
        "Participants": [
            {
                "Id": "<employee_id>",
                "FirstName": "Name",
                "LastName": "Last name",
            },
            {
                "Id": "<employee_id>",
                "FirstName": "Name",
                "LastName": "Last name",
            }
        ]
    }
]
```

### Create project
```
POST /api/project
```

**Request:**

```javascript
{
    "Name": "project_0",
    "State": "planned",
    "Progress": 0,
    "Owner": "<employee_id>",
    "Participants": [
        "<employee_id>",
        "<employee_id>"
    ]
}
```

### Update project
```
POST /api/project/:id
```

**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
id | INT | YES | Positive int.

**Request:**

```javascript
{
    "Name": "project_0",
    "State": "planned",
    "Progress": 0
}
```
