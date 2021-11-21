# API Documentation

 All responses come in standard JSON. All requests must include a `content-type` of `application/json` and the body must be valid JSON.


### Response Codes
```
200: Success
201: Created
400: Bad request
401: Unauthorized
404: Cannot be found
405: Method not allowed
50x: Server Error
```
### Error and Success Message Example

```json
  {
    "success":false,
    "error":"error" 
  }
  
  {
      "success": true, 
      "message": "message",  //w.r.t the API
      "data" : {
          "": ""
      }   //w.r.t the API
  }
```

## SignUp

**You send:**  You send the details required to signup.

**You get:** An `Error-Message` or a `Success-Message` depending on the status of the account created.

**Endpoint:** 
     /signup

**Authorization Token:** Not required

**Request:**
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{   
    "email":"yashi.gupta@gmail.com",
    "password":"qwerty",
    "type":1
    
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "Signup SuccessFul"
}
```

## Login

**You send:**  Your  login credentials.

**You get:** An `API-Token` and a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /login

**Authorization Token:** Not required

**Request:**
`POST HTTP/1.1`
```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "email":"yashi.gupta@gmail.com",
    "password":"qwerty"
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "Log In successful",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inlhc2hpLmd1cHRhQGdtYWlsLmNvbSIsInR5cGUiOjF9.w1qyPVtDE2bgavD-Q3rU-BjQrVDbw2AWodMimD3JJvo"
}
```
## Google Login

**You send:**  The Response fron the API /signup/google along with the 'type' of the user 

**You get:** An `API-Token` and a `Success-Message` with which you can make further actions.

**Endpoint:** 
First use /signup/google to get the response

Output-`{ "Name":"Yashi Gupta","Email":"vashishtiv@gmail.com" }`

type-> Student- 1
    -> Teacher- 2

then send this respose with the type to /login/google

**Authorization Token:** Not required

**Request:**
`POST HTTP/1.1`
```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
     "name":"Yashi Gupta",
     "email":"vashishtiv@gmail.com",
     "type":1
}
```


**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "Log In successful",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inlhc2hpLmd1cHRhQGdtYWlsLmNvbSIsInR5cGUiOjF9.w1qyPVtDE2bgavD-Q3rU-BjQrVDbw2AWodMimD3JJvo"
}
```

## Signup by google

**You send:**  You send the request to /signup/google and then with the recieved response send it to signup api.

**You get:** An `Error-Message` or a `Success mesage`

**Endpoint:** 
     /signup/google -> /signup

**Authorization Token:** No Token required

**Request: to signup**
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{   
    "name":"Yashi Gupta",
    "email":"vashishtiv@gmail.com",
    "type":1
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: image/png
Content-Length: xy
{
    "success": true,
    "message": "Signup SuccessFul"
}
```

## Create Class

**You send:**  the name of the subject.

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/add

**Authorization Token:** Teacher's auth token required

**Request:**
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{   
   "subject":"----"
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "Class created Successfully",
    "data": "000355shi" ----->this is the class code
}
```

## Join Class

**You send:**  the code of the class.

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/add

**Authorization Token:** Student's auth token required

**Request:**
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "class_code": "73fbbf"
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message":"------------"
}
```

## Get All Classes of a user

**You send:**  nothing

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/get

**Authorization Token:** user's auth token required


**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message":"------------",
    "data" : [{

    }]
    
}
```
## Create Anouncement

**You send:**  The details of the announcement.

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/announcement/add

**Authorization Token:** Teacher's auth token required

**Request:**
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "teacher_name":"yashi",
    "class_code":"e37e5dshi",
    "description":"uihnkiuygvbxjh",
    "timestamp":"12-Nov-21"
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "Announcement added Successfully"
}
```

## Get Anouncement

**You send:**  The code of the class to fetch the respective announcements.

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/announcement/get

**Authorization Token:** user's auth token required

**Request:**
`GET HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "class":"e37e5dshi"
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "class data fetch successful",
    "data": [
        {
            "_id": "619a1734c07b2ce54d45f227",
            "classcode": "e37e5dshi",
            "description": "uihnkiuygvbxjh",
            "id": 3009,
            "teachername": "",
            "timestamp": "12-Nov-21"
        },
        {
            "_id": "619a1bbdbeea287a6a8e623b",
            "classcode": "e37e5dshi",
            "description": "uihnkiuygvbxjh",
            "id": 28033,
            "teachername": "",
            "timestamp": "12-Nov-21"
        }
    ]
}
```
## Create Assignment

**You send:**  The details required for a pdf- assignment in the format of a form

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/assignment/add

**Authorization Token:** Teacher's auth token required

**Request:**

Form Fields

`"description"`

`"date"`

`"class_code"`

`"file"`

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "Assignment Added"
}
```
## Get Assignment

**You send:**  The code of the class to fetch the respective assignments.

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/assignment/get

**Authorization Token:** user's auth token required

**Request:**
`GET HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "class":"e37e5dshi"
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "class data fetch successful",
    "data": [
        {
            "_id": "619a3f274d41e6558ee29b29",
            "classcode": "e37e5dshi",
            "date": "0001-01-01T05:30:00+05:30",
            "description": "dfcdsfgvbdxhcju",
            "file": {
                "filename": "https://storage.googleapis.com/batbuck/Registration Successful - Registration of Migrants and Others for Travelling to Uttarakhand.pdf",
                "submissions": null
            },
            "form": {
                "answers": null,
                "ques": null,
                "soln": null
            },
            "id": 65448,
            "name": "",
            "type": 0
        },
        {
            "_id": "619a44fb4d41e6558ee29b2a",
            "classcode": "e37e5dshi",
            "date": "0001-01-01T05:30:00+05:30",
            "description": "grftesdxfcgvh",
            "file": {
                "filename": "https://storage.googleapis.com/batbuck/Registration Successful - Registration of Migrants and Others for Travelling to Uttarakhand.pdf",
                "submissions": null
            },
            "form": {
                "answers": null,
                "ques": null,
                "soln": null
            },
            "id": 35749,
            "name": "",
            "type": 0
        }
    ]
}
```

## Submit Assignment

**You send:** you upload the file and other details in the form 

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/assignment/sub

**Authorization Token:** Student's auth token required

**Request:**
`POST HTTP/1.1`

#### Form

`"id"`   ->the id of the assignment the user is submitting

`"file"` ->upload file

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "Assignment Submitted"
}
```
## get submissions of an assignment

**You send:**  The id of the assignment like this -> class/assignment/sub/`61283`

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/assignment/sub/{code}

**Authorization Token:** user's auth token required

**Request:**
`GET HTTP/1.1`

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "successful",
    "data": [
        {
            "filename": "https://storage.googleapis.com/batbuck/okay.txt.txt",
            "email": "vashishtiv@gmail.com",
            "timestamp": "2021-11-21 T 19:14"
        }
    ]
}
```

## Create MCQ Assignment

**You send:**  the questions and options along with class code

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/mcq/add

**Authorization Token:** user's auth token required

**Request:**
`GET HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "mcq":{
        "ques":[
            {
                "question":"usdhgbnchjv",
                "options":["a","b","c"]
            },
            {
                "question":"usdhgbnchjv",
                "options":["a","b","c"]
            }
        ],
        "answers":["a","c"]
    },
    "code":"e37e5dshi",
    "description":"kjuyhtgfvbnj",
    "date":"12-Dec-21" 
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "Assignment Added"
}
```
## submit MCQ Assignment

**You send:**  the id of the assignment and the selected answers

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/mcq/sub

**Authorization Token:** user's auth token required

**Request:**
`GET HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "id":2249,
    "ans":["a","c"]
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "Assignment Submitted"
}
```
## Environment variables

  
  `DbUrl` : to store the url of mongodb cluster or localhost.

  `secretkey` : to store encryption secret.

  `PORT`     : port on localhost to run application

 `SMTP_KEY` : to store the key of smtp server.

 `SMTP_PASS` : to store the pass of the smtp server.