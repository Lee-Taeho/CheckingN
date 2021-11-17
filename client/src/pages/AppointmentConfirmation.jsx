import React, {useState} from 'react';
import './Background.css';
import { useParams, Link, Redirect} from 'react-router-dom'

const AppointmentConfirmation = () => {
    const { course } = useParams();
    const { date } = useParams();
    const { time } = useParams();
    const { tutor } = useParams();
    const { location } = useParams();
    const [redirect, setRedirect] = useState(false);

    // current student user
    const user =  localStorage.getItem('profile');
    
    let startTime = date.toString()+ 'T' + time.toString() + 'Z';
    
    // compute End time
    var tempEndTime = new Date('2021-10-25 ' + time);
 //   console.log("temp " + tempEndTime);
    tempEndTime.setHours(tempEndTime.getHours()+1);
 //   console.log(tempEndTime)
    
    let endTime = date.toString()+ 'T' + tempEndTime.getHours().toString() + ':'
    + ("0" + tempEndTime.getMinutes()).slice(-2).toString() + ':'
    + ("0" + tempEndTime.getSeconds()).slice(-2) + 'Z';

    console.log("start " + startTime);
    console.log("end   " + endTime);
    console.log("tutor " + tutor);
    console.log("user   " + user);
    console.log("location " + location);


    const handleConfirm = async (e) => {
        console.log('confirm')
        e.preventDefault();

        var request = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                // tutor_email: "khang.d.nguyen@sjsu.edu",
                // student_email: "new@gmail.com",
                // course_code: "CS146",
                // meeting_location: "Zoom",
                // start_time: "2021-11-24T22:00:00Z",
                // end_time: "2021-11-24T23:00:00Z" 
                tutor_email: tutor,
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

    // uri: /tutoring/departments/:course/:date/:time/:tutor/:location/confirmation
    // eg. http://localhost:3000/tutoring/departments/CS174/2021-11-24/22:00:00/khang.d.nguyen@sjsu.edu/Zoom/confirmation
    return (
        <div className='confirmation-container'>
            <h4 className="title">New Appointment</h4>
            <div>
                <p><b>Your email:</b> {user}</p>
                <p><b>Tutor:</b> {tutor}</p>
                <p><b>Course:</b> {course}</p>
                <p><b>Date:</b> {date}</p>
                <p><b>Time:</b> {time}</p>
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