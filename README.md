# OCLS

    Ocls is an online assignment submission portal built for schools.

    The project is live [here]()

    Video demo:

<a href="" title="OCLS Demo"><img src="{image-url}" alt="Alternate Text" /></a>

## The Presentation [link]()

## The Video Demo [link]()

### The Idea 
    During the tough time of COVID-19 the childrens were affected very much in terms of academics, they had no option other than online studies.

    But online studies were very dicrete, they recieved assignments on one platform and they had to submit it on the other, Even sometimes they missed there assignments.
    
    Though forms were a great way to take tests, there was a problem that was, they were not in an integrated system which sometimes caused chaos.

    Considering such needs this project is conceptualized.

### Fetures

     -> Google oauth login possible.
     -> Seperate Dashboard for Student and Teacher.
     -> Techer can create classes.
     -> Two types of assignment can be created by the teacher : File based , Multiple Choice Question
     -> Email notification sent to all the students, of the particular class whenever any assignment is created.
     -> Remainder email to students.
     -> Auto grading for MCQ assignment.
     -> Teacher can download, the excel sheet of submissions of any assignment.
     -> Students can join any class.
     -> Students can attempt the assignment within the portal.

### Future Amendments/scope

     -> Improve UI/UX of the portal.
     -> Video calling for classes.
     -> Lock the browser while submitting mcq assignment to prevent cheating.
     -> Add Coding assignment section.
     -> In window editor for submitting assignment.
     -> Use ML/AI to detect piracy in assignments.
     -> Navigation link directly from mail to dashboard

# SetUp Project:

### Download [Go](https://go.dev/)

### Clone the project
### Crete 'creds.json' file and add it into the api/auth folder.
  
  This file is required for Oauth signin using google.

 `Attention`: In case you don't want to use this feature, comment the code in the file name- signup_google.go 

#### Steps to create creds.json file:

   1. Setup Oauth Consent Screen ->  [Docs](https://support.google.com/cloud/answer/10311615?hl=en&ref_topic=3473162)

   2. Create Oauth Client ID -> [Docs](https://support.google.com/cloud/answer/6158849?hl=en&ref_topic=3473162)

   3. Download the `Json` file and rename it as `creds.json`.

   4. modify the creds file in the following format.

    ```json
    {
        "client_id":"",
        "project_id":"",
        "auth_uri":"",
        "token_uri":"",
        "auth_provider_x509_cert_url":"",
        "client_secret":"",
        "redirect_uris":[],
        "javascript_origins":[]
    }
    ``` 

#### Steps to create Creds2.json file (required to upload file in bucket)

   1. Use [this](https://cloud.google.com/docs/authentication/production) to authenticate the service account.

   2. Download the file.

   3. Place it into the utility folder.
### Setup the Environment variables :

  `DbUrl` : to store the url of mongodb cluster or localhost.

  `secretkey` : to store encryption secret.

  `PORT`     : port on localhost to run application

  `SMTP_KEY` : to store the key of smtp server.

  `SMTP_PASS` : to store the pass of the smtp server.

### Run the Project using commands :
    
    go run main.go 
    or
    go build && ./ocls (parent folder name)
    

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

type-> Student- 1
    -> Teacher- 2

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
    "email":"",
    "password":"",
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
    "email":"",
    "password":""
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
    "data": {
        "User": {
            "_id": "6199--------------5",
            "classcode": [
                "8eiuhbdfh"
            ],
            "email": "",
            "password": "dfcvdchhgdbncj",
            "type": 2
        },
        "Token": "eyJhbGciOiJIUzI1NiIsInRCJ9.eyJlbWFpbCI6Inlhc2hpLmd1cHRheWFzaGkuZ3VwdI6IiIsInR5cGUiOjJ9.P_gdD2AlFu0Cu_bi2HvnYCbB0vk0ScQhrds"
    }
}
```
## Google Login

**You send:**  The Response fron the API /signup/google along with the 'type' of the user 

**You get:** An `API-Token` and a `Success-Message` with which you can make further actions.

**Endpoint:** 
First use /signup/google to get the response

Output-`{
     "Name":"",
     "Email":"jkmklkol@gmail.com"
}`

then send this respose with the type to /login/google

**Authorization Token:** Not required

**Request:**
`POST HTTP/1.1`
```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
     "name":"",
     "email":"",
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
    "data":  {
        "User": {
            "_id": "6199--------------5",
            "classcode": [
                "8eiuhbdfh"
            ],
            "email": "",
            "password": "dfcvdchhgdbncj",
            "type": 2
        },
        "Token": "eyJhbGciOiJIUzI1NiIsInRCJ9.eyJlbWFpbCI6Inlhc2hpLmd1cHRheWFzaGkuZ3VwdI6IiIsInR5cGUiOjJ9.P_gdD2AlFu0Cu_bi2HvnYCbB0vk0ScQhrds"
    }
}
```

## Signup by google

**You send:**  You send the request to /signup/google and then with the recieved response send it to signup api.

**You get:** An `Error-Message` or a `Success mesage`

**Endpoint:** 
     /signup/google --> /signup

**Authorization Token:** No Token required

**Request: to signup**
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{   
    "name":"",
    "email":"",
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

**Request:**
`GET HTTP/1.1`

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
    "teacher_name":"",
    "classcode":"",
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
`POST HTTP/1.1`

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
            "_id": "619a1bbdbee6a8e623b",
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

`"description"-> Text`

`"date" -> Text`

`"class_code" -> Text`

`"file" -> file`

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
## Get Assignment for Teacher dashboard

**You send:**  The code of the class to fetch the respective assignments.

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/assignment/get

**Authorization Token:** user's auth token required

**Request:**
`POST HTTP/1.1`

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
                "filename": "",
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
                "filename": "",
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

## Get Assignment for Student dashboard

**You send:**  The code of the class to fetch the respective assignments.

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /class/assignment/stu-get

**Authorization Token:** user's auth token required

**Request:**
`POST HTTP/1.1`

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
                "filename": "",
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
                "filename": "",
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
     /class/assignment/sub/{id}

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
            "timestamp": "2021-11-21 19:14"
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
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "mcq":{
        "ques":[
            {
                "question":"usdhgbnchjv",
                "options":["iuythgbn","b hn","bgvhjn"]
            },
            {
                "question":"usdhgbnchjv",
                "options":["true","false"]
            }
        ],
        "answers":["iuythgbn","false"]
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
`POST HTTP/1.1`

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

## update user details

**You send:**  the modifide data of the user for eg: to add "adm no" and "phone" we will do the following:
  
Fields of Techer are:
    "name"        string    -> only name can be changed 
	"email"       string 
    "post"        []int
    "class"       []string
    "assignment"  []int
	"type"        int      

Fields of Student are:
    "name"      string  ->can be changed
    "password"  string
    "email"     string
	"admno"     string  ->can be changed
	"phone"     string  ->can be changed
	"classcode" []string
	"type" int

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions along with complete updated data of user.

**Endpoint:** 
     /user/update

**Authorization Token:** user's auth token required

**Request:**
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
            "classcode": [
                "e37e5dshi"
            ],
            "name":"",
            "admno":"6789",
            "phone":"2345678901",
            "email": "",
            "password": "b1b3773a05c0ed01767",
            "type": 2
 }
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success": true,
    "message": "user updated successfully",
    "data": {
        "_id": "6199ef576240688196921bf5",
        "admno": "6789",
        "classcode": [
            "e37e5dshi"
        ],
        "email": "",
        "name": "",
        "password": "b1b377374ff0075f7521e",
        "phone": "2345678901",
        "type": 2
    }
}
```

## Download submission list

**You send:**  id of the assignment

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions along with complete updated data of user.

**Endpoint:** 
     /submission/{id}

**Authorization Token:** user's auth token required

**Request:**
`GET HTTP/1.1`

**Successful Response:**
 The excel file will be downloaded in the broweser.


# Frontend

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `yarn start`

Runs the app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

### `yarn test`

Launches the test runner in the interactive watch mode.<br />
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `yarn build`

Builds the app for production to the `build` folder.<br />
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br />
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `yarn eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

### Code Splitting

This section has moved here: https://facebook.github.io/create-react-app/docs/code-splitting

### Analyzing the Bundle Size

This section has moved here: https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size

### Making a Progressive Web App

This section has moved here: https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app

### Advanced Configuration

This section has moved here: https://facebook.github.io/create-react-app/docs/advanced-configuration

### Deployment

This section has moved here: https://facebook.github.io/create-react-app/docs/deployment

### `yarn build` fails to minify

This section has moved here: https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify


  

