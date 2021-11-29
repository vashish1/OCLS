import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import { useNavigate } from 'react-router';
import AssignmentModal from './AssignmentModal';

const useStyles = makeStyles({
  root: {
    minWidth: 275,
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
    marginBottom: 12,
  },
});

export default function AssignOutlinedCard(props) {
  const classes = useStyles();
  const bull = <span className={classes.bullet}>•</span>;
    const {name,description,classcode,date,file,form,id,id1}=props.props
    const history=useNavigate()
    
    
    localStorage.setItem('class-code',classcode)
    const userData=JSON.parse(localStorage.getItem('user'))

  
  const userType=userData.type
  const goToQuiz=()=>{
    localStorage.setItem('quizform',JSON.stringify(form))
    localStorage.setItem('assignmentId',id)
    history(`/class/${id}/assignment/takequiz`)
  }
  return (
             
                <Card className={classes.root} variant="outlined">
                <CardContent>
                <Typography className={classes.pos} color="textSecondary">
                   Assignment No. : {id}
                  </Typography>
                  
                  <Typography variant="h5" component="h2">

                    Teacher: {name}
                  </Typography>
                  <Typography variant="h5" component="h2">

                    Description: {description}
                  </Typography>
                 {file.filename?<Button href={file.filename} variant="contained" color="primary">Download</Button>:
                 <Button onClick={goToQuiz} variant="contained" color="primary">Submit Quiz</Button>}
                 <Typography className={classes.pos} color="textSecondary">
                   Last Submission date- {date}
                  </Typography>
                  {userType==2&&file.filename?<AssignmentModal/>:null}
                </CardContent>
               
              </Card>
        
  );
}