import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import { useNavigate } from 'react-router';
import { ArrowBack } from '@material-ui/icons';

const useStyles = makeStyles({
  root: {
    minWidth: 400,
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    fontSize: 14,
  },
  pos: {
      fontSize:24,
    margin: 12,
  },
});

export default function Profile() {
  const classes = useStyles();
  const bull = <span className={classes.bullet}>•</span>;
    const user=JSON.parse(localStorage.getItem('user'))
    const history=useNavigate()
    const goToProfile=()=>{
        history('/profile/update')
    }
    const backToDashboard=()=>{
        history('/dashboard')
    }
  return (
    <Card className={classes.root}>
      <CardContent>
      <ArrowBack onClick={backToDashboard}/>
      <Typography variant="h5" component="h2">
          User Profile
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Name: {user.name}
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
         Email: {user.email}
        </Typography>
        {user.type==2?(<Typography className={classes.pos} color="textSecondary">
        Admission No.: {user.admno}
       </Typography>):null}
       <Typography className={classes.pos} color="textSecondary">
        Phone No.: {user.phone}
       </Typography>
          {user.type==1?(<Typography className={classes.pos} color="textSecondary">Role: Teacher</Typography>):(<Typography className={classes.pos} color="textSecondary">Role: Stduent</Typography>)}
          
      </CardContent>
      <CardActions>
        <Button onClick={goToProfile} size="small" variant="outlined" color="secondary">Update profile</Button>
      </CardActions>
    </Card>

  );
}
