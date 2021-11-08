import React from 'react';
import './Background.css';
import { useParams } from 'react-router-dom'

const ChooseTutor = () => {
    const { course } = useParams()

    return (
        <div>Choose Tutor for {course} </div>
    );
};

export default ChooseTutor;