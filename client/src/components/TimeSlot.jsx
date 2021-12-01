import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import './TimeSlot.css';

const TimeSlot = (props) => 
{
    var course_code = (props.course)
    var month = (props.month) + 1
    var date = (props.day)
    var year = (props.year)
    var dayOfWeek = (props.dayOfWeek) - 1
    var fulldate = `${year}-${month}-${date}`
    var location = (props.location)

    const[tutors, setTutors] = useState([])
    const[error, setError] = useState(false)

    const jsonData = { "slots" : {
        "slot1": "9:00 AM",
        "slot2": "10:00 AM",
        "slot3": "11:00 AM",
        "slot4": "12:00 AM",
        "slot5": "1:00 PM",
        "slot6": "2:00 PM",
        "slot7": "3:00 PM",
        "slot8": "4:00 PM",
        "slot9": "5:00 PM",
        "slot10": "6:00 PM",
        "slot11": "7:00 PM",
        "slot12": "8:00 PM",
        "slot13": "9:00 PM"
     }
    }
    const slots = jsonData.slots
    const slotsarr = Object.keys(slots).map( function(k) {
        return (  
            <button className="buttonSlot">{slots[k]} </button>
        )
    });

    useEffect(() => {
        fetch(`http://localhost:8080/api/${course_code}/tutors/${year}/${month}/${date}`)
        .then(response => response.json())
        .then((result) => {
            setTutors(result)
            })
        .catch(e => {
            console.log(e)
            setError(true)
        })
    }, [])

    if(error) {
        return (
            <h4>No tutors available this day.</h4>
        )
    }

    const handlePress = (state, idx) => {
        
    }

    if(tutors) {
    return (
        
        <div className="slots">
        <h5 className="availableHours">Available hours for: </h5>
        {
        tutors.map((tutor) => {
            return (
                <div>
                    {/* display tutor name */}
                    <h5 className="tutorName">
                        {tutor["first_name"]} {''} {tutor["last_name"]}
                    </h5>
                    {
                        // parse through availability to determine day of week
                        tutor.availability.map((key, index2) => {
                            if(dayOfWeek == index2){
                            return (
                                <div>
                                    {
                                    // find available time slots
                                    key.map((time) => {
                                        if({time} != slotsarr[time]){
                                            return (
                                                // display available time slot
                                                <div className="availableBtnContainer">
                                                <Link to={`/tutoring/departments/${course_code}/${tutor["first_name"]}/${tutor["last_name"]}/${fulldate}/${time}/${tutor["email"]}/${location}`}>
                                                <button className="availableBtn" >
                                                    {time}:00
                                                    </button></Link>
                                                </div>
                                            )
                                        }
                                    })
                                    }
                                </div>
                            )  
                            }

                        }
                        )}
                </div>
            )
        })}
     
        <div>
        {
            // create buttons based on whether timeslot is booked => currently not working

            //  slotsarr.map((key) => {
            //      return(
            //         bookings.map((key2) => {
            //             if({key2} == slotsarr[key]) {
            //                 return(
            //                 <button className="bookedBtn" disabled>
            //                     {slots[key]}
            //                 </button>
            //                 )
            //             }
            //             else {
            //                 return(
            //                     <button className="availableBtn">
            //                         {slots[key]}
            //                     </button>
            //                 )
            //             }
            //         })
            //      )})
        }

        {/* create appointment button */}

        </div>
             {/* <div>
                <button className="buttonApp">Create Appointment</button>
            </div> */}
        </div>
        
        ) 
        
        } else{
            return <h4 className = "noTutor">No tutors available this day</h4>
        }
};

export default TimeSlot;
