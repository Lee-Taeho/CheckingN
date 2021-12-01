import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';

const Register = () => {
    const [first_name, setFirstName] = useState('');
    const [last_name, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e) => {
        e.preventDefault();

        await fetch('http://localhost:8080/api/save_new_user', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                first_name,
                last_name,
                email,
                password
            })
        });

        setRedirect(true);
    }

    if(redirect) {
        return <Redirect to="/login"/>;
    }

    return (
        <div className="form-signin">
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal" style={{textAlign: "center"}}>Create Your Account</h1>
            <label>First Name</label>
            <input type="text" name="first_name" className="form-control" placeholder="First Name" required
                onChange={e => setFirstName(e.target.value)}
            />
            <label>Last Name</label>
            <input type="text" name="last_name" className="form-control" placeholder="Last Name" required
                onChange={e => setLastName(e.target.value)}
            />
            <label>Email</label>
            <input type="email" name="email" className="form-control" placeholder="Email" required
                onChange={e => setEmail(e.target.value)}
            />
            <label>Password</label>
            <input type="password" name="password" className="form-control" placeholder="Password" required
                onChange={e => setPassword(e.target.value)}    
            />
            <button className="w-100 btn btn-lg btn-primary" type="submit">Create</button>

            <label>Already have an account? <a href="/login">Login</a></label>
        </form>
        </div>
    );
};

export default Register;
