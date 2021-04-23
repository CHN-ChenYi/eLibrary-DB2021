import React, { useState, useEffect } from 'react';
import { Form, Row, Col, Input, Button } from 'antd';
import { uniFetch } from '../utils/apiUtils';
import BookTable from './bookTable';

const BookSearch = () => {
  const [form] = Form.useForm();
  const [dataSource, setDataSource] = useState([]);
  const [constrain, setConstrain] = useState({});
  const label = ["分类", "标题", "出版社", "作者", "出版时间（上界）", "出版时间（下界）", "价格（上界）", "价格（下界）"];
  const name = ['category', 'title', 'press', 'author', 'year_upperbound', 'year_lowerbound', 'price_upperbound', 'price_lowerbound'];
  const count = name.length;

  useEffect(() => {
    const fetchData = async () => {
      try {
        let query = "";
        for (let i = 0; i < count; i++) {
          if (constrain[name[i]] !== undefined)
            query += `&${name[i]}=${constrain[name[i]]}`;
        }
        let result;
        if (query.length === 0)
          result = await uniFetch("/book/all");
        else
          result = await uniFetch(`/book/search?` + query.substring(1));
        setDataSource(result);
      } catch (e) {
        setDataSource([]);
        alert(e);
      }
    };
    fetchData();
  }, [constrain]);

  const getFields = () => {
    const children = [];

    for (let i = 0; i < count; i++) {
      children.push(
        <Col span={6} key={i}>
          <Form.Item
            name={`${name[i]}`}
            label={`${label[i]}`}
          >
            <Input />
          </Form.Item>
        </Col>,
      );
    }

    return children;
  };

  const onFinish = values => setConstrain(values);

  return (<>
    <Form
      form={form}
      name="advanced_search"
      className="ant-advanced-search-form"
      onFinish={onFinish}
    >
      <Row gutter={24}>{getFields()}</Row>
      <Row>
        <Col
          span={24}
          style={{
            textAlign: 'right',
          }}
        >
          <Button type="primary" htmlType="submit">
            Search
          </Button>
          <Button
            style={{
              margin: '0 8px',
            }}
            onClick={() => {
              form.resetFields();
              setConstrain({});
            }}
          >
            Clear
          </Button>
        </Col>
      </Row>
    </Form>
    <BookTable dataSource={dataSource} />
  </>);
};

export default BookSearch;
