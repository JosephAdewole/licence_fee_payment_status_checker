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
    
8
​
9
​
10
        



    
URL = https://mawakif.herokuapp.com   
   
POSTMAN DOCUMENTATION = https://www.getpostman.com/collections/133f448fed78751ce2ff



        
