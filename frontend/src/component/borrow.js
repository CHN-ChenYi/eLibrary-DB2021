import React from 'react';
import { uniFetch } from '../utils/apiUtils';
import { Form, Input, Button } from 'antd';

const Borrow = () => {
  const onFinish = async (values) => {
    try {
      await uniFetch(`/borrow/book?card_id=${values.cardID}&book_id=${values.bookID}`, { method: 'Post' });
      alert("借书成功");
    } catch (e) {
      alert(e);
    }
  };

  return (
    <Form
      name="basic"
      onFinish={onFinish}
    >
      <Form.Item
        label="借书证号"
        name="cardID"
        rules={[{ required: true, message: '请输入借书证号!' }]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="书号"
        name="bookID"
        rules={[{ required: true, message: '请输入书号!' }]}
      >
        <Input />
      </Form.Item>

      <Form.Item>
        <Button type="primary" htmlType="submit">
          借书
    </Button>
      </Form.Item>
    </Form>);
};

export default Borrow;
