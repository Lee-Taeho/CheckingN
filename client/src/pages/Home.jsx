import React, {useState, useEffect} from 'react';
import CardItem from '../components/CardItem';

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
            
            <div className='cards'>
            <h1>Choose an option</h1>
                <div className='cards_container'>
                    <div className='cards_wrapper'>
                    <ul className='cards_item'>
                            <CardItem
                                img = "assets/calendar.svg"
                                title = "View Appointments & Messages"
                                path = "/tutoring"
                            />
                            <CardItem
                                img = "assets/check.svg"
                                title = "Make Appointments & Tutoring Options"
                                path = "/tutoring"
                            />
                    </ul>
                </div>
            </div>
        </div>
        </div>
        
        )
    }

    return (
        <div> {homeText}</div> 
    );

};

export default Home;


