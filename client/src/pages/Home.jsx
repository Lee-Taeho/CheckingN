import React, {useState, useEffect} from 'react';
import CardItem from '../components/CardItem';
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faStar, faStarHalf} from "@fortawesome/free-solid-svg-icons"

const Home = (props) => {
    const [user, setUser] =  useState(localStorage.getItem('profile'));

    let homeText;
    if(user ==='' || user == null) {
        homeText=(
            <div>
                <section className='header'>
                    <div className='header-text'>
                        <h1 style={{fontSize: '70px'}}>Get Your Grades Up!</h1>
                        <h2>The best way to find a tutor</h2>
                    </div>
                </section>

                <section className='service'> 
                    <h1>Online Tutoring</h1>
                    <div className='row'>   
                        <div className='service-col'>
                            <h3>In-person Tutoring</h3>
                            <p>Our tutors are friendly and skilled in their subject fields.
                                 They undergo a thorough background check.</p>
                        </div>
                        <div className='service-col'>
                            <h3>Remote Learning</h3>
                            <p>Learn from anywhere, using your desktop, tablet, or mobile phone. 
                                Out tutoring is interactive and practical.</p>
                        </div>
                        <div className='service-col'>
                            <h3>Amazing Results</h3>
                            <p>Almost all of our students go on to achieve higher grades than
                                 they expected after completing tutoring with us.</p>
                        </div>
                    </div>
                </section>

                <section className='college'>
                    <h1>Our Partners</h1>
                    <div className='row'>
                        <div className='college-col'>
                            <img src='../../assets/sjsu.jpeg' />
                            <div class='layer'>
                                    <h3>SJSU</h3>
                            </div>
                        </div>

                        <div className='college-col'>
                            <img src='../../assets/sfsu.jpeg' />
                            <div class='layer'>
                                    <h3>SFSU</h3>
                            </div>
                        </div>

                        <div className='college-col'>
                            <img src='../../assets/deanza.jpeg' />
                            <div class='layer'>
                                    <h3>De Anza</h3>
                            </div>
                        </div>
                    </div>
                </section>

                <section className='reviews'>
                    <h1>What Our Users Say</h1>

                    <div className='row'>
                        <div className='reviews-col'>
                            <img src='../../assets/user1.jpeg'></img>
                            <div>
                            <p>I'm very grateful to all the tutors and CheckingN for the help they gave
                                    me in the run up to my final exam. I went into the exam with greater confidence
                                    and really felt like the tutoring session helped a lot.</p>
                                <h3>Nhien Lam</h3>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>

                            </div>
                        </div>

                        <div className='reviews-col'>
                            <img src='../../assets/user1.jpeg'></img>
                            <div>
                                <p>My thanks go out to CheckingN for the support. The technical we have had has been
                                    amazing. CheckingN makes it easier to schedule a tutoring session and get help instantly.
                                </p>
                                <h3>Ricky Lam</h3>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>
                                <FontAwesomeIcon icon={faStar} style={{color: "#f2ea00"}}/>

                            </div>
                        </div>

                    </div>
                </section>

                <section className='call-to-action'>
                    <h1>Schedule a tutoring session now! </h1>
                    <br/>
                    <a href="/register">JOIN US</a>
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


