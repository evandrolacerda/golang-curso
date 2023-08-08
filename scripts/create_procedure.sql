CREATE PROCEDURE listar_livros AS
BEGIN
	SELECT * FROM livros 
END


CREATE PROCEDURE listar_livros @id int AS
BEGIN
	SELECT * FROM livros WHERE id = @id
END
