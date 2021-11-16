import React, { useEffect, useState } from 'react';
import './TimeSlot.css';

const TimeSlot = (props) => 
{
    var course_code = (props.course)
    var month = (props.month) + 1
    var date = (props.day)
    var year = (props.year)
    var dayOfWeek = (props.dayOfWeek) - 1

    const[clicked, setClicked] = useState([])
    const[tutors, setTutors] = useState([])
    const[error, setError] = useState(false)
    const[booked, setBooked] = useState([])
    var bookings = [];

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
            // this.setState({tutors: result})
            })
        .catch(e => {
            console.log(e)
            setError(true)
        })
    }, [])

    // if(error) {
    //     return (
    //         <h1>No tutors available.</h1>
    //     )
    // }

    // const bookApp = (index) => {
    //     slotsarr.map((key) => {
    //         if(slotsarr[key] != index) {
    //             return(
    //             <button className="bookedBtn" disabled>
    //                 {slotsarr[key]}
    //             </button>
    //             )
    //         }
    //         else {
    //             return (
    //                 <button className="availableBtn">
    //                     {key}
    //                 </button>
    //             )
    //         }
    //     })
    // }


    return (
        <div className="slots">
        {
        tutors.map((index) => {
            return (
                <div>
                    <h5 className="tutorName">
                        {index["first_name"]} 
                    </h5>
                    <h5 className="tutorLastName">
                        {index["last_name"]}
                    </h5>
                    {
                        index.availability.map((key, index2) => {
                            if(dayOfWeek == index2){
                            return (
                                <div>
                                    {
                                    key.map((index3) => {
                                        if({index3} != slotsarr[index3]){
                                            return (
                                                <div>
                                                <h5 className="availableHours">Available hours: </h5>
                                                <button className="availableBtn">{index3}</button>
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
            // create buttons based on whether timeslot is booked

            //  slotsarr.map((key) => {
            //      return(
            //         bookings.map((key2) => {
            //             if({key2} == slotsarr[key]) {
            //                 return(
            //                 <button className="bookedBtn" disabled>
            //                     {key}
            //                 </button>
            //                 )
            //             }
            //             else {
            //                 return(
            //                     <button className="availableBtn">
            //                         {key}
            //                     </button>
            //                 )
            //             }
            //         })
            //      )})
        }
        </div>
             <div>
                <button className="buttonApp">Create Appointment</button>
            </div>
        </div>) 
};

export default TimeSlot;
