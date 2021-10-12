import React  from 'react';

const Home = (props) => {
    return (
        <div>
            {props.first_name ? 'Hi ' + props.first_name : 'You are not logged in'}
        </div>
        
    );
};

export default Home;
