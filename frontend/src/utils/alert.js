import { message } from 'antd';

export const success = (string) => {
  if (typeof(string) !== 'string') {
    alert(string);
  }
  message.success(string);
};

export const error = (string) => {
  if (typeof(string) !== 'string') {
    alert(string);
  }
  message.error(string);
};
