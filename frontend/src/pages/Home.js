import React, { useState,useEffect } from 'react';
import { useNavigate } from "react-router-dom";
import { Grid,Paper, Avatar, TextField, Button, Typography,Link } from '@material-ui/core'
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Axios from 'axios';
import { useGlobal } from '../context/SignupContext';
import { Public } from '@material-ui/icons';
const initialValues={
    email: '',
    password: ''
}

const Home=({handleChange})=>{
    
    const [loginValues, setLoginValues] = useState(initialValues);
    
    const paperStyle={
        padding :20,
        width:700, 
        
        margin:"0 auto",
        marginTop:300,
    }
    const avatarStyle={backgroundColor:'#1bbd7e'}
    const btnstyle={margin:'8px 0'}
  
    // const [inpData, setInpData] = useState(defData);
    const history = useNavigate();


    const goToLogin=()=>{
        history('/login')
    }
  
    
    return(
        <Grid>
            <Paper  style={paperStyle}>
                <Grid align='center'>
                     <Avatar style={avatarStyle}><Public/></Avatar>
                     <h1>Welcome To EAssign</h1>
                    <h2>Get Started</h2>
                    
                </Grid>
                
                <Button type='submit' onClick={goToLogin} color='primary' variant="contained" style={btnstyle} fullWidth>Log In / Sign Up</Button>
                
                
               
            </Paper>
        </Grid>
    )
}

export default Home
