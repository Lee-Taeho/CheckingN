import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);
    
    const submit = async (e) => {
        e.preventDefault();

        await fetch('http://localhost:8080/api/login_request', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
<<<<<<< HEAD
          //  credentials: 'include',
=======
            //credentials: 'include',
>>>>>>> fbcba355da181b51f740e32ad9bf21ca4873993b
            body: JSON.stringify({
                email,
                password
            })
        });

        // const content = await response.json();

        setRedirect(true);
        // props.setName(content.name);
    }

    if(redirect)
    {
        return <Redirect to="/"/>
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Sign in</h1>
            <label>Email</label> 
            <input type="email" name="Email" className="form-control" placeholder="Email" required
                onChange={e => setEmail(e.target.value)}
            />
            <label>Password</label>
            <input type="password" name="Password" className="form-control" placeholder="Password" required
                onChange={e => setPassword(e.target.value)}
            />
            <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
            <label>Don't have an account? <a href="/register">Register</a></label>
        </form>
        
    );
};

export default Login;
