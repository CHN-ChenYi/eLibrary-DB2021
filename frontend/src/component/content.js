import React from 'react';
import BookList from './bookList';
import BorrowList from './borrowList';

class Content extends React.Component {
  render() {
    if (this.props.page === "book:1")
      return (<BookList />);
    if (this.props.page === "borrow:1")
      return (<BorrowList />);
    return (
      <p>
        {this.props.page}
      </p>
    );
  }
}

export default Content;
