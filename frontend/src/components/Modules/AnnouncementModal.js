import React, {useState} from 'react';
import { Grid,Paper, Avatar, TextField, Button, Typography,Link } from '@material-ui/core'
import { makeStyles } from '@material-ui/core/styles';
import Modal from '@material-ui/core/Modal';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import List from '@material-ui/core/List'
import * as moment from 'moment';
import AddCircleOutlineIcon from '@material-ui/icons/AddCircleOutline';
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

export default function AnnouncementModal() {
  const classes = useStyles();
  // getModalStyle is not a pure function, we roll the style only on the first render
    const btnstyle={margin:'8px 0'}
    const textField={margin:'10px auto'};
    const textField2={
        margin:'10px auto',};
  const [modalStyle] = React.useState(getModalStyle);
  const [open, setOpen] = React.useState(false);
  const [teacherName,setTeacherName]=useState("");
    const [desc,setDesc]=useState("");
    const yourDate = new Date();
    const currentTime = moment(yourDate).format('DD-MMM-YY');
    const userToken=localStorage.getItem('token')
    const currentClass=localStorage.getItem('class-code')
    const userData=JSON.parse(localStorage.getItem('user'))
    const userType=userData.type
    const handleCreateAnnouncement= async ()=>{
        let item={teacherName,currentClass,desc,currentTime};
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/announcement/add",
        {
            method:"POST",
            headers:{
                "Content-Type": "application/json",
                "Accept": "application/json",
                authorization:`Bearer ${userToken}`
                
            },
            body:JSON.stringify({"teacher_name":teacherName,"classcode":currentClass,"description":desc,"timestamp":currentTime})
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

  const body =userType==1? (
    <div style={modalStyle} className={classes.paper}>
      <h2 id="simple-modal-title">Create Announcement</h2>
      <TextField style={textField} id="outlined-basic" onChange={(e)=>setTeacherName(e.target.value)} label='Teacher Name' placeholder='Enter Teacher Name' variant="outlined" fullWidth required/>
       <TextField style={textField} id="outlined-basic" onChange={(e)=>setDesc(e.target.value)} label='Description' placeholder='Enter Description' variant="outlined" fullWidth required/>
        
      <Button type='submit' onClick={handleCreateAnnouncement} color='primary' variant="contained" style={btnstyle} fullWidth>Create Announcement</Button>
    </div>
  ):null;

  return (
    <div>
    
    <List>
          
            {userType==1?(<ListItem button onClick={handleOpen}>
              <ListItemIcon> <AddCircleOutlineIcon/></ListItemIcon>
              <ListItemText label="Create Class"> Create Announcement</ListItemText>
            </ListItem>):null}
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
