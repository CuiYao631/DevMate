import React, {useEffect, useState} from 'react';
import './home.css';
import {
    DesktopOutlined,
    FileOutlined,
    PieChartOutlined,
    TeamOutlined,
    UserOutlined,
    SearchOutlined,
    NotificationOutlined,
} from '@ant-design/icons';
import {Avatar, Badge, Button, ConfigProvider, MenuProps, Switch} from 'antd';
import { Breadcrumb, Layout, Menu, theme } from 'antd';
import {Notify} from "../../../wailsjs/go/app/Application";


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

const items: MenuItem[] = [
    getItem('test 1', '1', <PieChartOutlined />),
    getItem('test 2', '2', <DesktopOutlined />),
    getItem('test 3', '3', <UserOutlined />, ),
    getItem('test 4', '4', <TeamOutlined />, ),
    getItem('test 5', '5', <FileOutlined />),
];

const Home = () => {
    const [collapsed, setCollapsed] = useState(true);
    const [show, setShow] = useState(true);
    const onChange = (checked: boolean) => {
        setShow(checked);
    };
    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();

    const notify = async () => {
        await Notify()

    }

    return (
                <Layout style={{ minHeight: '100vh' }}>
                <Sider theme="dark" collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
                    <div className="demo-logo-vertical" />

                    <Menu  theme="dark" defaultSelectedKeys={['1']} mode="inline" items={items} />
                </Sider>
            <Layout>
                <Header style={{ padding: 0, background: colorBgContainer }} >
                    <Button icon={<NotificationOutlined />}  onClick={notify}/>
                    <Badge dot={show}>
                        <Avatar shape="square" icon={<UserOutlined />} />
                    </Badge>
                    <Switch onChange={onChange} checked={show} />
                </Header>
                <Content style={{ margin: '0 16px' }}>
                    <Breadcrumb style={{ margin: '16px 0' }}>
                        <Breadcrumb.Item>User</Breadcrumb.Item>
                        <Breadcrumb.Item>Bill</Breadcrumb.Item>
                    </Breadcrumb>
                    <div
                        style={{
                            padding: 24,
                            minHeight: 360,
                            background: colorBgContainer,
                            borderRadius: borderRadiusLG,
                        }}
                    >
                        Bill is a cat.

                    </div>

                </Content>
                <Footer style={{ textAlign: 'center' }}>
                    Ant Design ©{new Date().getFullYear()} Created by Ant UED
                </Footer>

            </Layout>
        </Layout>
    );
};

export default Home;
