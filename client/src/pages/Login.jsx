import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import GoogleLogin from 'react-google-login';



const Login = (props) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);
    const [isGoogleLogin, setIsGoogleLogin] = useState(false);

	//const gOOGLE_CLIENT_ID = "533962375262-dvbofsom1ocmc8mjb6aq4alfd7rdisi0.apps.googleusercontent.com"

    // const onGoogleLoginSuccess = (response) => {
    //     console.log(response);
    //     console.log('Google Login Success:', response.profileObj);
    // }

    // const onGoogleLoginFailure = (response) => {
    //     console.log(response);
    //     console.log('Google Login Failed:', response.profileObj);
    // }

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
    
    // const handleGoogleLogin = () => {
    //    // return <Redirect to="/register"/>
    //     //console.log("click")
    //      return <Redirect to="/api/google_login_request"/>
    //     // var googleLoginRequest = {
    //     //     method: 'GET',
    //     //     headers: {'Content-Type': 'application/json'},
    //     // }

    //     // const googleResponse = await fetch('http://localhost:8080/api/google_login_request', googleLoginRequest)
    // }    

    

    if(isGoogleLogin) {
        var googleLoginRequest = {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
       }

       const authResponse = fetch('http://localhost:8080/api/google_login_request', googleLoginRequest);
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal" style={{textAlign: "center"}}>Sign in</h1>
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
        
            <button type="button" class="google-login-btn" onClick={() => setIsGoogleLogin(true)}>
                <svg width="40" height="40">
                    <g id="Google-Button" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
                        <rect x="0" y="0" width="40" height="40" rx="2" style={{fill: "rgb(255, 255, 255)"}}></rect>
                        <g id="logo_googleg_48dp" transform="translate(10.5, 10.5) scale(1.1)">
                            <path d="M17.64,9.20454545 C17.64,8.56636364 17.5827273,7.95272727  17.4763636,7.36363636 L9,7.36363636 L9,10.845 L13.8436364,10.845 C13.635,11.97 13.0009091,12.9231818 12.0477273,13.5613636 L12.0477273,15.8195455 L14.9563636,15.8195455 C16.6581818,14.2527273 17.64,11.9454545 17.64,9.20454545 L17.64,9.20454545 Z" id="Shape" fill="#4285F4"></path>
                            <path d="M9,18 C11.43,18 13.4672727,17.1940909 14.9563636,15.8195455 L12.0477273,13.5613636 C11.2418182,14.1013636 10.2109091,14.4204545 9,14.4204545 C6.65590909,14.4204545 4.67181818,12.8372727  3.96409091,10.71 L0.957272727,10.71 L0.957272727,13.0418182  C2.43818182,15.9831818 5.48181818,18 9,18 L9,18 Z" id="Shape" fill="#34A853"></path>
                            <path d="M3.96409091,10.71 C3.78409091,10.17 3.68181818,9.59318182 3.68181818,9  C3.68181818,8.40681818 3.78409091,7.83 3.96409091,7.29 L3.96409091,4.95818182  L0.957272727,4.95818182 C0.347727273,6.17318182 0,7.54772727 0,9 C0,10.4522727  0.347727273,11.8268182 0.957272727,13.0418182 L3.96409091,10.71 L3.96409091,10.71 Z" id="Shape" fill="#FBBC05"></path>
                            <path d="M9,3.57954545 C10.3213636,3.57954545 11.5077273,4.03363636 12.4404545,4.92545455  L15.0218182,2.34409091 C13.4631818,0.891818182 11.4259091,0 9,0 C5.48181818,0  2.43818182,2.01681818 0.957272727,4.95818182 L3.96409091,7.29 C4.67181818,5.16272727  6.65590909,3.57954545 9,3.57954545 L9,3.57954545 Z" id="Shape" fill="#EA4335"></path>
                            <path d="M0,0 L18,0 L18,18 L0,18 L0,0 Z" id="Shape"></path>
                        </g>
                    </g>
                </svg>
                <div class="google-btn-text-container">
                    <div class="google-btn-text">Log in with Google</div>
                </div>
            </button>

        </form>
        
    );
};

export default Login;
