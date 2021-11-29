import React, {useState} from 'react';
import { TextField, Button} from '@material-ui/core'
import { makeStyles } from '@material-ui/core/styles';
import Modal from '@material-ui/core/Modal';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import List from '@material-ui/core/List'

import { Note } from '@material-ui/icons';
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

export default function GetSubmissionModal() {
  const classes = useStyles();
  // getModalStyle is not a pure function, we roll the style only on the first render
    const btnstyle={margin:'8px 0'}
    const textField={margin:'10px auto'};
    
  const [modalStyle] = React.useState(getModalStyle);
  const [open, setOpen] = React.useState(false);
 
    // const yourDate = new Date();
   
    const [uniqueId,setUniqueId]=useState("")

    // const downloadSubmission =(uri, name) =>
    // {
    //     var link = document.createElement("a");
    //     link.setAttribute('download', name);
    //     link.href = uri;
    //     document.body.appendChild(link);
    //     link.click();
    //     link.remove();
    // }
    // const handleSubmission= async (e)=>{
    //     e.preventDefault()
    //     let result=await fetch(`https://thawing-mountain-02190.herokuapp.com/submission/${uniqueId}`,
    //     {
    //         method:"GET",
            
    //     });
    //     result = await result.json();
    //     console.log(result.data)
    //     downloadSubmission(result.data.filename,"submissions")
    //     handleClose()
    // }
  
    
  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };
  const handleSubInput=(e)=>{
    setUniqueId(e.target.value)
  }
  const body = (
    <div style={modalStyle} className={classes.paper}>
      <h2 id="simple-modal-title">Get Submissions</h2>
      <TextField style={textField} name="assignment no." id="outlined-basic" onChange={handleSubInput} label='Assignment No.' placeholder='Enter Assignment No.' variant="outlined" fullWidth required/>
       
     
      <Button type='submit' href={`https://thawing-mountain-02190.herokuapp.com/submission/${uniqueId}`} color='primary' variant="contained" style={btnstyle} fullWidth>Get Submissions</Button>
    </div>
  );

  return (
    <div>
    
    <List>
          
            <ListItem button onClick={handleOpen}>
              <ListItemIcon> <Note/></ListItemIcon>
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
