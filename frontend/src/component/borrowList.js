
import React, { useState, useEffect } from 'react';
import { uniFetch } from '../utils/apiUtils';
import BookTable from './bookTable';
import { Input } from 'antd';

const { Search } = Input;

const BorrowList = () => {
  let [dataSource, setDataSource] = useState([]);
  let [cardID, setCardID] = useState("12");

  useEffect(() => {
    const fetchData = async () => {
      if (cardID === undefined) {
        setDataSource([]);
      } else {
        try {
          const result = await uniFetch(`/borrow/book/all?card_id=${cardID}`);
          setDataSource(result);
        } catch (e) {
          setDataSource([]);
          alert(e);
        }
      }
    };
    fetchData();
  }, [cardID]);

  const onSearch = value => setCardID(value);

  return (<>
    <Search placeholder="借书证号" onSearch={onSearch} />
    <BookTable dataSource={dataSource} />
  </>);
};

export default BorrowList;
