import React, {useState} from 'react';
import './Background.css';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css'
import { useParams } from 'react-router-dom';
import TimeSlot from '../components/TimeSlot';


const ChooseTutor = () => {
    const { course } = useParams()
    const [selectedDate, setSelectedDate] = useState(null)
    
    let availableSlots;
    if(selectedDate != null){
        console.log(selectedDate);
        var d = new Date(selectedDate);
        var dday = d.getDate();
        var m = d.getMonth();
        var y = d.getFullYear();
        var dow = d.getUTCDay();
        availableSlots = (
            <div>         
                <label>{selectedDate.toDateString()}</label>
                <TimeSlot
                    course = {course}
                    month = {m}
                    day = {dday}
                    year = {y}
                    dayOfWeek = {dow}
                    />
            </div>
        )
    }



    return (
        <div>
            <h4>Tutor Selection for {course}</h4>

            <label>Choose a date</label>
            <DatePicker
                selected={selectedDate} 
                onChange={date => setSelectedDate(date)}
                minDate={new Date()}
                filterDate={date => date.getDay() != 6 && date.getDay() != 0}
            />
            <div> {availableSlots}</div>
        </div>
    );
};

export default ChooseTutor;