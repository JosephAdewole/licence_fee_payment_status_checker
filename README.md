# licence_fee_payment_status_checker
This part of a packing lot management software, It is an API (application Program Interface) that exposes endpoints allowing users(pakcing lot officers) to add new subscribers (cars), identiy and flag defaulters (non-subscribers using the parking lot).  
  
# Endpoints
Get all subscribers (both paid and defaulter) in a parking lot.
```json
method - GET
url - https://mawakif.herokuapp.com/api/subscribers

an example of response returned 
{
    "status": true,
    "code": 200,
    "data": [
        {
            "id": 184,
            "plate_number": "67P-78PL",
            "start_time": "2020-03-09T22:18:27Z",
            "status": false
        },
        {
            "id": 594,
            "plate_number": "67P-908PL",
            "start_time": "2021-05-17T18:18:27Z",
            "status": false
        }
    ]
}
```

| field                 | description                                  |
|-----------------------|----------------------------------------------|
| status                | returns true if successful                   |
|code                   | standard http status code                    |
|data                   | this contains the data/resource requested for|

# licence_fee_payment_status_checker
2
​
3
    
4
URL = https://mawakif.herokuapp.com
5
   
6
POSTMAN DOCUMENTATION = https://www.getpostman.com/collections/133f448fed78751ce2ff
7
  ```
  1. Endpoint for the Pi robot
- send the car plate number, packing space id, current time.
- returns status ok

/api/plate
POST 
{
    plate_number: "45-TY-90",
    packing_space_id: 5674, 
    current_time: 23:00:45
    is_empty: true/false
}

returns 
STATUS -  200 OK
{}


2. Endpoint for employees
- sends car plate number, status, current time.
-  returns status created if non exists, or status ok if it's updates a vehicle.

/api/subscribers/add
POST
{
    plate_number: "45-GH-46L",
    status: true,
    current_time: 21:50:59
}

status - 200/201
{
    message: "user registered"/ "user registration updated"
}

3. Endpoint for web interface 
- request for list of cars 
- returns a list of cars recorded by the employees and robot and their status

/api/subscribers
GET

returns 
[
    {
        plate_number: "34-534L-34"
        status: true/false,    (indicates if car is paid or unpaid)
        packing_space_id: 4233, 
        message: "upaid"/"paid"/"expired"
    },
    ...
]


4. Endpoint for admin
-  sets expiration duration 
/api/admin/ticket-duration
PUT
{
    ticket-duration: 24:00:00
}

returns 
STATUS 200 OK if successful.

Explanation:

When an employee records a paid user(car owner), the current time is designated as start time and status is 
true. If a car is detected by the robot and it's status is false, it does not exist in records or presence exceeds ticket duration (reset),  
the car is said to be unpaid thus this is recorded  or updated in database with a default start time of zero and a false status.

5. An endpoint that returns Packing space Checks
GET 
/api/checks

returns 
STATUS 200 OK
[
    {
        packing_space_id: 10,
        plate_number: "34L-24K-43",
        is_empty: true/false,
        created_at: "23:45:89"

    }
]

6. An endpoint that add and another that returns a list of packing space and their designation

GET

/api/space 

returns
[
    {
        id: 12,
        designation: "The President"
    }, 
    ...
]

POST

/api/space 

returns
[
    {
        id: 12,
        designation: "The President"
    }, 
    ...
]

///////
Create another database that logs events of image checks... (even if no plate number is detected)



TABLE 1  - subscribers
id, plate_number, packing_space_id, start_time, status

TABLE 2 - checks
id, is_empty, plate_number, packing_space_id, created_at

TABLE 3 -  packing_space
id, designation

TABLE 4 - config
id, name, value
 

note 3: has records of the different packing space that exists
  ```
8
​
9
​
10
        



    
URL = https://mawakif.herokuapp.com   
   
POSTMAN DOCUMENTATION = https://www.getpostman.com/collections/133f448fed78751ce2ff



        
