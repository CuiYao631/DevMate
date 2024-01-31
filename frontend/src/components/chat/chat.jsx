// ChatPage.jsx

import React, { useState, useEffect } from 'react';
import { MessageList, Input, Button } from 'react-chat-elements';
import 'react-chat-elements/dist/main.css';
import './chat.css';
import {EventsOff, EventsOn} from '../../../wailsjs/runtime/runtime.js'


const ChatPage = () => {
    const [messages, setMessages] = useState([]);
    const [inputValue, setInputValue] = useState('');
    const [messageIdCounter, setMessageIdCounter] = useState(1);
    const [users, setUsers] = useState([]);
    const [currentChatUser, setCurrentChatUser] = useState(null);

    const handleSendMessage = () => {
        if (inputValue.trim() !== '' && currentChatUser) {
            const newMessage = {
                position: 'right',
                type: 'text',
                text: inputValue,
                date: new Date(),
                id: messageIdCounter,
                userId: currentChatUser.id,
            };

            setMessages((prevMessages) => [...prevMessages, newMessage]);
            setInputValue('');
            setMessageIdCounter(messageIdCounter + 1);
        }
    };

    // const handleSendMessage2 = () => {
    //     if (inputValue.trim() !== '' && currentChatUser) {
    //         const replyMessage = {
    //             position: 'left',
    //             type: 'text',
    //             text: inputValue,
    //             date: new Date(),
    //             id: messageIdCounter + 1,
    //             userId: currentChatUser.id,
    //         };
    //
    //         setMessages([...messages, replyMessage]);
    //         setMessageIdCounter(messageIdCounter + 2);
    //     }
    // };



    const handleUserClick = (user) => {
        setCurrentChatUser(user);
        setMessages([]);
    };
    const addUser = (data) => {
        // 获取数据数组
        const usersData = Object.values(data);
        // 清空现有用户列表
        setUsers([]);
        // 遍历所有用户数据
        usersData.forEach((userData) => {
            // 检查用户是否已经存在
            const userExists = users.some((user) => user.name === userData.Name);
            if (!userExists) {
                // 添加新用户
                setUsers((prevUsers) => {
                    const newUser = {
                        id: prevUsers.length + 1,
                        name: userData.Name,
                        avatar: 'https://bpic.588ku.com/element_origin_min_pic/01/48/73/4357443aa53522b.jpg',
                    };
                    return [...prevUsers, newUser];
                });
            }
        });
    };

    useEffect(() => {
        const handleDataUpdated = (data) => {
            // 监听来自 Go 后端的数据更新事件
            addUser(data);
        };

        // 监听来自 Go 后端的数据更新事件
        EventsOn('dataUpdated', addUser);

        // 在组件卸载时清理监听器
        return () => {
            EventsOff('dataUpdated', addUser);
        };
    }, []); // 空数组确保只在组件加载时执行一次

    // 按照消息的日期进行排序
    const sortedMessages = messages.slice().sort((a, b) => a.date - b.date);

    return (
        <div className="chat-container">
            <div className="user-list">
                {/* 用户列表 */}
                <h2>用户列表</h2>
                <ul>
                    {users.map((user) => (
                        <li key={user.id} onClick={() => handleUserClick(user)}>
                            <img src={user.avatar} alt={`${user.name} Avatar`} className="user-avatar" />
                            {user.name}
                        </li>
                    ))}
                </ul>
                {/*<Button*/}
                {/*    color='white'*/}
                {/*    backgroundColor='black'*/}
                {/*    text='添加用户'*/}
                {/*    onClick={addUser}*/}
                {/*/>*/}
            </div>
            <div className="chat-room">
                {/* 聊天页面内容 */}
                <div className="message-list">
                    {sortedMessages.map((message) => (
                        (currentChatUser && message.userId === currentChatUser.id) && (
                            <div key={message.id} className={`message-bubble ${message.position}`}>
                                {message.position === 'left' && (
                                    <img
                                        src={message.position === 'right' ? 'https://bpic.588ku.com/element_origin_min_pic/01/48/73/4357443aa53522b.jpg' : 'https://bpic.588ku.com/element_origin_min_pic/01/48/73/4357443aa53522b.jpg'}
                                        alt="Avatar"
                                        className="avatar"
                                    />
                                )}
                                <div className="message-content">{message.text}</div>
                                {message.position === 'right' && (
                                    <img
                                        src={message.position === 'right' ? 'https://bpic.588ku.com/element_origin_min_pic/01/48/73/4357443aa53522b.jpg' : 'https://bpic.588ku.com/element_origin_min_pic/01/48/73/4357443aa53522b.jpg'}
                                        alt="Avatar"
                                        className="avatar"
                                    />
                                )}
                            </div>
                        )
                    ))}
                </div>
                <Input
                    className="send-button"
                    placeholder="输入消息..."
                    multiline={true}
                    onChange={(e) => setInputValue(e.target.value)}
                    value={inputValue}
                    rightButtons={
                        <Button
                            color='white'
                            backgroundColor='black'
                            text='发送'
                            onClick={handleSendMessage}
                        />
                    }
                />
            </div>
        </div>
    );
};

export default ChatPage;
