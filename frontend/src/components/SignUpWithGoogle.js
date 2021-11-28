import React, { useState } from 'react'
import { Grid, Paper, Avatar, Typography, TextField, Button } from '@material-ui/core'
import AddCircleOutlineOutlinedIcon from '@material-ui/icons/AddCircleOutlineOutlined';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';
// import FormLabel from '@material-ui/core/FormLabel';
// import Axios from 'axios';
import { useNavigate } from 'react-router';
const SignUpWithGoogle = () => {
    const paperStyle = { padding: 20, width: 300, margin: "0 auto" }
    const headerStyle = { margin: 0 }
    const avatarStyle = { backgroundColor: '#1bbd7e' }
    // const [userType, setUserType]=useState(0);
    const [value, setValue]=useState(1)
    // const handleSignUp= async (e)=>{
        
    //     const user={
    //         email:email,
    //         password:password,
    //         type:userType
    //     }
    //     const {data} = await Axios.post('https://thawing-mountain-02190.herokuapp.com/signup',
    //         user
    //       )
    //       console.log(data)
    // }

    const history=useNavigate();
    const handleSignUp= async (e)=>{
        let url_string =window.location.href
        var url = new URL(url_string);
        const email= url.searchParams.get("email");
        const name= url.searchParams.get("name");
        e.preventDefault();
        let item={email:email,name:name,type:value};
        console.log(item)
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/signup",
        {
            method:"POST",
            headers:{
                "Accept": "application/json",
                "Content-Type": "application/json"
            },
            body:JSON.stringify(item)
        });
        result = await result.json();
        console.log(result)
        history('/login')
        
    }
    const handleLogin= async (e)=>{
        let url_string =window.location.href
        var url = new URL(url_string);
        const email= url.searchParams.get("email");
        const name= url.searchParams.get("name");
        e.preventDefault();
        let item={email:email,name:name,type:value};
        console.log(item)
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/login/google",
        {
            method:"POST",
            headers:{
                "Accept": "application/json",
                "Content-Type": "application/json"
            },
            body:JSON.stringify(item)
        });
        result = await result.json();
        console.log(result)
        localStorage.setItem('user', JSON.stringify(result.data.User)) ;
        localStorage.setItem('token', result.data.Token) ; 
        history('/dashboard')
        
    }

    
    function handleClick(event) {
        if (event.target.value === "teacher") {
          setValue(1);
        } 
        else if(event.target.value === "student") {
          setValue(2);
        }
      }
      const login=localStorage.getItem('bool')
    return (
        <Grid>
            <Paper style={paperStyle}>
                <Grid align='center'>
                    <Avatar style={avatarStyle}>
                        <AddCircleOutlineOutlinedIcon />
                    </Avatar>
                    <h2 style={headerStyle}>Sign Up</h2>
                    <Typography component={"span"} variant='caption' gutterBottom>Please choose the role to create an account !</Typography>
                </Grid>
                <form onSubmit={handleSignUp}>
                    <FormControl component="fieldset">
                    <RadioGroup
                        aria-label="role"
                        defaultValue="teacher"
                        name="radio-buttons-group"
                    >
                        <FormControlLabel value="teacher" onClick={handleClick} control={<Radio />} label="teacher" />
                        <FormControlLabel value="student" onClick={handleClick} control={<Radio />} label="student" />
                    </RadioGroup>
                    {login?<Button type='submit' variant='contained' onClick={handleLogin} color='primary'>Sign up</Button>
                    :<Button type='submit' variant='contained' onClick={handleSignUp} color='primary'>Sign up</Button>}
                    </FormControl>
                   
                    
                    </form>
            </Paper>
        </Grid>
    )
}

export default SignUpWithGoogle;