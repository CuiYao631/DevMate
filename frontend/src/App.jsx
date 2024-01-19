import {useState} from 'react';
import './App.css';
import {BrowserRouter,Routes,Route,Navigate} from 'react-router-dom'
import {CreateKsuid} from "../wailsjs/go/app/Application";
import LoginCard from './components/login/loginCard.jsx'
import Home from './components/home/home.jsx'
import { Toaster, toast } from 'sonner'


function App() {
    const [resultText, setResultText] = useState("Ksuid");
    const updateResultText = (result) => setResultText(result);

    function greet() {
        CreateKsuid().then(result=>{
            updateResultText(result);
        })
    }
    const showMessage = {
        message: (e) => {
            toast(e);
        },
        errMessage: (e) => {
            toast.error(e,{
                style: {
                    backgroundColor: '#fff0f0',
                    borderColor: '#ffe0e1',
                    color: '#e60000',
                },
            })
        }
    };
    return (
        <div>
            <Toaster  position="top-center" expand={true} />
            <BrowserRouter>
                <Routes>
                    <Route path="/login" element={<LoginCard showMessage={showMessage}/>}/>
                    <Route path="/home" element={<Home/>}/>
                    <Route  path="/" element={<Navigate to="/login"/>}/>
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App
