import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import GoogleLogin from 'react-google-login';

const Login = (props) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e) => {
        e.preventDefault();

        var request = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            //credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        }

        // send email and password to login_request
        const response = await fetch('http://localhost:8080/api/login_request', request)
        // login_request returns a token info
        const data = await response.json()
        var key = data.key;
        var value = data.value;
        
        // use token info returned by login_request to set the header for api/authorized
        var authReq = {
             method: 'GET',
             headers: {'Authorization': String(value)}
        }

        const authResponse = await fetch('http://localhost:8080/api/authorized', authReq);

        const content = await authResponse.json();

        console.log(content);

        setRedirect(true);

        props.setFirstName(content.first_name);

    }

    if(redirect)
    {
        return <Redirect to="/"/>
    }
    
    // ***Google login
    // responseGoogle=(response) =>{
    //     console.log(response);
    //     console.log(response.profileObj);
    // }

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
            
            {/* <div>
                <GoogleLogin
                    // ask Ayush for clientId
                    clientId=""
                    buttonText="Sign in with Google"
                    onSuccess={this.responseGoogle}
                    onFailure={this.responseGoogle}
                    cookiePolicy={single_host_origin}
                />
            </div> */}
        </form>
        
    );
};

export default Login;
