import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import GoogleLogin from 'react-google-login';

const Login = (props) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    // Google api
	const GOOGLE_CLIENT_ID = "533962375262-sn2l2op591vabl5i85f6vf7sptad47tt.apps.googleusercontent.com";
    const GOOGLE_EMAIL_SCOPE    = "https://www.googleapis.com/auth/userinfo.email";
    const GOOGLE_PROFILE_SCOPE  = "https://www.googleapis.com/auth/userinfo.profile";
    const GOOGLE_CALENDAR_SCOPE = "https://www.googleapis.com/auth/calendar";

    // Successfully login with Google
    const onGoogleLoginSuccess = async (response) => {
        console.log(response);
        console.log('Google Login Success:', response.profileObj);
        
        const result = response?.profileObj;
        
        // redirect to home after user login with google
        setRedirect(true);

        // store user email and profileObj in local storage
        localStorage.setItem('profile', response.profileObj.email);
        localStorage.setItem('header', response.profileObj);    
        props.setFirstName(response.profileObj.email);

        // fetch user info to server side to store into database
        var googleLoginRequest = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                id: response.profileObj.googleId,
                email: response.profileObj.email,
                first_name: response.profileObj.givenName,
                last_name: response.profileObj.familyName,
                picture_link: response.profileObj.imageUrl
            })
        }

        const googleResponse = await fetch('http://localhost:8080/api/google_login_request', googleLoginRequest);
    }

    // Fail to login with Google
    const onGoogleLoginFailure = (response) => {
        console.log(response);
        console.log('Google Login Failed:', response.profileObj);
    }

    // Regular login
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

        // redirect to home after user login
        setRedirect(true);

        // store user info through refreshes
        localStorage.setItem('profile', content.email);
        props.setFirstName(content.email);
    }

    // const handleGoogleLogin = async (e) => {

    //     var googleLoginRequest = {
    //         method: 'GET',
    //         headers: {
    //          //   'X-Requested-With': 'XMLHttpRequest',
    //             'Content-Type': 'application/json'
    //         }
    //     }

    //     const googleResponse = await fetch('http://localhost:8080/api/google_callback', googleLoginRequest)
    //         // .catch(err => {
    //         //     throw new Error(err)
    //         // });

    //     const googleData = await googleResponse.json()
    //     console.log(googleData);
    //     var gKey = googleData.key;
    //     var gValue = googleData.value;
        
    //     // use token info returned by login_request to set the header for api/authorized
    //     var gAuthReq = {
    //             method: 'GET',
    //             headers: {'Authorization': String(gValue)}
    //     }

    //     const gAuthResponse = fetch('http://localhost:8080/api/authorized', gAuthReq);

    //     const gContent = gAuthResponse.json();

    //     console.log(gContent);
        
    //     setRedirect(true);

    //     props.setFirstName(gContent.first_name);
    // }


    // redirect to home
    if(redirect)
    {
        return <Redirect to="/"/>
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal" style={{textAlign: "center"}}>Sign in</h1>
            <br/>
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
            
            <hr/>
            <div className="or-container">
                <label>or</label>
            </div>
        
            <div>
                <GoogleLogin
                    clientId={GOOGLE_CLIENT_ID}
                    buttonText="Log in with Google"
                    onSuccess={onGoogleLoginSuccess}
                    onFailure={onGoogleLoginFailure}
                    cookiePolicy={'single_host_origin'}
                    //isSignedIn={true}
                    scope={GOOGLE_EMAIL_SCOPE, GOOGLE_PROFILE_SCOPE, GOOGLE_CALENDAR_SCOPE}
                />
            </div>
        </form>
        
    );
};

export default Login;
