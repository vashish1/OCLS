import React,{useEffect, useState} from 'react';

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
import GetClass from '../getClasses';
import CreateClass from '../createClass';
import { useGlobal } from '../../context/SignupContext';
import OutlinedCard from './ClassCard';
const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
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
  entitle:{
    marginLeft:100,
    marginRight:100
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
}));

const MiniDrawer = props =>{
  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);
  const [auth, setAuth] = React.useState(true);
  const [anchorEl, setAnchorEl] = React.useState(null);
  const [userClass,setUserClass]=useState([])
  const opened = Boolean(anchorEl);
  // const { isLoading,setIsLoading } = useGlobal();
  // const allClass=!isLoading?JSON.parse(localStorage.getItem('classes')):null
  // const allClass=JSON.parse(localStorage.getItem('classes'))
  const [toggle,setToggle]=useState(true)
  const navigateto = useNavigate()
  const handleLogout = () => {
      localStorage.clear();
      navigateto('/login')
  };
  useEffect(() => {
    handleGetClasses()
  }, [])

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
  const userToken =localStorage.getItem('token')
  const handleGetClasses= async ()=>{
      setToggle(true)
      let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/get",
      {
          method:"GET",
          headers:{
              "Content-Type": "application/json",
              "Accept": "application/json",
              authorization: `Bearer ${userToken}`,
          },
      });
      result = await result.json();
      setUserClass(result.data)
      // localStorage.setItem('classes',JSON.stringify(result.data))
      setToggle(false)
  }
const handleProfile=()=>{
  navigateto('/profile')
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
          <Typography className={classes.entitle} variant="h4">Dashboard</Typography>
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
                            <MenuItem onClick={handleProfile}>Profile</MenuItem>
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
          
        <List>
        <ListItem button onClick={handleGetClasses}>
              <ListItemIcon> <ClassIcon /></ListItemIcon>
              <ListItemText label="Create Class">Classes</ListItemText>
            </ListItem>
        </List>
            
            <CreateClass/>
            
          
        </List>
      </Drawer>
      <main className={classes.content}>
        <div className={classes.toolbar} />
        {userClass.map(({subject,code},id)=>{
          
          return (
            <OutlinedCard key={id} props={{subject,code,id}}/>
          );})
        
        
          }
        
      </main>
    </div>
  );
}

export default MiniDrawer
// const things={userClass.map(({subject,code},id)=>{
          
            // return (
            //   <div className="lead-list__row" key={id}>
            //     <p>
            //       {id + 1}. {subject}
            //     </p>
            //     <p>code {code}</p>
            //   </div>
            // );})
          
          
            // }
