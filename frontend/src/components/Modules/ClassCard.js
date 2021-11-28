import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import { ArrowForward } from '@material-ui/icons';
import { useNavigate } from 'react-router';

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

export default function OutlinedCard(props) {
  const classes = useStyles();
  const bull = <span className={classes.bullet}>•</span>;
    const {subject,code,id}=props.props
    const history=useNavigate()
    const handleAnnouncement=()=>{

      localStorage.setItem('class-code',code)
        history(`/class/${id}/announcement`)
    }
    const handleAssignment=()=>{

    localStorage.setItem('class-code',code)
      history(`/class/${id}/assignment`)
  }
    

  return (
             
                <Card className={classes.root} variant="outlined">
                <CardContent>
                <Typography className={classes.pos} color="textSecondary">
                   {id+1}.
                  </Typography>
                  <Typography className={classes.title} color="textSecondary" gutterBottom>
                  Code: {code}
                  </Typography>
                  <Typography variant="h5" component="h2">

                    {subject}
                  </Typography>
                  
                </CardContent>
                <CardActions>
                  <Button size="small" onClick={handleAnnouncement}>Announcement<ArrowForward/></Button>
                </CardActions>
                <CardActions>
                  <Button size="small" onClick={handleAssignment}>Assignment<ArrowForward/></Button>
                </CardActions>
              </Card>
        
  );
}