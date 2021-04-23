import React, { useState, useEffect } from 'react';
import { uniFetch } from '../utils/apiUtils';
import BookTable from './bookTable';
import { Form, Input, Button } from 'antd';

const { Search } = Input;

const Return = () => {
  const [dataSource, setDataSource] = useState([]);
  const [cardID, setCardID] = useState();
  const [searchCnt, setSearchCnt] = useState(0);

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

  const onReturn = async (value) => {
    try {
      await uniFetch(`/borrow/book?card_id=${cardID}&book_id=${value.bookID}`, { method: 'Delete' });
      setSearchCnt(searchCnt + 1);
    } catch (e) {
      alert(e);
    }
  };

  return (<>
    <Form><Form.Item label="借书证号" >
      <Search allowClear onSearch={onSearch} enterButton />
    </Form.Item></Form>
    <BookTable dataSource={dataSource} />
    <div style={{ padding: "20px 0 0 0" }}>
      <Form onFinish={onReturn}>
        <Form.Item label="书号" name="bookID" rules={[{ required: true, message: '请输入书号!' }]}>
          <Input type="text" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            还书
        </Button>
        </Form.Item>
      </Form>
    </div>
  </>);
};

export default Return;
