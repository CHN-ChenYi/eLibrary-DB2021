import React, { useState, useEffect } from 'react';
import { uniFetch } from '../utils/apiUtils';
import BookTable from './bookTable';
import { Input, Button } from 'antd';

const { Search } = Input;

const Return = () => {
  let [dataSource, setDataSource] = useState([]);
  let [cardID, setCardID] = useState();
  let [bookID, setBookID] = useState();
  let [searchCnt, setSearchCnt] = useState(0);

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
  }, [cardID, searchCnt]);

  const onSearch = value => setCardID(value);

  const onChange = value => { setBookID(value.target.value); };

  const onReturn = async () => {
    try {
      await uniFetch(`/borrow/book?card_id=${cardID}&book_id=${bookID}`, { method: 'Delete' });
      setSearchCnt(searchCnt + 1);
    } catch (e) {
      alert(e);
    }
  };

  return (<>
    <Search placeholder="借书证号" allowClear onSearch={onSearch} enterButton />
    <BookTable dataSource={dataSource} />
    <span>
      <Input type="text" placeholder="书号" onChange={onChange} />
      <Button type="primary" htmlType="submit" onClick={onReturn}>
        还书
      </Button>
    </span>
  </>);
};

export default Return;
