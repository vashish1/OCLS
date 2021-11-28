import React,{useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import { useNavigate } from 'react-router';
import { Avatar, Grid, Paper, TextField} from '@material-ui/core';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Axios from 'axios';
import { People } from '@material-ui/icons';

const useStyles = makeStyles({
  root: {
    minWidth: 400,
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    fontSize: 14,
  },
  pos: {
      fontSize:24,
    margin: 12,
  },
});

export default function UpdateProfile() {
    const user=JSON.parse(localStorage.getItem('user'))
    
    const defaultUser=user.type==2?{
        classcode:user.class,
        name:'',
        admno:'',
        phone:'',
        email: '',
        password:user.password,
        type:user.type
        
    }:{
        classcode:user.class,
        name:'',
        phone:'',
        email: user.email,
        password:user.password,
        type:user.type
    }
    const [userProfile, setUserProfile] = useState(defaultUser);
    const handleInput = (e) => {
        
        setUserProfile({...userProfile, [e.target.name]: e.target.value});
    }
    const textField={margin:'10px auto'}
    const paperStyle={
        padding :20,
        width:300, 
        margin:"0 auto"
    }
    const avatarStyle={backgroundColor:'#1bbd7e'}
    const btnstyle={margin:'8px 0'}
    const navigateto=useNavigate()
   const userToken=localStorage.getItem('token')
   
    const handleUpdateProfile= async (e)=>{
        e.preventDefault();
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/user/update",
        {
            method:"POST",
            headers:{
                "Content-Type": "application/json",
                "Accept": "application/json",
                authorization: `Bearer ${userToken}`,
            },
            body:JSON.stringify(userProfile)
        });
        result = await result.json();
        console.log(result)
        // localStorage.setItem('user', JSON.stringify(result.data.User)) ;
        navigateto('/profile')
    }
    
    
  const classes = useStyles();
  const bull = <span className={classes.bullet}>•</span>;
    
    
    const backToProfile=()=>{
        navigateto('/profile')
    }
  return (
    <Grid>
            <Paper  style={paperStyle}>
                <Grid align='center'>
                    <People/>
                    <h2>Update Profile</h2>
                </Grid>
                
                <TextField style={textField} name="name" id="outlined-basic" onChange={handleInput} variant="outlined" label='name' placeholder='Enter name'  fullWidth required/>
                {user.type==2?(<TextField style={textField} name="admno" id="outlined-basic" onChange={handleInput} variant="outlined" label='admission no' placeholder='Enter admission no' fullWidth required/>):null}
                <TextField style={textField} name="phone" id="outlined-basic" onChange={handleInput} variant="outlined" label='phone no.' placeholder='Enter phone no' fullWidth required/>
                
                <Button type='submit' onClick={handleUpdateProfile} color='primary' variant="contained" style={btnstyle} fullWidth>Update Profile</Button>
                <Button type='submit' onClick={backToProfile} color='primary' variant="contained" style={btnstyle} fullWidth>back to profile</Button>
        
                
            </Paper>
        </Grid>

  );
}
