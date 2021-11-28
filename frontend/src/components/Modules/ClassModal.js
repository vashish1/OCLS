import React, {useEffect, useState} from 'react';
import { Grid,Paper, Avatar, TextField, Button, Typography,Link } from '@material-ui/core'
import { makeStyles } from '@material-ui/core/styles';
import Modal from '@material-ui/core/Modal';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import List from '@material-ui/core/List'
import NotificationImportantIcon from '@material-ui/icons/NotificationImportant';
function rand() {
  return Math.round(Math.random() * 20) - 10;
}

function getModalStyle() {
  const top = 50 + rand();
  const left = 50 + rand();

  return {
    top: `${top}%`,
    left: `${left}%`,
    transform: `translate(-${top}%, -${left}%)`,
  };
}

const useStyles = makeStyles((theme) => ({
  paper: {
    position: 'absolute',
    width: 400,
    backgroundColor: theme.palette.background.paper,
    border: '2px solid #000',
    boxShadow: theme.shadows[5],
    padding: theme.spacing(2, 4, 3),
  },
}));

export default function SimpleModal() {
  const classes = useStyles();
  // getModalStyle is not a pure function, we roll the style only on the first render
  const avatarStyle={backgroundColor:'#1bbd7e'}
    const btnstyle={margin:'8px 0'}
    const textField={margin:'10px auto'};
  const [modalStyle] = React.useState(getModalStyle);
  const [open, setOpen] = React.useState(false);
  const [subject,setSubject]=useState("");
  const userToken= localStorage.getItem('token')
  const userData=JSON.parse(localStorage.getItem('user'))
  const userType=userData.type
  useEffect(() => {
    
  }, [])
    const handleCreateClass= async ()=>{
        let item={subject};
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/add",
        {
            method:"POST",
            headers:{
                "Content-Type": "application/json",
                "Accept": "application/json",
                authorization: `Bearer ${userToken}`,
                
            },
            body:JSON.stringify(item)
        });
        result = await result.json();
        console.log(result)
        handleClose()
    }
    const [classCode,setClassCode]=useState()
    const handleJoinClass= async ()=>{
        let item={classCode};
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/join",
        {
            method:"POST",
            headers:{
                "Content-Type": "application/json",
                "Accept": "application/json",
                authorization: `Bearer ${userToken}`,
                
            },
            body:JSON.stringify(item)
        });
        result = await result.json();
        console.log(result)
        handleClose()
    }
  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const body = userType==1?(
    
    <div style={modalStyle} className={classes.paper}>
    
      <h2 id="simple-modal-title">Create Class</h2>
      <TextField style={textField} id="outlined-basic" onChange={(e)=>setSubject(e.target.value)} label='Subject' placeholder='Enter Subject' variant="outlined" fullWidth required/>
                
      <Button type='submit' onClick={handleCreateClass} color='primary' variant="contained" style={btnstyle} fullWidth>Create Class</Button>
    </div>
  ):(
    
    <div style={modalStyle} className={classes.paper}>
    
      <h2 id="simple-modal-title">Join Class</h2>
      <TextField style={textField} id="outlined-basic" onChange={(e)=>setClassCode(e.target.value)} label='Class Code' placeholder='Enter Class Code' variant="outlined" fullWidth required/>
                
      <Button type='submit' onClick={handleJoinClass} color='primary' variant="contained" style={btnstyle} fullWidth>Join Class</Button>
    </div>
  );

  return (
    <div>
    
    <List>
          
            <ListItem button onClick={handleOpen}>
              <ListItemIcon> <NotificationImportantIcon /></ListItemIcon>
              {userType==1?<ListItemText label="Create Class"> Create Class</ListItemText>:<ListItemText label="Join Class"> Join Class</ListItemText>}
              
            </ListItem>
        </List>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="simple-modal-title"
        aria-describedby="simple-modal-description"
      >
        {body}
      </Modal>
    </div>
  );
}
