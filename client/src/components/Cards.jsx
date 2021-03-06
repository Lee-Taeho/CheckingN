import React from 'react';
import CardItem from './CardItem';
import './Cards.css';

const Cards = props => {
    return (
        <div className="form-signin">
        <div className='cards'>
            <h1>Choose Tutoring Option</h1>
            <div className='cards_container'>
                <div className='cards_wrapper'>
                <ul className='cards_item'>
                        <CardItem
                            img = "assets/calendar.svg"
                            title = "Booking Appointments"
                            path = '/tutoring/departments'
                        />
                        <CardItem
                            img = "assets/message-circle.svg"
                            title = "Drop-in Tutoring"
                            path = "/tutoring"
                        />
                        <CardItem
                            img = "assets/mail.svg"
                            title = "Asynchronous Tutoring"
                            path = "/tutoring"
                        />
                    </ul>
                    </div>
            </div>
        </div>
        </div>
    );
}

export default Cards;