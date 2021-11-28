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
        width:300, 
        margin:"0 auto"
    }
    const avatarStyle={backgroundColor:'#1bbd7e'}
    const btnstyle={margin:'8px 0'}
    const textField={margin:'10px auto'};


    const [teacherName,setTeacherName]=useState("");
    const [classCode,setClassCode]=useState("");
    const [desc,setDesc]=useState("");
    const handleAssignment= async ()=>{
        let item={classCode};
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/assignment/get",
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
                <Grid align='center'>
                     <Avatar style={avatarStyle}><LockOutlinedIcon/></Avatar>
                    <h2>Create Your Class</h2>
                </Grid>
                <TextField style={textField} id="outlined-basic" onChange={(e)=>setTeacherName(e.target.value)} label='Teacher Name' placeholder='Enter username' variant="outlined" fullWidth required/>
                <TextField style={textField} id="outlined-basic" onChange={(e)=>setClassCode(e.target.value)} label='Class Code' placeholder='Enter class Code' variant="outlined" fullWidth required/>
                <TextField style={textField} id="outlined-basic" onChange={(e)=>setDesc(e.target.value)} label='description' placeholder='Enter Description' variant="outlined" fullWidth required/>
                <Typography > 
                  `${currentTime}`
                </Typography>
                <Button type='submit' onClick={handleAssignment} color='primary' variant="contained" style={btnstyle} fullWidth>Create Announcement</Button>
                
            </Paper>
        </Grid>
        </div>
    )
}

export default GetAnnouncement
