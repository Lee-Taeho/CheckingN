import React, {useState, useEffect} from 'react';

const Home = (props) => {
    const [user, setUser] =  useState(localStorage.getItem('profile'));

    let homeText;
    if(user ==='' || user == null) {
        homeText=(<div>You are not logged in</div>)
    }
    else{
        homeText=(<div>Hi {user} !</div>)
    }

    return (
        <div> {homeText}</div>
        
    );
};

export default Home;
