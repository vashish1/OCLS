## Signup

Input- `{
    "email":"yashi.gupta@gmail.com",
    "password":"qwerty",
    "type":1
}`

Output- `{
    "success": true,
    "message": "Signup SuccessFul"
}

In this we will consecutively call two api's one will give the name and email of the user or error
and then we will ask the user type Student or teacher on that we will call the sighnup api


/callback Response 
Output-`{ "Name":"Yashi Gupta","Email":"vashishtiv@gmail.com" }`

add the type-> Student- 1
            -> Teacher- 2

Call the signup api

## login by google :

Do the same as done in google signup and then call the /login/google api

Input-`{ "Name":"Yashi Gupta","Email":"vashishtiv@gmail.com" }`

Output- `{
    "success": true,
    "message": "Log In successful",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inlhc2hpLmd1cHRhQGdtYWlsLmNvbSIsInR5cGUiOjF9.w1qyPVtDE2bgavD-Q3rU-BjQrVDbw2AWodMimD3JJvo"
}`
or
`
`{
    "success": false,
    "message": "------",
    "error": "-----"
}`

## Create Class

Input- `{
    "subject":"----"
}

Output- `{
    "success": true,
    "message": "Class created Successfully",
    "data": "000355shi" ----->this is the class code
}`

or
{
    "success": false,
    "error": "-----"
}

## join class

Input- `{
    "class_code": "73fbbf"
}`

Output- `{
    "success": true,
    "message":"------------"
}`
 or
{
    "success": false,
    "error": "Error while joining the class"
}

## Get all the classes 

