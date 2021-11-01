import React, {useState, useEffect} from 'react';

const Home = (props) => {
    const [user, setUser] =  useState(localStorage.getItem('profile'));

    let homeText;
    if(user ==='' || user == null) {
        homeText=(<div>You are not logged in</div>)
    }
    else{
        homeText=(
        <div>
            <div>Hi {user} !</div>
            <h1>Choose Options</h1>
            <ul className="Home-options">
                <li className="Home-options-elements">
                    <button className="Home-options-elements-button">
                        <input type="image" src="../../../Calendar.png" alt="Calendar" width="200" height="200" onClick={()=>{console.log("Clicked!")}} />
                    </button>
                </li>
            </ul>

            <ul className="Home-options">
                <li className="Home-options-elements">
                    <button className="Home-options-elements-button">
                        <input type="image" src="../../../Check.png" alt="Calendar" width="200" height="200" onClick={()=>{console.log("Clicked!")}} />
                    </button>
                </li>
            </ul>
        </div>
        
        )
    }

    return (
        <div> {homeText}</div>
        
    );

};

export default Home;