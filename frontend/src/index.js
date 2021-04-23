import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

import Content from './component/content';

import { PageHeader, Menu } from 'antd';
import { BookOutlined, TeamOutlined, UserSwitchOutlined } from '@ant-design/icons';

const { SubMenu } = Menu;

class App extends React.Component {
  state = {
    current: 'book:1'
  };

  handleClick = e => {
    console.log('click ', e);
    this.setState({ current: e.key });
  };

  render() {
    const { current } = this.state;
    return (
      <>
        <PageHeader
          className="site-page-header"
          title="图书管理系统"
          subTitle="eLibrary-DB2021"
        />
        <Menu onClick={this.handleClick} selectedKeys={[current]} mode="horizontal">
          <SubMenu key="Book" icon={<BookOutlined />} title="Book">
            <Menu.Item key="book:1">Search</Menu.Item>
            <Menu.Item key="book:2">Modify</Menu.Item>
          </SubMenu>
          <SubMenu key="Card" icon={<TeamOutlined />} title="Card">
            <Menu.Item key="card:1">Add</Menu.Item>
            <Menu.Item key="card:2">Modify</Menu.Item>
            <Menu.Item key="card:3">Delete</Menu.Item>
          </SubMenu>
          <SubMenu key="Borrow" icon={<UserSwitchOutlined />} title="Return/Borrow">
            <Menu.Item key="borrow:1">Return</Menu.Item>
            <Menu.Item key="borrow:2">Borrow</Menu.Item>
          </SubMenu>
        </Menu>
        <Content page={current} />
      </>
    );
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
