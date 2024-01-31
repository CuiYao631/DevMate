import React from 'react';
import { Layout, Menu, Breadcrumb, Form, Input, Button, Select, Switch } from 'antd';
import {
    UserOutlined,
    LaptopOutlined,
    GlobalOutlined,
    HighlightOutlined,
    SettingOutlined,
    BackwardFilled
} from '@ant-design/icons';

const { SubMenu } = Menu;
const { Content, Sider } = Layout;

const SettingsPage = () => {
    const onFinish = (values) => {
        console.log('Received values:', values);
    };

    return (
        <Layout style={{ minHeight: '100vh', background: 'transparent'}}>
            <Sider width={200} className="site-layout-background">
                <Menu
                    mode="inline"
                    defaultSelectedKeys={['1']}
                    defaultOpenKeys={['sub1']}
                    style={{ flex: 1, height: '100%', borderRight: 0 }}
                >
                    <Menu.Item key="1"
                               icon={<UserOutlined />}
                               style={
                                   {marginTop: '40px'}
                    }>
                        注册和登录
                    </Menu.Item>
                    <Menu.Item key="2" icon={<GlobalOutlined />}>
                        语言选择
                    </Menu.Item>
                    <Menu.Item key="3" icon={<HighlightOutlined />}>
                        主题选择
                    </Menu.Item>
                    <Menu.Item key="4" icon={<SettingOutlined />}>
                        快捷键设置
                    </Menu.Item>
                    <Menu.Item key="5" icon={<LaptopOutlined />}>
                        网络设置
                    </Menu.Item>
                    <Menu.Item key="6" icon={<BackwardFilled/>} onClick={() => window.location.href = '/'} style={{ marginBottom: 'auto' }}>
                        返回主页
                    </Menu.Item>
                </Menu>
            </Sider>
            <Layout style={{ padding: '0 24px 24px' }}>
                <Breadcrumb style={{ margin: '16px 0' }}>
                    <Breadcrumb.Item>首页</Breadcrumb.Item>
                    <Breadcrumb.Item>设置</Breadcrumb.Item>
                </Breadcrumb>
                <Content
                    className="site-layout-background"
                    style={{
                        padding: 24,
                        margin: 0,
                        minHeight: 280,
                        textAlign: 'left', // 文字左对齐
                        color: 'black',    // 文字颜色为黑色
                    }}
                >
                    <div>
                        {/* 注册和登录 */}

                        {/* 语言选择 */}
                        <div>
                            <h2>语言选择</h2>
                            <Select defaultValue="zh-CN" style={{ width: 120 }}>
                                <Select.Option value="zh-CN">中文</Select.Option>
                                <Select.Option value="en-US">English</Select.Option>
                            </Select>
                        </div>

                        {/* 主题选择 */}
                        <div>
                            <h2>主题选择</h2>
                            <Switch defaultChecked onChange={(checked) => console.log(`Switch to ${checked}`)} />
                        </div>

                        {/* 快捷键设置 */}
                        <div>
                            <h2>快捷键设置</h2>
                            {/* 这里可以添加快捷键设置的表单 */}
                        </div>

                        {/* 是否开启局域网发现功能 */}
                        <div>
                            <h2>是否开启局域网发现功能</h2>
                            <Switch defaultChecked />
                        </div>
                    </div>
                </Content>
            </Layout>
        </Layout>
    );
};

export default SettingsPage;
