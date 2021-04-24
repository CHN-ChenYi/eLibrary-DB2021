import React from 'react';
import BookSearch from './bookSearch';
import BookModify from './bookModify';
import CardModify from './cardModify';
import Return from './return';
import Borrow from './borrow';

class Content extends React.Component {
  render() {
    if (this.props.page === "book:1")
      return (<BookSearch />);
    if (this.props.page === "book:2")
      return (<BookModify />);
    if (this.props.page === "Card")
      return (<CardModify />);
    if (this.props.page === "borrow:1")
      return (<Return />);
    if (this.props.page === "borrow:2")
      return (<Borrow />);
    return (
      <p>
        {this.props.page}
      </p>
    );
  }
}

export default Content;
