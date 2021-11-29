import React, {useState, useEffect} from 'react';
import AppointmentList from '../components/AppointmentList';
import './Background.css';



class ViewAppointment extends React.Component{
    state = {
        email : localStorage.getItem('profile'),
        fname : localStorage.getItem('fname'),
        lname : localStorage.getItem('lname'),
        profilePic : localStorage.getItem('profilePic'),
        appointments : {},
        appList : []
    }

    fetchData = async () => {     
        console.log("reached until fetchData")
        try{
            const response = await fetch(`http://localhost:8080/api/appointment/student/${this.state.email}`).then((result) => result.json());
            this.setState(() => ({ appointments: response  }));
            console.log(this.state.appointments, "is the appointment list of the user ", this.state.email);
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
                let response = await fetch(`http://localhost:8080/api/appointment/${id}`).then(result => result.json());
                return response
            }

            const List = [];
            appID.map(id => fetchData(id).then(result => {List.push(result)}))
            // let response = await fetch(`http://localhost:8080/api/appointment/${appID[1]}`).then(result => result.json());
            console.log(List,"is the appointment info")
            this.setState(() => ({ appList: List }));
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
        console.log(appointments);
        console.log(appList);
        let appointmentList = (
            <div className="Appointments-list">
                <h4>Appointments</h4>
                {/* <AppointmentList  mail={this.state.email} /> */}
            </div>
        );
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
    

}

// const ViewAppointment = (props) => {

//     const [email, setUser] =  useState(localStorage.getItem('profile'));
//     const [fname , setFname] =useState(localStorage.getItem('fname')); 
//     const [lname , setLname] =useState(localStorage.getItem('lname')); 
//     const [profilePic, setProfilePic] = useState(localStorage.getItem('profilePic'));
//     const[appointmentIDList, setappointmentIDList] = useState(null);
//     console.log(email , "is the email");

//     // let profileimgUrl = userObj.imageUrl
//     // console.log(profileimgUrl)

//     let appointmentList = (
//         <div className="Appointments-list">
//             <h4>Appointments</h4>
//             {/* <AppointmentList  mail={email} /> */}
//         </div>
//     )

//     useEffect(()=>{
//         let url = `http://localhost:8080/api/appointment/student/${email}`;
//         // let url = `http://localhost:8080/api/appointment/student/new@gmail.com`;


//         const fetchData = async () => {     
//             console.log("reached until fetchData")
//             try{
//                 const response = await fetch(url);
//                 setappointmentIDList(response);
//                 console.log(appointmentIDList, "is the api response")
//             }catch(error){
//                 console.log("error" , error);
//             }
//         };

//         fetchData();
//     } , [email])
    
//     return (
//         <div className="Appointments">
//             <div className="profile">
//                 <h3>Profile</h3>
//                 <img src= {profilePic? profilePic : "/assets/user-circle.svg"} style={{"object-fit":"cover", "border-radius":"50%"}} alt="profile" width="100" height="100" />
//                 <ol>
//                     <li>Name : {fname} {lname}</li>
//                     <li>Email : {email}</li>
//                 </ol>
//             </div>
//             <div>
//                 {appointmentList}
//             </div>
            
//         </div>
//     );
    
// };

export default ViewAppointment;


