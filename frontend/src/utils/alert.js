import { message } from 'antd';

export const success = (string) => {
  message.success(string);
};

export const error = (string) => {
  message.error(string);
};
