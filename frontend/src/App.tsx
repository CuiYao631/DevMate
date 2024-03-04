
import './App.css';
import {BrowserRouter,Routes,Route} from "react-router-dom"
import Home from "./components/home/home";

function App() {


    return (
        <div id="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Home/>}/>
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App
