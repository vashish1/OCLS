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

const AssignmentModal= props => {
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
    const [imageAsFile,setImageAsFile]= useState()
    const userData=JSON.parse(localStorage.getItem('user'))
    const userType=userData.type
    const [uniqueId,setUniqueId]=useState("")
  const handleImageAsFile=(e)=>{
    const image=e.target.files[0];
    
    setImageAsFile((imageFile)=>image)
  }
  const classid=JSON.parse(localStorage.getItem('classid'))
  const handleCreateAssignment= async (e)=>{
    e.preventDefault()
    const item={desc,currentTime,currentClass,imageAsFile}
    
    const requestData = new FormData();
  requestData.set('description',desc);
  requestData.set('date',currentTime);
  requestData.set('class_code',currentClass);
  requestData.set('file', imageAsFile);
  console.log(currentTime)
  console.log(currentClass)
    let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/assignment/add",
    {
        method:"POST",
        headers:{
            "Accept": "*/*",
            authorization:`Bearer ${userToken}`
        },
        body:requestData
    });
    result = await result.json();
    console.log(result)
    handleClose()
}
    
    const handleSubmitAssignment= async (e)=>{
      
      const requestData = new FormData();
      requestData.set('id',uniqueId);
      requestData.set('file', imageAsFile);
      
      let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/assignment/sub",
      {
          method:"POST",
          headers:{
              "Accept": "*/*",
              authorization:`Bearer ${userToken}`
              
          },
          body:requestData
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
  const handleSubInput = (e) => {
    setUniqueId(e.target.value)
    
  };
  
  const body = userType==1?(
    <div style={modalStyle} className={classes.paper}>
      <h2 id="simple-modal-title">Create Assignment</h2>
       <TextField style={textField} id="outlined-basic" onChange={(e)=>setDesc(e.target.value)} label='Description' placeholder='Enter Subject' variant="outlined" fullWidth required/>
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
      <Button type='submit' onClick={handleCreateAssignment} color='primary' variant="contained" style={btnstyle} fullWidth>Create Assignment</Button>
    </div>
  ):(<div style={modalStyle} className={classes.paper}>
    <h2 id="simple-modal-title">Submit Assignment</h2>
     <input
     accept="image/*,file_extension/*"
     className={classes.input}
     id="contained-button-file"
     className="newUpload"
     multiple
     type="file"
     onChange={handleImageAsFile}
   />
  
      <TextField style={textField} name="assignment no." id="outlined-basic" onChange={handleSubInput} label='Assignment No.' placeholder='Enter Assignment No.' variant="outlined" fullWidth required/>
    <Button id="assignsub" type='submit' onClick={handleSubmitAssignment} color='primary' variant="contained" style={btnstyle} fullWidth>Submit Assignment</Button>
  </div>);

  return (
    <div>
    
    <List>
          
            <ListItem button onClick={handleOpen}>
              <ListItemIcon> <NotificationImportantIcon /></ListItemIcon>
              {userType==1?<ListItemText label="Create Class"> Create Assignment</ListItemText>:<ListItemText label="Create Class"> Submit Assignment</ListItemText>}
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
export default AssignmentModal