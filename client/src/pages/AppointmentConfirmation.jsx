import React from 'react';
import './Background.css';
import { useParams, Link, useHistory } from 'react-router-dom'

const AppointmentConfirmation = () => {
    const { course } = useParams();
    const { date } = useParams();
    const { time } = useParams();
    const { tutor } = useParams();
    const { location } = useParams();

    let history = useHistory();

    // uri: /tutoring/departments/:course/:date/:time/:tutor/:location/confirmation
    // eg. http://localhost:3000/tutoring/departments/CS174/12-23-2021/12:22:00/Nhien%20Lam/online/confirmation
    return (
        <div>
            <h4>New Appointment</h4>
            <div>
                <p>Tutor: {tutor}</p>
                <p>Course: {course}</p>
                <p>Date: {date}</p>
                <p>Time: {time}</p>
                <p>Location: {location}</p>
            </div>
            <button>Confirm</button>
            <button onClick={() => {
                history.push("/tutoring/departments/" + course + "/" + date + "/" + time + "/" + tutor + "/" + location);
            }}>Cancel</button>
            <Link to={"/tutoring/departments/" + course}>
                <button>Cancel</button>
            </Link>
        </div>
    );
};

export default AppointmentConfirmation;