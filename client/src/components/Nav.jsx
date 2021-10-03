import React from 'react';
import { Link } from 'react-router-dom';

const Nav = () => {
    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
        <div className="container-fluid">
          <Link to="/" className="navbar-brand">Home</Link>
          
          <div>
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
              <li className="nav-item">
                <Link to="/login" className="nav-link active">Login</Link>
              </li>
              <li className="nav-item">
                <Link to="/register" className="nav-link active">Register</Link>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    );
};

export default Nav;
