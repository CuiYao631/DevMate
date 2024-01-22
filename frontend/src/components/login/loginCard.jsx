import React, {useEffect, useState} from 'react';
import './LoginCard.css';
import {useNavigate} from 'react-router-dom'
import {EventsOn,EventsOff} from '../../../wailsjs/runtime/runtime.js'
import {PushDataToReact} from '../../../wailsjs/go/app/Application.js'

const LoginCard = ({showMessage}) => {
     const navigate=useNavigate();
    const[username,setUsername]=useState("");
    const[password,setPassword]=useState("");
    const [dataFromGo, setDataFromGo] = useState('');

    const handleUsernameChange=(e)=>{
        setUsername(e.target.value);
    }
    const handlePasswordChange=(e)=>{
        setPassword(e.target.value);
    }
    const handleSubmit=(e)=>{
        if (username===""){
            showMessage.errMessage("用户名不能为空")
            return;
        }
        if (password==="") {
            showMessage.errMessage("密码不能为空")
            return;
        }
        if (username==="admin"&&password==="admin") {
            navigate("/home");
        }else{
            showMessage.errMessage("用户名或密码错误")
        }
    }

    return (
        <div className="login-card">
            <div className="card-header">
                <div className="log">Login</div>
            </div>
            <form>
                <div className="form-group">
                    <label htmlFor="username">Username:</label>
                    <input required="" name="username" id="username" type="text" value={username} onChange={handleUsernameChange} />
                </div>
                <div className="form-group">
                    <label htmlFor="password">Password:</label>
                    <input required="" name="password" id="password" type="password" value={password} onChange={handlePasswordChange}/>
                </div>
                <div className="form-group">
                    <input value="Login" type="button" onClick={handleSubmit} />
                </div>
            </form>

        </div>
    );
}

export default LoginCard;
