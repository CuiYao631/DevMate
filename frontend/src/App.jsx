import {useState} from 'react';
import './App.css';
import {CreateKsuid} from "../wailsjs/go/app/Application";

function App() {
    const [resultText, setResultText] = useState("Ksuid");
    const updateResultText = (result) => setResultText(result);

    function greet() {
        CreateKsuid().then(result=>{
            updateResultText(result);
        })
    }

    return (
        <div id="App">
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <button className="btn" onClick={greet}>生成Ksuid</button>
            </div>
        </div>
    )
}

export default App
