import React, { useState } from 'react';
import { useNavigate } from 'react-router';
import Style from './TakeQuiz.module.css'
var studentAnswers=[];
export default function TakeQuiz() {
	const form=JSON.parse(localStorage.getItem('quizform'))
	const assignmentId=JSON.parse(localStorage.getItem('assignmentId'))
	const [currentQuestion, setCurrentQuestion] = useState(0);
	const [showScore, setShowScore] = useState(false);
	const [score, setScore] = useState(0);
	const userToken=localStorage.getItem('token')
	const classcode=localStorage.getItem('class-code')
	const history=useNavigate()
	const handleSubmitMcq= async ()=>{
		let item={id:assignmentId,ans:studentAnswers};
        let result=await fetch("https://thawing-mountain-02190.herokuapp.com/class/mcq/sub",
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
        history(`/class/${classcode}/assignment`)
		
	}
	const handleAnswerOptionClick = (studentAnswer) => {
		studentAnswers.push(studentAnswer)
		if (studentAnswer==form.answers) {
			setScore(score + 1);
		}
		console.log(studentAnswers)
		const nextQuestion = currentQuestion + 1;
		if (nextQuestion < form.ques.length) {
			setCurrentQuestion(nextQuestion);
		}
		else{
			setShowScore(true)
			handleSubmitMcq()
		}
	};
	return (
		<div className={Style.bodydiv}>
			<div className={Style.app}>
				{showScore ? (
					<div className={Style.scoresection}>
						Thanks You. Your quiz has been submitted.
					</div>
				) : (
					<>
						<div className={Style.questionsection}>
							<div className={Style.questioncount}>
								<span>Question {currentQuestion + 1}</span>/{form.ques.length}
							</div>
							<div className={Style.questiontext}>{form.ques[currentQuestion].question}</div>
						</div>
						<div className={Style.answersection}>
							{form.ques[currentQuestion].options.map((options,i) => (
								<button className={Style.optionbutton} onClick={()=>handleAnswerOptionClick(options)}>{options}</button>
							))}
						</div>
					</>
				)}
			</div>
		</div>
	);
}

// const que={ques: [{question: "abcd", options: ["efg", "hij", "klm", "nop"]},…], answers: ["efg", "efgh"]}
