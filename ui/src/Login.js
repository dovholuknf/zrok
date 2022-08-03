import {useState} from "react";
import * as identity from './api/identity';

const Login = (props) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = async e => {
        e.preventDefault()

        identity.login({body: {"email": email, "password": password}})
            .then(resp => {
                if (!resp.error) {
                    let user = {
                        "email": email,
                        "token": resp.data
                    }
                    props.loginSuccess(user)
                    localStorage.setItem('user', JSON.stringify(user))
                    console.log(user)
                    console.log('login succeeded', resp)
                } else {
                    console.log('login failed')
                }
            })
            .catch((resp) => {
                console.log('login failed', resp)
            });
    };

    return (
        <div className={"login"}>
            <img src={"ziggy.svg"} width={200}/>
            <h1>zrok</h1>
            <form onSubmit={handleSubmit}>
                <fieldset>
                    <legend>Log In</legend>
                    <p><label htmlFor="email">email: </label><input type="text" value={email} placeholder="enter an email" onChange={({target}) => setEmail(target.value)}/></p>
                    <p><label htmlFor="password">password: </label><input type="password" value={password} placeholder="enter a password" onChange={({target}) => setPassword(target.value)}/><button type="submit">Log In</button></p>
                </fieldset>
            </form>
        </div>
    )
}

export default Login;