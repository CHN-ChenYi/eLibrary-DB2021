import React from 'react';
import { uniFetch } from '../utils/apiUtils';
import { success, error } from '../utils/alert';
import { Form, Input, Button } from 'antd';
import { layout, tailLayout } from './formLayout';

const Borrow = () => {
  const onFinish = async (values) => {
    try {
      await uniFetch(`/borrow/book?card_id=${values.cardID}&book_id=${values.bookID}`, { method: 'Post' });
      success("借书成功");
    } catch (e) {
      error(e);
    }
  };

  return (
    <Form
      {...layout}
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

      <Form.Item {...tailLayout}>
        <Button type="primary" htmlType="submit">
          借书
        </Button>
      </Form.Item>
    </Form>);
};

export default Borrow;
