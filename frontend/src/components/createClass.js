import React, { useState } from 'react'
import { Grid,Paper, Avatar, TextField, Button, Typography,Link } from '@material-ui/core'
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import SimpleModal from './Modules/ClassModal';
const CreateClass = () => {
    const paperStyle={
        padding :20,
        width:300, 
        margin:"0 auto"
    }
  
    return (
        <div>
        <SimpleModal/>
        </div>
    )
}

export default CreateClass
