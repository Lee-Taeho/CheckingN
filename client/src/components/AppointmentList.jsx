import React, { useState ,useEffect } from 'react';

const AppointmentList = (props) => {
    console.log(props.mail, "is the props.mail")
    let email = (props.mail);
    const[appointmentIDList, setappointmentIDList] = useState([]);
    console.log("reached until start");

    // const[error, setError] = useState(false)
    useEffect(() => {
        // fetch(`http://localhost:8080/api/appointment/student/${email}`)
        // .then(response => response.json())
        // .then((result) => {
        //     setappointmentIDList(result)
        //     console.log(appointmentIDList)
        //     // this.setState({tutors: result})
        //     })
        // .catch(e => {
        //     console.log(e)
        //     // setError(true)
        // })

        let url = `http://localhost:8080/api/appointment/student/${email}`;
        // let url = `http://localhost:8080/api/appointment/student/new@gmail.com`;


        const fetchData = async () => {     
            console.log("reached until fetchData")
            try{
                const response = await fetch(url);
                console.log(response, "is the response");
                setappointmentIDList(response.json());
            }catch(error){
                console.log("error" , error);
            }
        };

        const IDList = appointmentIDList;
        console.log(IDList , "is a list a ID");
        Object.keys(IDList).map( (key, index) => console.log(key, index , " are the returned values from the IDList") );


        // console.log(appointmentIDList);
        fetchData();
        // appointmentIDList.map( (appointment) => {console.log(appointment, "is the appointment ID"); return {appointment}});
        const fetchAppointmentData = async () => {
            try{
                console.log("I am in fetchAppointmentData");
                // console.log(appointmentIDList.Array);
                const response = await fetch(`http://localhost:8080/api/appointment/6196d97fad77a7bcb19870ea`);
                console.log(response.json() , "is the appointment information");
                
            }catch(error){
                console.log("error" , error);
            }
        }

        fetchAppointmentData();

    }, [])

    const getAppointmentIDs  = () => {
        
    }
    
    
    return(
        <div>
            
        </div>
    );
    
    
    
    
}





export default AppointmentList;