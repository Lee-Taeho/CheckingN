import React, {useState, useEffect} from 'react';
import AppointmentSummary from '../components/AppointmentSummary';
import './Background.css';



class ViewAppointment extends React.Component{
    state = {
        email : localStorage.getItem('profile'),
        fname : localStorage.getItem('fname'),
        lname : localStorage.getItem('lname'),
        profilePic : localStorage.getItem('profilePic'),
        appointments : {"Not Yet" : null},
        appList : [],
        // isLoading : true
    }

    fetchData = async () => {     
     //   console.log("reached until fetchData")
        try{
            const response = await fetch(`http://localhost:8080/api/appointment/student/${this.state.email}`).then((result) => result.json());
            this.setState(() => ({ appointments: response  }));
            // console.log(this.state.appointments, "is the appointment list of the user ", this.state.email);
            return this.state.appointments;
        }catch(error){
            console.log("error" , error);
        }
    };

    fetchAppointmentData = async (appID) => {
        try{
            // console.log("I am in fetchAppointmentData with appointmentList ", appID[1]);
            // console.log(appointmentIDList.Array);

            const fetchData = async (id) =>
            {
                try{
                    const response = await fetch(`http://localhost:8080/api/appointment/${id}`).then((result) => result.json());
                    return response;
                }catch(error){
                    console.log("error" , error);
                }
            }


            // console.log(fetchData(appID[0]).then(result => result), "is appointment[0]");
         //   console.log(fetchData(appID[0]), "is appointment[0]")
            const list = [];
            appID.map(id =>  fetchData(id).then((result) => {list.push(result); this.setState(() => ({ appList: list  }));}))
            // let response = await fetch(`http://localhost:8080/api/appointment/${appID[1]}`).then(result => result.json());
            // console.log(List,"is the appointment info")
            this.setState(() => ({ appList: list  }));
            // this.setState({appList : List });
            list.map((promise) => {promise.then((result) => console.log(result , "is the mapping"))})
         //   console.log(this.state.appList[0], "is element 0")
            // this.setState(() => ({ appointmentList: appList }));


            
            
        }catch(error){
            console.log("error" , error);
        }
    };

    componentDidMount() {
        this.fetchData(this.state.email)
        .then( id => {this.fetchAppointmentData(id)});
    }

    render() {
        const {email , fname , lname , profilePic, appointments, appList} = this.state;
        // const {course_code, start_time, end_time, tutor_email, meeting_location} = appList[0];
      //  console.log(appointments);
        // if(!this.state.isLoading){
      //      console.log( appList, "is appList, and its length is " , appList.length);
        // }

        //just for testing!!! delete this code after testing.
        // let AppList = [
        //     {
        //         "tutor_email": "ekaterina.kazantseva@sjsu.edu",
        //         "student_email": "new@gmail.com",
        //         "course_code": "CS146",
        //         "meeting_location": "Zoom",
        //         "start_time": "2021-12-06T17:00:00Z",
        //         "end_time": "2021-12-06T18:00:00Z",
        //         "join_link": "https://us04web.zoom.us/j/7400707319?pwd=TVZsWWhhM2tTa2pNWGxzZG5FWnE3UT09",
        //         "start_link": "https://us04web.zoom.us/s/7400707319?zak=eyJ0eXAiOiJKV1QiLCJzdiI6IjAwMDAwMSIsInptX3NrbSI6InptX28ybSIsImFsZyI6IkhTMjU2In0.eyJhdWQiOiJjbGllbnRzbSIsInVpZCI6InFNNmttRkVNUnlHb1hDZ0NxQU92N3ciLCJpc3MiOiJ3ZWIiLCJzdHkiOjEsIndjZCI6InVzMDQiLCJjbHQiOjAsIm1udW0iOiI3NDAwNzA3MzE5Iiwic3RrIjoiS1lVeWFTLTJRaF80QllqS05yRXJCbnFsTExUMTNFMFAyeHMxejRXSnIyVS5BRy5jUzBDeE9wOXpGTVBjQ2o0dHQ2M2N1SndGbi1rWVRDbFAxNkI3T0s2VlBudVRYbkZabDNtWndYaVJNVUdiVGhHRWE5M244ZEdWN0QtbGh4Si5xN29xOWlUaE5CNjJxZlg4UnAyVTRnLmJCX295b0Z1Xy1FcEJmMEkiLCJleHAiOjE2MzgxNTAxODYsImlhdCI6MTYzODE0Mjk4NiwiYWlkIjoiWlZiakpfNHFSRXU1SGlobEhTZWN4dyIsImNpZCI6IiJ9.G85c8OVqbcJDsME7TEsRNaCtXbFoFhMruBIBYNbKZF4"
        //     },
        //     {
        //         "tutor_email": "ekaterina.kazantseva@sjsu.edu",
        //         "student_email": "new@gmail.com",
        //         "course_code": "CS146",
        //         "meeting_location": "In-person",
        //         "start_time": "2021-11-29T17:00:00Z",
        //         "end_time": "2021-11-29T18:00:00Z",
        //         "join_link": "",
        //         "start_link": ""
        //     }
        // ];

        let appointmentSummary;
        //AppList -> appList after testing!!!
        if(appList){
      //      console.log("appList is ", appList);
            appointmentSummary = (
                <div className="Appointments-list">
                    <h4>Appointments</h4>
                    {/* {this.fetchAppointmentData} */}
                    {
                        appList.map( (appointment) => {
                            return (
                                <AppointmentSummary course_code={appointment.course_code} start_time={appointment.start_time} end_time={appointment.end_time} tutor_email={appointment.tutor_email} meeting_location={appointment.meeting_location} />
                                
                            )
                        })
                    }
                </div>
            )
        }else {
            appointmentSummary = (<h4>no apppointments</h4>)
        }
        

        appList.map(appointment => {console.log(appointment, "is the appointment element"); return appointment})
        // let appointmentSummary = () => (
            
        //     <div className="Appointments-list">
        //         <h4>Appointments</h4>
        //         {/* {this.fetchAppointmentData} */}
        //         {
        //             this.state.appList.map( (appointment) => {
        //                 return (
        //                     <AppointmentSummary course_code={appointment.course_code} start_time={appointment.start_time} end_time={appointment.end_time} tutor_email={appointment.tutor_email} meeting_location={appointment.meeting_location} />
        //                 )
        //             })
        //         }
                
        //     </div>
        // );
       
        return (
            <div className="form-signin">
                <div className="profile">
                    <h3>Profile</h3>
                    <img src= {profilePic? profilePic : "/assets/user-circle.svg"} style={{"object-fit":"cover", "border-radius":"50%"}} alt="profile" width="100" height="100" />
                    <ol>
                        <li>Name : {fname} {lname}</li>
                        <li>Email : {email}</li>
                    </ol>
                </div>
                <div className = "apppointment-lists">
                    {appointmentSummary}
                </div>
                
            </div>
        );
    };
    

}

export default ViewAppointment;


