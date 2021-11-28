import React, { useState,useEffect } from 'react';
import { useNavigate } from "react-router-dom";
import { Grid,Paper, Avatar, TextField, Button, Typography,Link } from '@material-ui/core'
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Axios from 'axios';
import LoginWithGoogle from './LoginWithGoogle';
import { useGlobal } from '../context/SignupContext';
const initialValues={
    email: '',
    password: ''
}

const Login=({handleChange})=>{
    
    const [loginValues, setLoginValues] = useState(initialValues);
    
    const paperStyle={
        padding :20,
        width:300, 
        margin:"0 auto"
    }
    const avatarStyle={backgroundColor:'#1bbd7e'}
    const btnstyle={margin:'8px 0'}
    const textField={margin:'10px auto'}
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [user,setUser]=useState('')
    const [token,setToken]=useState('')
    // const [inpData, setInpData] = useState(defData);
    const history = useNavigate();
    const userToken =localStorage.getItem('token')
    useEffect(() => {
        const loggedInUser = JSON.parse(localStorage.getItem("user-info"));
        const tokenHandler= localStorage.getItem('token')
        
        if (loggedInUser && tokenHandler)
        {
        //   const foundUser = JSON.parse(loggedInUser);
          setUser(loggedInUser);
          history("/dashboard")
        }
      }, []);

    const handleInput = (e) => {
        
        setLoginValues({...loginValues, [e.target.name]: e.target.value});
    }
  
    const handleSignIn= async (e)=>{
        e.preventDefault();
        let item={email,password};

        const {data} = await Axios.post('https://thawing-mountain-02190.herokuapp.com/login',
            loginValues
          )
          console.log(data)

        // set the state of the user
        // store the user in localStorage
        
        console.log(data.data.Token)
        setUser(data.data);
        setToken(data.data.Token);
        localStorage.setItem('user', JSON.stringify(data.data.User)) ;
        localStorage.setItem('token', data.data.Token) ; 
        history('/dashboard')
    }
    const handleLogInWithGoogle= async (e)=>{
        // let url_string =window.location.href
        // var url = new URL(url_string);
        // const email= url.searchParams.get("email");
        // const name= url.searchParams.get("name");
        const login="login"
        localStorage.setItem('bool',login)
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
    const goToSignUp=()=>{
        history('/signup')
    }
    return(
        <Grid>
            <Paper  style={paperStyle}>
                <Grid align='center'>
                     <Avatar style={avatarStyle}><LockOutlinedIcon/></Avatar>
                    <h2>Sign In</h2>
                </Grid>
                <TextField style={textField} name="email" id="outlined-basic" onChange={handleInput} label='Username' placeholder='Enter username' variant="outlined" fullWidth required/>
                <TextField style={textField} name="password" id="outlined-basic" onChange={handleInput} variant="outlined" label='Password' placeholder='Enter password' type='password' fullWidth required/>
                
                <Button type='submit' onClick={handleSignIn} color='primary' variant="contained" style={btnstyle} fullWidth>Sign in</Button>
                <br/>
                <Button type='submit' onClick={handleLogInWithGoogle} color='primary' variant="contained" style={btnstyle} fullWidth>Google Sign in</Button>
                
                
                <Typography > Do you have an account ?
                     <Link href="#" onClick={goToSignUp} >
                        Sign Up 
                </Link>
                </Typography>
            </Paper>
        </Grid>
    )
}

export default Login