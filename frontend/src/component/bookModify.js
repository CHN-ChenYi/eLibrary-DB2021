import React, { useState, useEffect } from 'react';
import { Form, Input, InputNumber, Button, Radio } from 'antd';
import { uniFetch } from '../utils/apiUtils';
import { layout, tailLayout } from './formLayout';

const BookModify = () => {
  const [form] = Form.useForm();
  const [book, setBook] = useState({});

  useEffect(() => {
    form.setFieldsValue(book);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [book]);

  const onFinish = async (values) => {
    try {
      if (values.operation === 'Get') {
        const result = await uniFetch(`/book?book_id=${values.book_id}`);
        setBook(result);
        alert('查询成功');
      } else {
        await uniFetch(`/book`, { method: values.operation, body: values });
        alert(values.operation === 'Post' ? '新建成功' : '修改成功');
      }
    } catch (e) {
      alert(e);
    }
  };

  const label = ["书号", "分类", "标题", "出版社", "作者", "出版时间", "价格", "总数", "在架数"];
  const name = ['book_id', 'category', 'title', 'press', 'author', 'year', 'price', 'total', 'stock'];
  const count = name.length;

  const getFields = () => {
    const children = [];

    for (let i = 0; i < count; i++) {
      if (i === 0) {
        children.push(
          <Form.Item
            name={`${name[i]}`}
            label={`${label[i]}`}
            rules={[{ required: true, message: '请输入书号!' }]}
          >
            <Input />
          </Form.Item>
        );
      } else if (i > 4) {
        children.push(
          <Form.Item
            name={`${name[i]}`}
            label={`${label[i]}`}
          >
            <InputNumber style={{ width: '100%' }} />
          </Form.Item >
        );
      } else {
        children.push(
          <Form.Item
            name={`${name[i]}`}
            label={`${label[i]}`}
          >
            <Input />
          </Form.Item>
        );
      }
    }

    return children;
  };

  return (
    <Form {...layout} form={form} onFinish={onFinish}>
      {getFields()}
      <Form.Item label="操作" name="operation">
        <Radio.Group>
          <Radio.Button value="Get">查询</Radio.Button>
          <Radio.Button value="Post">新建</Radio.Button>
          <Radio.Button value="Put">修改</Radio.Button>
        </Radio.Group>
      </Form.Item>
      <Form.Item {...tailLayout}>
        <Button type="primary" htmlType="submit">
          提交
        </Button>
      </Form.Item>
    </Form>
  );
};

export default BookModify;
