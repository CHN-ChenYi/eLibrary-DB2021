import React from 'react';
import BookList from './bookList';

class Content extends React.Component {
  render() {
    if (this.props.page === "book:1")
      return (<BookList />);
    return (
      <p>
        {this.props.page}
      </p>
    );
  }
}

export default Content;
