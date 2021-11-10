import './App.css';
import Login from './pages/Login';
import Nav from './components/Nav';
import {BrowserRouter, Route} from "react-router-dom";
import Home from './pages/Home';
import Register from './pages/Register';
import React, {useEffect, useState} from 'react';
import GoogleLogin from 'react-google-login';
import TutoringOptions from './pages/TutoringOptions';
import ChooseDepartment from './pages/ChooseDepartment';
import ChooseTutor from './pages/ChooseTutor';

function App() {
  const [first_name, setFirstName] = useState('');
    
  // useEffect(() => {
  //     (
  //         // which api is user api to get user data 
  //         async () => {
  //             const response = await fetch('http://localhost:8080/api/****', {
  //                 headers: {'Content-Type': 'application/json'},
  //                 //credentials: 'include',
  //             })

  //             const content = await response.json();

  //             setFirstName(content.first_name);
  //         }
  //     )();
  // });

  return ( 
    <div className="App">
      <BrowserRouter>
        <Nav first_name={first_name} setFirstName={setFirstName}/>

        <main className="form-signin">
            <Route path="/" exact component={() => <Home first_name={first_name}/>}/>
            <Route path="/login" component={() => <Login setFirstName={setFirstName}/>}/>
            <Route path="/register" component={Register}/>
            <Route path="/tutoring" exact component={TutoringOptions}/>
            <Route path="/tutoring/departments" exact component={ChooseDepartment}/>
            <Route path="/tutoring/departments/:course" component={ChooseTutor}/>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
