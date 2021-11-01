import React, { useState } from 'react';
import './DepartmentList.css'

const DepartmentList = (props) =>  {

    const [clicked, setClicked] = useState(false);

    var btnStyle = {
        backgroundColor: 'blue'
    }

    if(clicked) {
        return (
            btnStyle = {
                backgroundColor: 'red'
            }
        );
    }    

    return (
        <div className="button_container" >
            <h1>Choose a department</h1>
            <button
                className="button"
                onClick={() => setClicked(true)}
            >
            Biology / BIOL
            </button>
            <button
                className="button"
                onClick={() => setClicked(true)}
            >
            Business / BUS
            </button>
            <button
                className="button"
                onClick={() => setClicked(true)}
            >
            Chemistry / CHEM
            </button>
            <button
                className="button"
                onClick={() => setClicked(true)}
            >
            Computer Science / CS
            </button>
            <button
                className="button"
                onClick={() => setClicked(true)}
            >
            English / ENG
            </button>
            <button
                className="button"
                onClick={() => setClicked(true)}
            >
            Geology / GEOL
            </button>
            <button
                className="button"
                onClick={() => setClicked(true)}
            >
            Mathematics / MATH
            </button>
        </div>


        // <div className='button_list'>
        //     <button 
        //         style={{backgroundColor:"blue", width:100, height:100}} 
        //         onClick = {() => setClicked(true)}
        //         /> 
        // </div>
    );
};

export default DepartmentList;