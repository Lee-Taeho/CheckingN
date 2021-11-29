import React, {useState, useEffect} from 'react';
import AppointmentList from '../components/AppointmentList';
import './Background.css';

const ViewAppointment = (props) => {

    const [email, setUser] =  useState(localStorage.getItem('profile'));
    const [fname , setFname] =useState(localStorage.getItem('fname')); 
    const [lname , setLname] =useState(localStorage.getItem('lname')); 
    const [profilePic, setProfilePic] = useState(localStorage.getItem('profilePic'));
    console.log(email , "is the email");

    // let profileimgUrl = userObj.imageUrl
    // console.log(profileimgUrl)

    let appointmentList = (
        <div className="Appointments-list">
            <h4>Appointments</h4>
            <AppointmentList  mail={email} />
        </div>
    )

    
    return (
        <div className="Appointments">
            <div className="profile">
                <h3>Profile</h3>
                <img src= {profilePic? profilePic : "/assets/user-circle.svg"} style={{"object-fit":"cover", "border-radius":"50%"}} alt="profile" width="100" height="100" />
                <ol>
                    <li>Name : {fname} {lname}</li>
                    <li>Email : {email}</li>
                </ol>
            </div>
            <div>
                {appointmentList}
            </div>
            
        </div>
    );
    
};

export default ViewAppointment;