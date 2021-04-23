
import React, { useState, useEffect } from 'react';
import { Table } from 'antd';
import { uniFetch } from '../utils/apiUtils';

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

const BookList = () => {
  let [dataSource, setDataSource] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const result = await uniFetch("/book/all");
        setDataSource(result);
      } catch (e) {
        setDataSource([]);
        alert(e);
      }
    };
    fetchData();
  }, []);

  return (<Table dataSource={dataSource} columns={columns} />);
}

export default BookList;
