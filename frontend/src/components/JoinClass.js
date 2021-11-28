// import React, { useState } from 'react'
// import { Grid,Paper, Avatar, TextField, Button, Typography,Link } from '@material-ui/core'
// import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
// import FormControlLabel from '@material-ui/core/FormControlLabel';
// import Checkbox from '@material-ui/core/Checkbox';
// const JoinClass = () => {
//     const paperStyle={
//         padding :20,
//         width:300, 
//         margin:"0 auto"
//     }
//     const avatarStyle={backgroundColor:'#1bbd7e'}
//     const btnstyle={margin:'8px 0'}
//     const textField={margin:'10px auto'};
//     const [classCode,setClassCode]=useState("");
//     const handleJoinClass= async ()=>{
//         let item={classCode};
//         let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/add",
//         {
//             method:"POST",
//             headers:{
//                 "Content-Type": "application/json",
//                 "Accept": "application/json"
                
//             },
//             body:JSON.stringify(item)
//         });
//         result = await result.json();
//         console.log(result)
//     }
//     return (
//         <div>
//         <Grid>
//             <Paper  style={paperStyle}>
//                 <Grid align='center'>
//                      <Avatar style={avatarStyle}><LockOutlinedIcon/></Avatar>
//                     <h2>Create Your Class</h2>
//                 </Grid>
//                 <TextField style={textField} id="outlined-basic" onChange={(e)=>setClassCode(e.target.value)} label='Username' placeholder='Enter username' variant="outlined" fullWidth required/>
                
//                 <Button type='submit' onClick={handleJoinClass} color='primary' variant="contained" style={btnstyle} fullWidth>Join Class</Button>
                
//             </Paper>
//         </Grid>
//         </div>
//     )
// }

// export default JoinClass
