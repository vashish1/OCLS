import React, { useState,useEffect } from 'react'
import { Grid,Paper, Avatar, TextField, Button, Typography,Link } from '@material-ui/core'
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import {List,ListItem,ListItemIcon,ListItemText} from '@material-ui/core';
import ClassIcon from '@material-ui/icons/Class';
import { useGlobal } from '../context/SignupContext';
const GetClass = () => {
    const paperStyle={
        padding :20,
        width:300, 
        margin:"0 auto"
    }
    const avatarStyle={backgroundColor:'#1bbd7e'}
    const btnstyle={margin:'8px 0'}
    const textField={margin:'10px auto'};
    const [classCode,setClassCode]=useState(""); 
    
//   const userClasses=JSON.parse(localStorage.getItem('classes'))
const userToken =localStorage.getItem('token')
  const handleGetClasses= async ()=>{
   
      let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/get",
      {
          method:"GET",
          headers:{
              "Content-Type": "application/json",
              "Accept": "application/json",
              authorization: `Bearer ${userToken}`,
          },
      });
      result = await result.json();
      localStorage.setItem('classes',JSON.stringify(result.data))
      
  }
  

    return (
        <div>
        <List>
        <ListItem button onClick={handleGetClasses}>
              <ListItemIcon> <ClassIcon /></ListItemIcon>
              <ListItemText label="Create Class">Classes</ListItemText>
            </ListItem>
        </List>
        </div>
    )
}

export default GetClass
