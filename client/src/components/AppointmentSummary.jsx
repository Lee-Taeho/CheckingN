// import React, { useState ,useEffect } from 'react';

function AppointmentSummary ({course_id, course_code, start_time, end_time, tutor_email, meeting_location, join_link}){
    return(
        <div className = "Appointments-element">
            <button className="Appointment-button"> 
                <h5>Class : {course_code}</h5>
                <ul>
                    <li>Class : {course_code}</li>
                    <li>Tutor : {tutor_email}</li>
                    <li>Meeting Type: {meeting_location}</li>
                    {join_link != "" &&
                    <li><a href={`${join_link}`}>Join</a></li>}
                    <li>Time {start_time} ~ {end_time}</li>
                </ul>
            </button>
                
        </div>
    );
}




// const AppointmentList = (props) => {
//     console.log(props.mail, "is the props.mail")
//     let email = (props.mail);
//     const[appointmentIDList, setappointmentIDList] = useState([]);
//     console.log("reached until start");

//     // const[error, setError] = useState(false)
//     useEffect(() => {
//         // fetch(`http://localhost:8080/api/appointment/student/${email}`)
//         // .then(response => response.json())
//         // .then((result) => {
//         //     setappointmentIDList(result)
//         //     console.log(appointmentIDList)
//         //     // this.setState({tutors: result})
//         //     })
//         // .catch(e => {
//         //     console.log(e)
//         //     // setError(true)
//         // })

//         let url = `http://localhost:8080/api/appointment/student/${email}`;
//         // let url = `http://localhost:8080/api/appointment/student/new@gmail.com`;


//         const fetchData = async () => {     
//             console.log("reached until fetchData")
//             try{
//                 const response = await fetch(url);
//                 console.log(response.json(), "is the response");
//                 // setappointmentIDList(response.json());
//             }catch(error){
//                 console.log("error" , error);
//             }
//         };

        


//         // console.log(appointmentIDList);
//         fetchData();

//         const IDList = appointmentIDList;
//         console.log(appointmentIDList , "is a list a ID");
//         Object.keys(appointmentIDList).map( (key, index) => console.log(key, index , " are the returned values from the IDList") );
//         appointmentIDList.map( (appointment) => {console.log(appointment, "is the appointment ID"); return {appointment}});
//         const fetchAppointmentData = async () => {
//             try{
//                 console.log("I am in fetchAppointmentData");
//                 // console.log(appointmentIDList.Array);
//                 const response = await fetch(`http://localhost:8080/api/appointment/61a4128457b68ac9c92a5d5a`);
//                 console.log(response.json() , "is the appointment information");
                
//             }catch(error){
//                 console.log("error" , error);
//             }
//         }

//         fetchAppointmentData();

//     }, [])

//     const getAppointmentIDs  = () => {
        
//     }
    
    
//     return(
//         <div>
            
//         </div>
//     );
    
    
    
    
// }





export default AppointmentSummary;