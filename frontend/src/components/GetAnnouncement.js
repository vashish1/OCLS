import React, { useState } from 'react'
import { Grid,Paper, Avatar, TextField, Button, Typography,Link } from '@material-ui/core'
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import * as moment from 'moment'
import MiniDrawer from './Modules/NavBar'
const yourDate = new Date()
const currentTime = moment(yourDate).format('DD-MMM-YY');
const GetAnnouncement = () => {
    
    const paperStyle={
        padding :20,
        margin:"0 auto"
    }
    const avatarStyle={backgroundColor:'#1bbd7e'}
    const btnstyle={margin:'8px 0'}
    const textField={margin:'10px auto'};
    const [teacherName,setTeacherName]=useState("");
    const [classCode,setClassCode]=useState("");
    const handleAnnouncement= async ()=>{
        let item={classCode};
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/announcement/get",
        {
            method:"POST",
            headers:{
                "Content-Type": "application/json",
                "Accept": "application/json"
                
            },
            body:JSON.stringify(item)
        });
        result = await result.json();
        console.log(result)
    }
    return (
        <div>
        <Grid>
            <Paper  style={paperStyle}>
                
                 <Typography> 
                  `${currentTime}`
                </Typography>
                <Button type='submit' onClick={handleAnnouncement} color='primary' variant="contained" style={btnstyle} fullWidth>Get Announcement</Button>
                
            </Paper>
        </Grid>
        </div>
    )
}

export default GetAnnouncement
