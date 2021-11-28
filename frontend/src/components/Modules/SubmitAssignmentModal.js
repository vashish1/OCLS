import React, {useState} from 'react';
import { Grid,Paper, Avatar, TextField, Button, Typography,Link } from '@material-ui/core'
import { makeStyles } from '@material-ui/core/styles';
import Modal from '@material-ui/core/Modal';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import List from '@material-ui/core/List'
import * as moment from 'moment';
import NotificationImportantIcon from '@material-ui/icons/NotificationImportant';
import { v4 as uuidv4 } from 'uuid';
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
const assignmentId=uuidv4()
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

export default function AssignmentModal() {
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
    const currentTime = moment(yourDate, 'DD-MMM-YY');
    const userToken=localStorage.getItem('token')
    const currentClass=localStorage.getItem('class-code')
    const [imageAsFile,setImageAsFile]= useState()
  const handleImageAsFile=(e)=>{
    const image=e.target.files[0];
    setImageAsFile((imageFile)=>image)
  }
  
    const handleSubmitAssignment= async (e)=>{
        e.preventDefault()
        let item={assignmentId,file:imageAsFile};
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/assignment/sub",
        {
            method:"POST",
            headers:{
                "Content-Type": "application/json",
                "Accept": "application/json",
                authorization:`Bearer ${userToken}`
                
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
  
  const body = (
    <div style={modalStyle} className={classes.paper}>
      <h2 id="simple-modal-title">Upload Assignment</h2>
       
       <input
       accept="image/*,file_extension/*"
       className={classes.input}
       id="contained-button-file"
       className="newUpload"
       multiple
       type="file"
       onChange={handleImageAsFile}
     />
     <label htmlFor="contained-button-file">
       <Button variant="contained" color="primary" component="span">
         Upload
       </Button>
     </label>
      <Button type='submit' onClick={handleSubmitAssignment} color='primary' variant="contained" style={btnstyle} fullWidth>Create Announcement</Button>
    </div>
  );

  return (
    <div>
    
    <List>
          
            <ListItem button onClick={handleOpen}>
              <ListItemIcon> <NotificationImportantIcon /></ListItemIcon>
              <ListItemText label="Create Class"> Submit Assignment</ListItemText>
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
