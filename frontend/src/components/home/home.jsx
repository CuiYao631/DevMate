// home.jsx

import React, {useEffect, useState} from 'react';
import './home.css';
import {useNavigate} from 'react-router-dom'

const Home = () => {
    function pageDimensions() {
        console.log('pageDimensions')
    }

    const [dimensions, setDimensions] = useState(pageDimensions)

    useEffect(() => {
        // the function you want to call
    }, [dimensions])

    return (
        <div className="main-container">
            <div className="menu-bar">
                {/* 菜单栏的内容放在这里 */}
                <ul>
                    <li>菜单项1</li>
                    <li>菜单项2</li>
                    <li>菜单项3</li>
                    {/* 添加更多菜单项 */}
                </ul>
            </div>
            <div className="content">
                {/* 主要内容区域放在这里 */}
                <h1>欢迎来到软件主页</h1>
                <p>这是主页的内容，你可以在这里添加更多信息。</p>
            </div>
        </div>
    );
};

export default Home;
