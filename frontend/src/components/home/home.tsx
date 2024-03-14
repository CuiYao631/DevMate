import React, {useEffect, useState} from 'react';
import './home.css';
import {
    DesktopOutlined,
    FileOutlined,
    PieChartOutlined,
    TeamOutlined,
    UserOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    SearchOutlined,
    NotificationOutlined,
} from '@ant-design/icons';
import {Avatar, Badge, Button, ConfigProvider, Flex, FlexProps, Input, MenuProps, Radio, Switch} from 'antd';
import { Breadcrumb, Layout, Menu, theme } from 'antd';
import Page1 from "../page1/page1";


const { Header, Content, Footer, Sider } = Layout;

type MenuItem = Required<MenuProps>['items'][number];

function getItem(
    label: React.ReactNode,
    key: React.Key,
    icon?: React.ReactNode,
    children?: MenuItem[],
    ): MenuItem {
        return {
            key,
            icon,
            children,
            label,
        } as MenuItem;
    }



const Home = () => {
    const [collapsed, setCollapsed] = useState(true);
    const [show, setShow] = useState(false);
    const [selectedMenu, setSelectedMenu] = useState('1'); // 默认选中第一个菜单项
    const [title,setTitle] = useState('');
    const onChange = (checked: boolean) => {
        setShow(checked);
    };
    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();

    const items: MenuItem[] = [
        getItem('test 1', '1', <PieChartOutlined />),
        getItem('test 2', '2', <DesktopOutlined />),
        getItem('test 3', '3', <UserOutlined />, ),
        getItem('test 4', '4', <TeamOutlined />, ),
        getItem('test 5', '5', <FileOutlined />),
    ];


    // const notify = async () => {
    //     await Notify()
    //     setShow(true);
    // }

    // 渲染不同的内容
    const renderContent = () => {
        switch (selectedMenu) {
            case '1':
                return <Page1/>;
            case '2':
                return <div>Desktop Content</div>;
            case '3':
                return <div>User Content</div>;
            case '4':
                return <div>Team Content</div>;
            case '5':
                return <div>File Content</div>;
            default:
                return null;
        }
    };

    return (
        <Layout  style={{ minHeight: '100vh' }}>
            <Sider trigger={null} collapsible collapsed={collapsed}>
                <Flex gap="middle" vertical>
                    <Flex >
                        <img  className="logo"  src="../../../src/assets/images/log.png" alt="logo"/>
                        <text className="title">{title}</text>
                    </Flex>
                </Flex>
                <Menu
                    theme="dark"
                    mode="inline"
                    defaultSelectedKeys={['1']}
                    items={items}
                    onSelect={({key}) => setSelectedMenu(key.toString())}
                />
            </Sider>
            <Layout>
                <Header style={{padding: 0, background: colorBgContainer}}>
                    <Flex justify="end" align="middle" style={{height: '100%'}}>
                    <Button
                            type="text"
                            icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                            onClick={() => {setCollapsed(!collapsed)

                            if (collapsed) {
                                setTitle('DevMate')
                            } else {
                                setTitle('')
                            }
                            }}
                            style={{
                                fontSize: '16px',
                                width: 64,
                                height: 64,
                            }}
                        />
                        <Input style={{
                            fontSize: '16px',
                            height: 40,marginTop: 12,marginLeft: 12,marginRight: 12,
                        }} size="large" placeholder="" prefix={<SearchOutlined />} />

                        <div style={{ marginLeft: 'auto', display: 'flex', alignItems: 'center', fontSize: '16px',
                            width: 64,
                            height: 64, }}>
                            <Badge dot={show}>
                                <Avatar style={{
                                    width: 40,
                                    height: 40,
                                }}
                                        shape="square"
                                        icon={<UserOutlined />}
                                        src={"https://avatars.githubusercontent.com/u/48278927?s=400&u=53dfef89d0f441243d82e2dc4f85d3897a67f3a6&v=4"}
                                />
                            </Badge>
                        </div>
                    </Flex>
                </Header>
                <Content style={{ margin: '24px 16px', padding: 24, minHeight: 280, background: colorBgContainer, borderRadius: borderRadiusLG }}>
                    {renderContent()} {/* 渲染不同的内容 */}
                </Content>
            </Layout>
        </Layout>
    );
};

export default Home;
