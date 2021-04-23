import React, { useState, useEffect } from 'react';
import { Form, Row, Col, Input, InputNumber, Button } from 'antd';
import { uniFetch } from '../utils/apiUtils';
import { success, error } from '../utils/alert';
import BookTable from './bookTable';

const layout = {
  labelCol: {
    span: 12,
  },
  wrapperCol: {
    span: 12,
  },
};

const BookSearch = () => {
  const [form] = Form.useForm();
  const [dataSource, setDataSource] = useState([]);
  const [constrain, setConstrain] = useState({});
  const label = ["分类", "标题", "出版社", "作者", "出版时间", "价格"];
  const name = ['category', 'title', 'press', 'author', 'year', 'price'];
  const count = name.length;

  useEffect(() => {
    const fetchData = async () => {
      try {
        console.log(constrain);
        let query = "";
        const appendQuery = (index) => {
          if (constrain[index] !== undefined && constrain[index] !== null)
            query += `&${index}=${constrain[index]}`;
        };
        for (let i = 0; i < count; i++) {
          if (i > 3) {
            appendQuery(`${name[i]}_lowerbound`);
            appendQuery(`${name[i]}_upperbound`);
          } else {
            appendQuery(`${name[i]}`);
          }
        }
        let result;
        if (query.length === 0)
          result = await uniFetch("/book/all");
        else
          result = await uniFetch(`/book/search?` + query.substring(1));
        setDataSource(result);
        success("查询成功");
      } catch (e) {
        setDataSource([]);
        error(e);
      }
    };
    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [constrain]);

  const getFields = () => {
    const children = [];

    for (let i = 0; i < count; i++) {
      if (i > 3) {
        children.push(
          <Col key={i}>
            <Form.Item
              label={`${name[i]}`}>
              <Input.Group compact>
                <Form.Item name={`${name[i]}_lowerbound`} >
                  <Input
                    style={{ width: 100, textAlign: 'center' }}
                    placeholder="Minimum"
                  />
                </Form.Item>
                <Input
                  style={{
                    width: 30,
                    borderLeft: 0,
                    borderRight: 0,
                    pointerEvents: 'none',
                  }}
                  placeholder="~"
                  disabled
                />
                <Form.Item name={`${name[i]}_upperbound`} >
                  <Input
                    style={{ width: 100, textAlign: 'center' }}
                    placeholder="Maximum"
                  />
                </Form.Item>
              </Input.Group>
            </Form.Item>
            {/* <InputNumber style={{ width: '100%' }} /> */}
          </Col>,
        );
      } else {
        children.push(
          <Col key={i}>
            <Form.Item
              name={`${name[i]}`}
              label={`${label[i]}`}
            >
              <Input />
            </Form.Item>
          </Col>,
        );
      }
    }

    return children;
  };

  const onFinish = values => setConstrain(values);

  return (<>
    <Form
      {...layout}
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
    <div style={{ padding: "10px 0 0 0" }}>
      <BookTable dataSource={dataSource} />
    </div>
  </>);
};

export default BookSearch;
