import React from 'react';
import CardItem from './CardItem';
import './Cards.css';

const Cards = props => {
    return (
        <div className='cards'>
            <h1>Choose an option</h1>
            <div className='cards_container'>
                <div className='cards_wrapper'>
                <ul className='cards_item'>
                        <CardItem
                            img = "assets/calendar.svg"
                            title = "Booking Appointments"
                            path = '/choosedepartment'
                        />
                        <CardItem
                            img = "assets/message-circle.svg"
                            title = "Drop-in Tutoring"
                        />
                        <CardItem
                            img = "assets/mail.svg"
                            title = "Asynchronous Tutoring"
                        />
                    </ul>
                    </div>
            </div>
        </div>
    );
}

export default Cards;