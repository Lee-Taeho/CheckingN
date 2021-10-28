import React, {useState, useEffect} from 'react';
import { Link } from 'react-router-dom';


const Nav = (props) => {
    const logout = async () => {
    //   await fetch('http://localhost:8080/api/logout', {
    //     method: 'POST',
    //     headers: {'Content-Type': 'application/json'},
    //     //credentials: 'include',
    // });
      
    // clear the local storage if user logout
      localStorage.clear();  
      setUser('');
      props.setFirstName('');
    } 

    const [user, setUser] =  useState(localStorage.getItem('profile'));
    console.log("user " + user);
    useEffect(() => {
      const token = user?.token;
      setUser(localStorage.getItem('profile'))
    })

    let menu;

    if(user ==='' || user == null) {
        menu = (
          <ul className="navbar-nav me-auto mb-2 mb-md-0">
              <li className="nav-item">
                <Link to="/login" className="nav-link active">Login</Link>
              </li>
              <li className="nav-item">
                <Link to="/register" className="nav-link active">Register</Link>
              </li>
          </ul>
        )
    } else {
      menu = (
        <ul className="navbar-nav me-auto mb-2 mb-md-0">
              <li className="nav-item">
                <Link to="/" className="nav-link active">{user}</Link>
              </li>
              <li className="nav-item">
                <Link to="/login" className="nav-link active" onClick={logout}>Logout</Link>
              </li>
            </ul>
      )
    }


    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
        <div className="container-fluid">
          <Link to="/" className="navbar-brand">Home</Link>
          
          <div>
            {menu}
          </div>

        </div>
      </nav>
    );
};

export default Nav;
