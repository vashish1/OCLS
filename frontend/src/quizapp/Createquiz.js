import React, { useState } from 'react'
import { Grid,Paper, Avatar, TextField, Button } from '@material-ui/core'
import { QuestionAnswer } from '@material-ui/icons'
import moment from 'moment'
import { useNavigate } from 'react-router'

var new_list = []
var answer=[]
const Createquiz = () => {
    const paperStyle={
        padding :20,
        width:800, 
        margin:"0 auto"
    }
    const avatarStyle={backgroundColor:'#1bbd7e'}
    const btnstyle={margin:'8px 0'}
    const textField={margin:'10px auto'}
    const submitQuiz={backgroundColor:'#116530',color:'#ffffff'}
    
    
    let mcqvalues={
        ques:new_list,
        answers:answer
    }
    // const [qa, setQa] = useState(mcqvalues)
    
    const [question,setQuestion]=useState("")
    
    let options=[]
    const [noOfQuestions,setNoOfQuestions]=useState(0)
    const [optionA, setOptionA] = useState("")
    const [optionB, setOptionB] = useState("")
    const [optionC, setOptionC] = useState("")
    const [optionD, setOptionD] = useState("")
    const [correctAnswer,setCorrectAnswer]=useState("")
    const [desc, setDesc] = useState("")
    const classcode=localStorage.getItem('class-code')
    const yourDate = new Date()
const currentTime = moment(yourDate).format('DD-MMM-YY');

    const handleQuestion=(e)=>{
        e.preventDefault()
        setQuestion(e.target.value)
    }
    const handleOptionA=(e)=>{
        e.preventDefault()
        setOptionA(e.target.value)
       
    }
    const handleOptionB=(e)=>{
        e.preventDefault()
        setOptionB(e.target.value)
       
    }
    const handleOptionC=(e)=>{
        e.preventDefault()
        setOptionC(e.target.value)
       
        
    }
    const handleOptionD=(e)=>{
        e.preventDefault()
        setOptionD(e.target.value)
        
    }
    const handleInputQuestions=(e)=>{
        e.preventDefault()
        setNoOfQuestions(e.target.value)
    }
    const handleAddQuestion=(e)=>{
        e.preventDefault()
        options.push(optionA)
        options.push(optionB)
        options.push(optionC)
        options.push(optionD)
        answer.push(correctAnswer)
        // let newList=JSON.parse(localStorage.getItem('qlist'))
        let item={question:question,options:options}
        
        new_list.push(item)
        

    }
    
    const handleCorrectAnswer=(e)=>{
        setCorrectAnswer(e.target.value)
        
    }
    const handleDesc=(e)=>{
        setDesc(e.target.value)
    }
    const history=useNavigate()
    const userToken=localStorage.getItem('token')
    const handleCreateQuiz=async (e)=>{
        let item={mcq:mcqvalues,code:classcode,description:desc,date:currentTime};
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/mcq/add",
        {
            method:"POST",
            headers:{
                "Content-Type": "application/json",
                "Accept": "application/json",
                authorization:`Bearer ${userToken}`
            },
            body:JSON.stringify(item)
        });
        result = await result.json();
        history('/dashboard')
    }
    
    return (
        <div>
        
        <TextField style={textField} name="description" id="outlined-basic" onChange={handleDesc} variant="outlined" label='Description' placeholder='Enter Description of Quiz' fullWidth required/>
        Enter no. of questions: <TextField style={textField} name="no. of questions" id="outlined-basic" onChange={handleInputQuestions} variant="outlined" label='No. of questions' placeholder='Enter No. of questions' fullWidth required/>
        
        {Array.apply(null, { length: noOfQuestions }).map((e, i) =>
        (
                <Grid key={i} >
    <Paper  style={paperStyle}>
        <Grid align='center'>
             <Avatar style={avatarStyle}><QuestionAnswer/></Avatar>
            <h2>Question {i+1}</h2>
        </Grid>
        <TextField style={textField} name="question" id="outlined-basic" onChange={handleQuestion} variant="outlined" label='Enter Question' placeholder='Enter Question' fullWidth required/>
       
        <TextField style={textField} name="A" id="outlined-basic" onChange={handleOptionA} label='Option A' placeholder='Enter OptionD' variant="outlined" fullWidth required/>
        <TextField style={textField} name="B" id="outlined-basic" onChange={handleOptionB} variant="outlined" label='Option B' placeholder='Enter OptionB' fullWidth required/>
        <TextField style={textField} name="C" id="outlined-basic" onChange={handleOptionC} variant="outlined" label='Option C' placeholder='Enter OptionC'  fullWidth required/>
        <TextField style={textField} name="D" id="outlined-basic" onChange={handleOptionD} variant="outlined" label='Option D' placeholder='Enter OptionD' fullWidth required/>
        <TextField style={textField} name="correct" id="outlined-basic" onChange={handleCorrectAnswer} variant="outlined" label='Enter Correct Answer' placeholder='Enter OptionD' fullWidth required/>
        
        <Button type='submit' onClick={handleAddQuestion} color='primary' variant="contained" style={btnstyle} fullWidth>Add question</Button>
        
        
    </Paper>
</Grid>
        )
        )}
        <Button style={submitQuiz} type='submit' onClick={handleCreateQuiz} color='primary' variant="contained" style={btnstyle} fullWidth>CreateQuiz</Button>
        

        </div>
    )
}

export default Createquiz


// const things= {
//     "mcq":{
//         "ques":[
//             {
//                 "question":"usdhgbnchjv",
//                 "options":["a","b","c"]
//             },
//             {
//                 "question":"usdhgbnchjv",
//                 "options":["a","b","c"]
//             }
//         ],
//         "answers":["a","c"]
//     },
//     "code":"e37e5dshi",
//     "description":"kjuyhtgfvbnj",
//     "date":"12-Dec-21" 
// }


