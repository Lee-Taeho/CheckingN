import React, {useState} from 'react';
import './Background.css';
import { useParams, Link, Redirect} from 'react-router-dom'

const AppointmentConfirmation = () => {
    const { course } = useParams();
    const { fulldate } = useParams();
    const { time } = useParams();
    const { email } = useParams();
    const { firstname } = useParams();
    const { lastname } = useParams();
    const { location } = useParams();
    const [redirect, setRedirect] = useState(false);
    
    // formats the time
    var formattedTime
    if(time.toString().length < 2){
        formattedTime = '0' + time +':00:00';
    }
    else{
        formattedTime = time +':00:00';
    }

    // current student user
    const user =  localStorage.getItem('profile');

    // start time PST timezone format
    var startTime = fulldate.toString()+ 'T' + formattedTime.toString()+ "-08:00";
    
    // compute End time
    var tempEndTime = new Date('2021-10-25 ' + formattedTime);
    tempEndTime.setHours(tempEndTime.getHours()+1);
    var eTime = ("0" + tempEndTime.getHours()).slice(-2).toString() + ':'
    + ("0" + tempEndTime.getMinutes()).slice(-2).toString() + ':'
    + ("0" + tempEndTime.getSeconds()).slice(-2);

    // end time PST timezone format
    var endTime = fulldate.toString()+ 'T' + ("0" + tempEndTime.getHours()).slice(-2).toString() + ':'
    + ("0" + tempEndTime.getMinutes()).slice(-2).toString() + ':'
    + ("0" + tempEndTime.getSeconds()).slice(-2)
    + "-08:00";

    console.log("tutor_email " + email);
    console.log("student_email   " + user);
    console.log("course_code   " + course);
    console.log("meeting_location " + location);
    console.log("start_time " + startTime);
    console.log("end_time   " + endTime);
    
    // fetch api to create an appointment in db when user clicks confirm
    const handleConfirm = async (e) => {
        console.log('confirm')
        e.preventDefault();

        var request = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                tutor_email: email,
                student_email: user,
                course_code: course,
                meeting_location: location,
                start_time: startTime,
                end_time: endTime           
            })
        }

        // send appointment info to backend to create appointment in db
        const response = await fetch('http://localhost:8080/api/appointment', request);   

         // redirect to home after user login
         setRedirect(true);

    }

      // redirect to Home
      if(redirect)
      { 
          return <Redirect to="/"/>
      }

    return (
        <div className='confirmation-container'>
            <h4 className="title">New Appointment</h4>
            <div>
                <p><b>Your email:</b> {user}</p>
                <p><b>Tutor:</b> {firstname} {lastname}</p>
                <p><b>Course:</b> {course}</p>
                <p><b>Date:</b> {fulldate}</p>
                <p><b>Time:</b> {formattedTime} - {eTime}</p>
                <p><b>Location:</b> {location}</p>
            </div>
            <div className='btn-container'>
                <button className='btn-primary' onClick={handleConfirm}>Confirm</button>
                <Link to={"/tutoring/departments/" + course}>
                    <button className='btn-primary'>Cancel</button>
                </Link>
            </div>
        </div>
    );
};

export default AppointmentConfirmation;