import React, { useEffect, useState } from 'react';
import './DepartmentList.css'
import { Link } from 'react-router-dom';


const DepartmentList = (props) =>  {

    const [clicked, setClicked] = useState([]);
    const [courses, setCourses] = useState()
    const [isLoaded, setLoaded] = useState(false)

    useEffect(() => {
        fetch("http://localhost:8080/api/courses_by_departments")
            .then(res => res.json())
            .then((result) => {
                setCourses(result)
                setLoaded(true)
                const initialState = Object.keys(result).map((key, index) => {
                    return false
                })
                setClicked(initialState)
            })
    }, [])

    const handlePress = (state, idx) => {
        const newArr = [...state]
        newArr[idx] = !newArr[idx]
        setClicked(newArr)
    }

    if (isLoaded) {
        return (
            <div className="button_container">
            {
            Object.keys(courses).map((key, index) => {
                return (
                    <div>
                    <button
                        className="button"
                        onClick={() => handlePress(clicked, index)}>
                            {key}
                    </button>
                    {clicked[index] && courses[key].map((course) => {
                        return (
                        <Link to={`/tutoring/departments/${course["class_code"]}`}>
                        <button
                            className="button"
                            >
                                {course["class_code"]} {course["name"]}
                        </button></Link>)
                    })}
                    </div>
                )
            })}
            </div>) 
    } else if (!isLoaded) {
        return <div>Loading...</div>
    }
/*
        
        // <div className='button_list'>
        //     <button 
        //         style={{backgroundColor:"blue", width:100, height:100}} 
        //         onClick = {() => setClicked(true)}
        //         /> 
        // </div>
    ); */
};

export default DepartmentList;