import React from 'react';
import { Table } from 'antd';

const columns = [
  {
    title: '书号',
    dataIndex: 'book_id',
    key: 'book_id',
  },
  {
    title: '分类',
    dataIndex: 'category',
    key: 'category',
  },
  {
    title: '标题',
    dataIndex: 'title',
    key: 'title',
  },
  {
    title: '出版社',
    dataIndex: 'press',
    key: 'press',
  },
  {
    title: '出版时间',
    dataIndex: 'year',
    key: 'year',
  },
  {
    title: '作者',
    dataIndex: 'author',
    key: 'author',
  },
  {
    title: '单价',
    dataIndex: 'price',
    key: 'price',
  },
  {
    title: '总数',
    dataIndex: 'total',
    key: 'total',
  },
  {
    title: '在架数',
    dataIndex: 'stock',
    key: 'stock',
  },
];

class BookTable extends React.Component {
  render() {
    return (<Table dataSource={this.props.dataSource} columns={columns}  pagination={{ pageSize: 5 }}/>);
  }
};

export default BookTable;
