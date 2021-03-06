import React, { useState, useEffect } from 'react';
import { Form, Input, Button, Radio } from 'antd';
import { uniFetch } from '../utils/apiUtils';
import { success, error } from '../utils/alert';
import { layout, tailLayout } from './formLayout';

const CardModify = () => {
  const [form] = Form.useForm();
  const [card, setCard] = useState({});

  useEffect(() => {
    form.setFieldsValue(card);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [card]);

  const onFinish = async (values) => {
    try {
      if (values.operation === 'Get') {
        const result = await uniFetch(`/card?card_id=${values.card_id}`);
        setCard(result);
        success('查询成功');
      } else if (values.operation === 'Delete') {
        await uniFetch(`/card?card_id=${values.card_id}`, { method: 'Delete' });
        success('删除成功');
      } else {
        await uniFetch(`/card`, { method: values.operation, body: values });
        success(values.operation === 'Post' ? '新建成功' : '修改成功');
      }
    } catch (e) {
      error(e);
    }
  };

  const label = ["借书证号", "部门", "类型"];
  const name = ['card_id', 'department', 'type'];
  const count = name.length;

  const getFields = () => {
    const children = [];

    for (let i = 0; i < count; i++) {
      if (i === 0) {
        children.push(
          <Form.Item
            name={`${name[i]}`}
            label={`${label[i]}`}
            rules={[{ required: true, message: '请输入借书证号!' }]}
          >
            <Input />
          </Form.Item>
        );
      } else if (i > 1) {
        children.push(
          <Form.Item
            name={`${name[i]}`}
            label={`${label[i]}`}
          >
            <Radio.Group>
              <Radio.Button value="S">学生</Radio.Button>
              <Radio.Button value="T">教师</Radio.Button>
            </Radio.Group>
          </Form.Item>
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
          <Radio.Button value="Delete">删除</Radio.Button>
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

export default CardModify;
