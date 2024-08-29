# Boulder Trails API

## Overview

The Boulder Trails API provides information about various trails in Boulder, Colorado. Built using Go and the Gin web framework, this API allows you to query trails based on different criteria, such as biking, fishing, and picnic facilities.

## Features

- **Get Trail by ID**: Retrieve details of a specific trail by its ID.
- **Get Free Biking Trails**: List biking trails that are free of charge.
- **Get Difficult Biking Trails**: List the most difficult biking trails.
- **Get Free Fishing Trails**: List fishing trails that are free of charge.
- **Get Picnic Spots with Restrooms**: List picnic spots that have restrooms.

## Endpoints

### `GET /trails/:id`

**Description:** Retrieve details of a specific trail by its ID.

**Parameters:**
- `id` (path parameter): The ID of the trail to retrieve.

**Response:**
- Status Code: `200 OK`
- Body: JSON object containing details of the trail.


### `GET /freebiketrails`

**Description:** List biking trails that are free of charge.

**Response:**
- Status Code: `200 OK`
- Body: JSON array of trails where biking is allowed and the fee is "No".

**Example Response:**
```json
[
    {
        "ID": "1",
        "Name": "Trail A",
        "BikeTrail": "Yes",
        "Fishing": "No",
        "Difficuly": "Easy",
        "FEE": "No",
        "Picnic": "Yes",
        "Restrooms": "Yes",
        "Address": "123 Trail Road"
    }
]

GET /difficultbiketrails
Description: List the most difficult biking trails.

Response:

Status Code: 200 OK
Body: JSON array of the most difficult biking trails.
Example Response:

json
Copy code
[
    {
        "ID": "2",
        "Name": "Trail B",
        "BikeTrail": "Yes",
        "Fishing": "No",
        "Difficuly": "Most Difficult",
        "FEE": "Yes",
        "Picnic": "No",
        "Restrooms": "No",
        "Address": "456 Trail Avenue"
    }
]
GET /freefishingtrails
Description: List fishing trails that are free of charge.

Response:

Status Code: 200 OK
Body: JSON array of trails where fishing is allowed and the fee is "No".
Example Response:

json
Copy code
[
    {
        "ID": "3",
        "Name": "Trail C",
        "BikeTrail": "No",
        "Fishing": "Yes",
        "Difficuly": "Moderate",
        "FEE": "No",
        "Picnic": "Yes",
        "Restrooms": "Yes",
        "Address": "789 Trail Boulevard"
    }
]
GET /picnicwithrestroom
Description: List picnic spots that have restrooms.

Response:

Status Code: 200 OK
Body: JSON array of trails with picnic facilities and restrooms.
Example Response:

json
Copy code
[
    {
        "ID": "1",
        "Name": "Trail A",
        "BikeTrail": "Yes",
        "Fishing": "No",
        "Difficuly": "Easy",
        "FEE": "No",
        "Picnic": "Yes",
        "Restrooms": "Yes",
        "Address": "123 Trail Road"
    }
]
Setup
Prerequisites
Go 1.16 or higher
Gin framework

Clone the Repository : https://github.com/ayushd96/Boulder-hills-project.git


Ensure the CSV file named BoulderTrailHeads.csv is placed in the root directory of the project.

Navigate to the main directory and start the server:
The server will start on http://localhost:8080.

Testing
Unit tests are yet to be written. 