
//page1.tsx
import React  from "react";
import {Button} from 'antd';
import {Notify} from "../../../wailsjs/go/app/Application";
import './page1.css';

const Page1 = () => {
    const notify = async () => {
        await Notify()
    }
    return (
        <div>
            <h1>Page 1</h1>
            <Button onClick={notify}>Notify</Button>
        </div>
    )
}
export default Page1
