import React from 'react';
import { BrowserRouter as Router, Route, Routes} from "react-router-dom";
import './App.css';
import Home from './pages/Home';
// import SignInSignUp from './pages/SignInSignUp';
import NotFound from './pages/NotFound';
import SignUp from './components/SignUp';
import Login from './components/login';
import Dashboard from './pages/Dashboard';
import AnnouncementMiniDrawer from './components/Modules/ClassAnnouncementNavBar';
import AssignmentMiniDrawer from './components/Modules/ClassAssignmentNavBar';
import Profile from './components/Profile';
import UpdateProfile from './components/Modules/UpdateProfile';
import Createquiz from './quizapp/Createquiz';
import TakeQuiz from './quizapp/TakeQuiz';
import SignUpWithGoogle from './components/SignUpWithGoogle';
import Private from './pages/Private';
function App() {
  return (
    <div>
    <Router>
    
      <Routes>
      <Route exact path='/dashboard' element={<Private/>}>
        <Route exact path="/dashboard" element={<Dashboard/>} />
        </Route>
        <Route exact path='/profile' element={<Private/>}>
        <Route exact path="/profile" element={<Profile/>} />
        </Route>
        <Route exact path='/class/:id/assignment/createquiz' element={<Private/>}>
        <Route exact path="/class/:id/assignment/createquiz" element={<Createquiz/>} />
        </Route>
        <Route exact path='/class/:id/assignment/takequiz' element={<Private/>}>
        <Route exact path="/class/:id/assignment/takequiz" element={<TakeQuiz/>} />
        </Route>
        <Route exact path='/profile/update' element={<Private/>}>
        <Route exact path="/profile/update" element={<UpdateProfile/>} />
        </Route>
        <Route exact path="/type" element={<SignUpWithGoogle/>} />
      
        <Route exact path="/signup" element={<SignUp/>} />
        <Route exact path='/class/:id/announcement' element={<Private/>}>
        <Route exact path="/class/:id/announcement" element={<AnnouncementMiniDrawer/>} />
        </Route>
        <Route exact path='/class/:id/assignment' element={<Private/>}>
        <Route exact path="/class/:id/assignment" element={<AssignmentMiniDrawer/>} />
        </Route>
        <Route exact path="/login" element={<Login/>} />
        <Route exact path="/" element={<Home/>} />
        <Route element={<NotFound/>} />
      </Routes>
  </Router>
    </div>
  );
}

export default App;
