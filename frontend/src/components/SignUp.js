import React, { useState,useEffect } from 'react'
import { Grid, Paper, Avatar, Typography, TextField, Button } from '@material-ui/core'
import AddCircleOutlineOutlinedIcon from '@material-ui/icons/AddCircleOutlineOutlined';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';
// import FormLabel from '@material-ui/core/FormLabel';
import Checkbox from '@material-ui/core/Checkbox';
// import Axios from 'axios';
import SignUpWithGoogle from './SignUpWithGoogle';
import { useNavigate } from 'react-router';
const SignUp = () => {
    const paperStyle = { padding: 20, width: 300, margin: "0 auto" }
    const headerStyle = { margin: 0 }
    const avatarStyle = { backgroundColor: '#1bbd7e' }
    const marginTop = { marginTop: 5 }
    const textField={margin:'10px auto'}
    const [email, setEmail] = useState('');
    const [password, setPassword]=useState('');
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
        e.preventDefault();
        let item={email,password,value};
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
        history('/login')
        
    }
    const handleSignUpWithGoogle= async (e)=>{
        console.log("clicked")
      let Url;
      fetch("http://localhost:9000/signup/google").then(
        res=>{
             console.log(res.url)

          Url=res.url
          window.location.replace(Url)
          // console.log(url)
      

          },
     
        // (res) => console.log(res.url)
      )
        
    }

    
    function handleClick(event) {
        if (event.target.value === "teacher") {
          setValue(1);
        } 
        else if(event.target.value === "student") {
          setValue(2);
        }
      }
      const goToLogin=()=>{
          history('/login')
      }
    return (
        <Grid>
            <Paper style={paperStyle}>
                <Grid align='center'>
                    <Avatar style={avatarStyle}>
                        <AddCircleOutlineOutlinedIcon />
                    </Avatar>
                    <h2 style={headerStyle}>Sign Up</h2>
                    <Typography component={"span"} variant='caption' gutterBottom>Please fill this form to create an account !</Typography>
                </Grid>
                <form onSubmit={handleSignUp}>
                    <TextField style={textField} value={email} onChange={(e)=>setEmail(e.target.value)} id="outlined-basic" variant="outlined" fullWidth label='Email' placeholder="Enter your email" />
                    <TextField style={textField} value={password} onChange={(e)=>setPassword(e.target.value)} id="outlined-basic" variant="outlined" fullWidth label='Password' placeholder="Enter your password"/>
                    <FormControl component="fieldset">
                    <RadioGroup
                        aria-label="gender"
                        defaultValue="female"
                        name="radio-buttons-group"
                    >
                        <FormControlLabel value="teacher" onClick={handleClick} control={<Radio />} label="teacher" />
                        <FormControlLabel value="student" onClick={handleClick} control={<Radio />} label="student" />
                    </RadioGroup>
                    </FormControl>
                    <Typography> Already A user?
                     <a href="#" onClick={goToLogin} >
                        Login
                    </a>
                    </Typography>
                    <Button type='submit' variant='contained' onClick={handleSignUp} color='primary'>Sign up</Button><br/><br/>
                    <Button type='submit' variant='contained' onClick={handleSignUpWithGoogle} color='primary'>Google SignUp</Button>
                    </form>
            </Paper>
        </Grid>
    )
}

export default SignUp;