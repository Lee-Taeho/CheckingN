import React, {useState, useEffect} from 'react';
import CardItem from '../components/CardItem';

const Home = (props) => {
    const [user, setUser] =  useState(localStorage.getItem('profile'));

    let homeText;
    if(user ==='' || user == null) {
        homeText=(
            <div>
                <section className='header'>
                    <div className='header-text'>
                        <h1>The best way to find a tutor</h1>
                        <p>asdfasdfasdfa</p>
                    </div>
                </section>

                <section className='service'> 
                    <h1>Online Tutoring</h1>
                    <div className="row">   
                        <div className="col">
                            <h3>In-person Tutoring</h3>
                            <p>Our tutors are friendly and skilled in their subject fields.
                                 They undergo a thorough background check</p>
                        </div>
                        <div className="col">
                            <h3>Remote Learning</h3>
                            <p>Learn from anywhere, using your desktop, tablet, or mobile phone. 
                                Out tutoring is interactive and practical</p>
                        </div>
                        <div className="col">
                            <h3>Amazing Results</h3>
                            <p>Almost all of our students go on to achieve higher grades than
                                 they expected after completing tutoring with us</p>
                        </div>
                    </div>
                </section>
            </div>
        )
    }
    else{
        homeText=(
        <div className="form-signin">         
            <div className='cards'>
            <h1>Choose an option</h1>
                <div className='cards_container'>
                    <div className='cards_wrapper'>
                    <ul className='cards_item'>
                            <CardItem
                                img = "assets/calendar.svg"
                                title = "View Appointments"
                                path = "/tutoring/appointments"
                            />
                            <CardItem
                                img = "assets/check.svg"
                                title = "Tutoring Options"
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


