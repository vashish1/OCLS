import React, {useContext,useState} from 'react'
import { useHistory } from "react-router-dom";
import axios from "axios";

const GlobalContext=React.createContext();
export const useGlobal=()=>{
  return useContext(GlobalContext)
}
export const SigninupProvider = ({ children }) => {
   
  const [isLoading,setIsLoading]=useState(false)
  const userToken =localStorage.getItem('token')
  const userClasses=JSON.parse(localStorage.getItem('classes'))
  const handleGetClasses= async ()=>{
    setIsLoading(true)
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
      localStorage.setItem('classes',JSON.stringify(result.data))
      setIsLoading(false)
  }
  
 const value={
  
    handleGetClasses,
    setIsLoading,
    isLoading,
    userClasses,
    userToken
 };
  return(
    
    <GlobalContext.Provider value={value} >{children}</GlobalContext.Provider>
  )
}

