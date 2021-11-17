<<<<<<< HEAD
import React from 'react';
import './Background.css';
import { useParams } from 'react-router-dom'

const ChooseTutor = () => {
    const { course } = useParams()

    return (
        <div>Choose Tutor for {course} </div>
=======
import React, {useState} from 'react';
import './Background.css';
import './ChooseTutor.css';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css'
import { useParams } from 'react-router-dom';
import TimeSlot from '../components/TimeSlot';


const ChooseTutor = () => {
    const { course } = useParams()
    const [selectedDate, setSelectedDate] = useState(null)
    const [location, setLocation] = useState();

    var d = new Date(selectedDate);
    var dday = d.getDate();
    var m = d.getMonth();
    var y = d.getFullYear();
    var dow = d.getUTCDay();

    let availableSlots;
    if(selectedDate != null){
        console.log(selectedDate);
        availableSlots = (
            <div>
                    <div>
                    <label className="locationOptn">Location for tutoring session?</label>
                    </div>
                    <ul className="btnLocationWrap">
                    <button className ="btnLocation"
                        onClick={() => setLocation("online")}> Online
                    </button>
                    <button className ="btnLocation"
                        onClick={() => setLocation("in-person")}> In-Person
                    </button>
                </ul>         
            </div>
        )
    }

    if(location) {
        return (
            <div>
            <label className ="timeSlotDate">{selectedDate.toDateString()}</label>
            <TimeSlot
                course = {course}
                month = {m}
                day = {dday}
                year = {y}
                dayOfWeek = {dow}
                location = {location}
                />
                </div>
        )
    }


    return (
        <div>
            <h4 className="tutorSelectTitle">Tutor Selection for {course}</h4>

            <label className="dateOptn">Choose a date</label>
            <DatePicker
                selected={selectedDate} 
                onChange={date => setSelectedDate(date)}
                minDate={new Date()}
                filterDate={date => date.getDay() != 6 && date.getDay() != 0}
            />
            <div> {availableSlots}</div>
        </div>
>>>>>>> 76102fc9a373a9cabfdad6ee815a9bac5206baf2
    );
};

export default ChooseTutor;