import React, { useState, useEffect } from 'react';
import { Space, Form, Input, InputNumber, Button } from 'antd';
import { uniFetch } from '../utils/apiUtils';
import { success, error } from '../utils/alert';
import BookTable from './bookTable';
import { layout } from './formLayout';

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
        children.push(<>
          <Form.Item {...layout}
            label={`${label[i]}`}
            style={{ margin: "0 0 0 0" }}
            key={i}
          >
            <Input.Group compact>
              <div style={{ width: '45%' }}><Form.Item name={`${name[i]}_lowerbound`} >
                <InputNumber
                  style={{ width: '100%', textAlign: 'center' }}
                  placeholder="Minimum"
                />
              </Form.Item></div>
              <Input
                style={{
                  width: '10%',
                  borderLeft: 0,
                  borderRight: 0,
                  pointerEvents: 'none',
                  textAlign: 'center',
                }}
                placeholder="~"
                disabled
              />
              <div style={{ width: '45%' }}><Form.Item name={`${name[i]}_upperbound`} >
                <InputNumber
                  style={{ width: '100%', textAlign: 'center' }}
                  placeholder="Maximum"
                />
              </Form.Item></div>
            </Input.Group>
          </Form.Item>
        </>
          ,
        );
      } else {
        children.push(
          <Form.Item {...layout}
            name={`${name[i]}`}
            label={`${label[i]}`}
            key={i}
          >
            <Input />
          </Form.Item>,
        );
      }
    }

    return children;
  };

  const onFinish = values => setConstrain(values);

  return (<>
    <Space direction="vertical">
      <Form
        form={form}
        name="advanced_search"
        className="ant-advanced-search-form"
        onFinish={onFinish}
      >
        {getFields()}
      </Form>
      <div style={{ width: "50%", marginLeft: "250px" }}>
        <Space direction="horizontal">
          <Button type="primary" htmlType="submit" style={{ width: '80px' }}>
            Search
          </Button>
          <Button
            style={{
              margin: '0 8px',
              width: '80px'
            }}
            onClick={() => {
              form.resetFields();
              setConstrain({});
            }}
          >
            Clear
          </Button>
        </Space></div>
      <div style={{ paddingTop: "10px", width: "50%", marginLeft: "35px" }}>
        <BookTable dataSource={dataSource} />
      </div>
    </Space>
  </>
  );
};

export default BookSearch;
