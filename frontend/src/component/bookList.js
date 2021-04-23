import React, { useState, useEffect } from 'react';
import { uniFetch } from '../utils/apiUtils';
import BookTable from './bookTable';

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

  return (<BookTable dataSource={dataSource} />);
};

export default BookList;
