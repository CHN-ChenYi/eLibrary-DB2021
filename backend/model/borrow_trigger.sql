CREATE TRIGGER borrow_book BEFORE INSERT ON borrows
FOR EACH ROW
BEGIN
  IF 0 = (SELECT count(*) FROM books WHERE book_id = new.book_id) THEN
    SIGNAL SQLSTATE '45000' SET message_text = 'book_id invalid';
  END IF;
  IF 0 = (SELECT stock FROM books WHERE book_id = new.book_id) THEN
    SIGNAL SQLSTATE '45000' SET message_text = 'out of stock';
  END IF;
  UPDATE books SET stock = stock - 1 WHERE book_id = new.book_id;
END

CREATE TRIGGER return_book BEFORE UPDATE ON borrows
FOR EACH ROW
BEGIN
  IF old.return_date IS NULL and new.return_date IS NOT NULL
     AND old.id = new.id AND old.book_id = new.book_id
     AND old.card_id = new.card_id AND old.borrow_date = new.borrow_date THEN
    UPDATE books SET stock = stock + 1 WHERE book_id = new.book_id;
  END IF;
END
