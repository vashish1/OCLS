// import React, { useState } from 'react'
// import { Grid, Paper, Avatar, Typography, TextField, Button } from '@material-ui/core'
// import AddCircleOutlineOutlinedIcon from '@material-ui/icons/AddCircleOutlineOutlined';
// import Radio from '@material-ui/core/Radio';
// import RadioGroup from '@material-ui/core/RadioGroup';
// import FormControlLabel from '@material-ui/core/FormControlLabel';
// import FormControl from '@material-ui/core/FormControl';
// // import FormLabel from '@material-ui/core/FormLabel';
// // import Axios from 'axios';
// import { useNavigate } from 'react-router';
// const SignUpWithGoogle = () => {
//     const paperStyle = { padding: 20, width: 300, margin: "0 auto" }
//     const headerStyle = { margin: 0 }
//     const avatarStyle = { backgroundColor: '#1bbd7e' }
//     // const [userType, setUserType]=useState(0);
//     const [value, setValue]=useState(1)
//     // const handleSignUp= async (e)=>{
        
//     //     const user={
//     //         email:email,
//     //         password:password,
//     //         type:userType
//     //     }
//     //     const {data} = await Axios.post('https://thawing-mountain-02190.herokuapp.com/signup',
//     //         user
//     //       )
//     //       console.log(data)
//     // }

//     const history=useNavigate();

    
//     function handleClick(event) {
//         if (event.target.value === "teacher") {
//           setValue(1);
//         } 
//         else if(event.target.value === "student") {
//           setValue(2);
//         }
//       }
      
//     return (
//         <Grid>
//             <Paper style={paperStyle}>
//                 <Grid align='center'>
//                     <Avatar style={avatarStyle}>
//                         <AddCircleOutlineOutlinedIcon />
//                     </Avatar>
//                     <h2 style={headerStyle}>Sign Up</h2>
//                     <Typography component={"span"} variant='caption' gutterBottom>Please choose the role to create an account !</Typography>
//                 </Grid>
//                 <form onSubmit={handleSignUp}>
//                     <FormControl component="fieldset">
//                     <RadioGroup
//                         aria-label="role"
//                         defaultValue="teacher"
//                         name="radio-buttons-group"
//                     >
//                         <FormControlLabel value="teacher" onClick={handleClick} control={<Radio />} label="teacher" />
//                         <FormControlLabel value="student" onClick={handleClick} control={<Radio />} label="student" />
//                     </RadioGroup>
//                     <Button type='submit' variant='contained' onClick={handleSignUp} color='primary'>Login</Button>
//                     </FormControl>
                   
                    
//                     </form>
//             </Paper>
//         </Grid>
//     )
// }

// export default SignUpWithGoogle;