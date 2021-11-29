import React,{useState,useEffect} from 'react';

import { useNavigate } from 'react-router-dom';
import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import {AppBar, Tabs, Tab} from '@material-ui/core';
import Toolbar from '@material-ui/core/Toolbar';
import List from '@material-ui/core/List';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import AccountCircle from '@material-ui/icons/AccountCircle';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import NotificationImportantIcon from '@material-ui/icons/NotificationImportant';
import ClassIcon from '@material-ui/icons/Class';
import AssignmentIcon from '@material-ui/icons/Assignment';
import GetAnnouncement from '../GetAnnouncement';
import SimpleTabs from '../Containers/AnnounceAssign';
import AnnouncementModal from './AnnouncementModal';
import OutlinedCard from './ClassCard';
import { ArrowBack } from '@material-ui/icons';
import AnnounceOutlinedCard from './AnnounceClassCard';
const drawerWidth = 400;
const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
  },
  entitle:{
    marginLeft:100,
    marginRight:100
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  menuButton: {
    marginRight: 36,
  },
  hide: {
    display: 'none',
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
    whiteSpace: 'nowrap',
  },
  drawerOpen: {
    width: drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerClose: {
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    overflowX: 'hidden',
    width: theme.spacing(7) + 1,
    [theme.breakpoints.up('sm')]: {
      width: theme.spacing(9) + 1,
    },
  },
  toolbar: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: theme.spacing(0, 1),
    // necessary for content to be below app bar
    ...theme.mixins.toolbar,
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
}));

const AnnouncementMiniDrawer = props =>{
  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);
  const [auth, setAuth] = React.useState(true);
  const [anchorEl, setAnchorEl] = React.useState(null);
  const opened = Boolean(anchorEl);

  useEffect(() => {
    handleGetAnnouncement()
  }, [ ])
  const navigateto = useNavigate()
  const handleLogout = () => {
      localStorage.clear();
      navigateto('/login')
  };

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };
  const handleChange = (event) => {
    setAuth(event.target.checked);
  };

  const handleMenu = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };
  const [announce,setAnnounce]=useState([])
  const classCode=localStorage.getItem('class-code')
  const userToken=localStorage.getItem('token')
  const [loadingDone,setLoadingDone]=useState(false)
  const [classData, setClassData] = useState(false)
  const handleGetAnnouncement= async ()=>{
    
    let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/announcement/get",
    {
        method:"POST",
        headers:{
            "Content-Type": "application/json",
            "Accept": "application/json",
            authorization:`Bearer ${userToken}`
        },
        body:JSON.stringify({"class":classCode})
    });
    result = await result.json();
    setAnnounce(result.data)
    {result.data?setClassData(true):setClassData(false)}
    setLoadingDone(true)
}

const history=useNavigate()
const handleBackward=()=>{
  history('/dashboard');
}
const goToProfile=()=>{
  history('/profile')
}
  return (
    <div className={classes.root}>
      <CssBaseline />
      <AppBar
        position="fixed"
        className={clsx(classes.appBar, {
          [classes.appBarShift]: open,
        })}
      >
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            edge="start"
            className={clsx(classes.menuButton, {
              [classes.hide]: open,
            })}
          >
            <MenuIcon />
          </IconButton>
          
        <Typography className={classes.entitle} variant="h4">ANNOUNCEMENTS</Typography>
          {auth && (
                        <div>
                        
                          <IconButton
                            aria-label="account of current user"
                            aria-controls="menu-appbar"
                            aria-haspopup="true"
                            onClick={handleMenu}
                            color="inherit"
                          >
                            <AccountCircle />
                          </IconButton>
                          
                          <Menu
                            id="menu-appbar"
                            anchorEl={anchorEl}
                            anchorOrigin={{
                              vertical: 'top',
                              horizontal: 'right',
                            }}
                            keepMounted
                            transformOrigin={{
                              vertical: 'top',
                              horizontal: 'right',
                            }}
                            open={opened}
                            onClose={handleClose}
                          >
                          
                        
                            <MenuItem onClick={goToProfile}>Profile</MenuItem>
                            <MenuItem onClick={handleLogout}>Logout</MenuItem>
                          </Menu>
                          
                        </div>
                      )}
        </Toolbar>
        
      </AppBar>
      <Drawer
        variant="permanent"
        className={clsx(classes.drawer, {
          [classes.drawerOpen]: open,
          [classes.drawerClose]: !open,
        })}
        classes={{
          paper: clsx({
            [classes.drawerOpen]: open,
            [classes.drawerClose]: !open,
          }),
        }}
      >
        <div className={classes.toolbar}>
          <IconButton onClick={handleDrawerClose}>
            {theme.direction === 'rtl' ? <ChevronRightIcon /> : <ChevronLeftIcon />}
          </IconButton>
        </div>
        <Divider />
        <List>
        <ListItem button onClick={handleBackward}>
        <ListItemIcon><IconButton
        color="inherit"
        aria-label="open drawer"
        
        edge="start"
      >
        <ArrowBack/>
      </IconButton>
      </ListItemIcon>
      <ListItemText label="back to Dashboard">Back To Dashboard</ListItemText>
      </ListItem>
        <ListItem button onClick={handleGetAnnouncement}>
        <ListItemIcon> <ClassIcon /></ListItemIcon>
        <ListItemText label="Create Class">Announcements</ListItemText>
      </ListItem>

            <AnnouncementModal />
          
        </List>
      </Drawer>
      <main className={classes.content}>
        <div className={classes.toolbar} />
        {classData?(loadingDone&&(announce.map(({teachername,description,classcode,timestamp},id)=>{
          
            return (
              <AnnounceOutlinedCard key={id} props={{teachername,description,classcode,timestamp,id}}/>
            );})
          
          
           )):null }
      </main>
    </div>
  );
}

export default AnnouncementMiniDrawer

// const things= {loadingDone&&(announce.map(({teachername,description,classcode},id)=>{
          
        //   return (
        //     <AnnounceOutlinedCard key={id} props={{teachername,description,classcode,id}}/>
        //   );})
        
        
        //  ) }
