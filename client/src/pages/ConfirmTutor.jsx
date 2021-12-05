import React, {useState} from 'react';
import './Background.css';
import { useParams } from 'react-router-dom';


const ConfirmTutor = () => {
    const { course } = useParams()
    const { firstname } = useParams()
    const { lastname } = useParams()
    const { fulldate } = useParams()
    const { time } = useParams()
    const { email } = useParams()
    const { location } = useParams()

    return (
        <div className="form-signin">
            {course}
            <div>
                {firstname} {''} {lastname}
                <div>
                    {fulldate}
                    <div>{time}</div>
                    <div>{email}</div>
                    <div>{location}</div>
                </div>
            </div>
        </div>
    );
};

export default ConfirmTutor;