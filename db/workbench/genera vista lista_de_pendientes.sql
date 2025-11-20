CREATE VIEW `Lista_de_pendientes` AS
    SELECT 
        p.id, p.Titulo, p.Descripcion, p.Estado, a.Fecha_Avance
    FROM
        Pendientes AS p
            LEFT JOIN
        Avances AS a ON a.id = (SELECT 
                a2.id
            FROM
                Avances AS a2
            WHERE
                a2.Pendientes_id = p.id
            ORDER BY a2.Fecha_Avance DESC
            LIMIT 1)
    WHERE
        p.Finalizada = 0;
        
select * from `lista de pedientes`;