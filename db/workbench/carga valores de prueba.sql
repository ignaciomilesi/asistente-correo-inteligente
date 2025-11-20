INSERT INTO usuarios (nombre) VALUES
('Mateo García'),
('Valentina López'),
('Lucas Fernández'),
('Isabella Martínez'),
('Thiago Rodríguez'),
('Sofía Torres');

INSERT INTO Pendientes 
(Titulo, Descripcion, Estado, Finalizada, Fecha_iniciada, Fecha_finalizada, Cierre, asignado)
VALUES
('Revisión de contratos', 'Analizar y actualizar los contratos del último trimestre.', 'En progreso', 0, '2025-10-20', NULL, NULL, 2),
('Actualización del sistema', 'Implementar la nueva versión del software interno.', 'Pendiente', 0, '2025-11-01', NULL, NULL, 4),
('Informe de ventas Q3', 'Preparar el informe de resultados de ventas del tercer trimestre.', 'Finalizado', 1, '2025-09-10', '2025-10-01', 'Enviado a gerencia', 1),
('Capacitación del equipo', 'Organizar capacitación sobre ciberseguridad para empleados.', 'En progreso', 0, '2025-10-15', NULL, NULL, 5),
('Revisión presupuestaria', 'Evaluar gastos y ajustar presupuesto mensual.', 'Pendiente', 0, '2025-11-05', NULL, NULL, 3),
('Diseño de campaña publicitaria', 'Crear materiales para la campaña de fin de año.', 'En progreso', 0, '2025-10-28', NULL, NULL, 6),
('Optimización del sitio web', 'Reducir tiempos de carga y mejorar SEO.', 'Finalizado', 1, '2025-08-12', '2025-09-05', 'Cierre aprobado', 4),
('Revisión de inventario', 'Auditar el inventario del almacén principal.', 'Pendiente', 0, '2025-11-02', NULL, NULL, 2),
('Preparación del evento anual', 'Coordinar logística y proveedores para el evento.', 'En progreso', 0, '2025-10-10', NULL, NULL, 1),
('Actualización de políticas internas', 'Revisar políticas de recursos humanos.', 'Pendiente', 0, '2025-11-09', NULL, NULL, 3);

INSERT INTO Avances 
(Fecha_Avance, Descripcion, ubicacion_mail, Pendientes_id)
VALUES
('2025-10-20', 'Pendiente registrada', '', 1),
('2025-10-22', 'Se revisaron contratos iniciales', '', 1),
('2025-10-25', 'Ajustes finales en contratos', '', 1),
('2025-11-01', 'Pendiente registrada', '', 2),
('2025-11-04', 'Versión beta instalada', '', 2),
('2025-09-10', 'Pendiente registrada', '', 3),
('2025-09-15', 'Informe preliminar enviado', '', 3),
('2025-09-25', 'Informe final revisado', '', 3),
('2025-10-15', 'Pendiente registrada', '', 4),
('2025-10-18', 'Primera sesión de capacitación', '', 4),
('2025-11-05', 'Pendiente registrada', '', 5),
('2025-11-07', 'Presupuesto preliminar evaluado', '', 5),
('2025-11-10', 'Ajustes realizados y aprobado', '', 5),
('2025-10-28', 'Pendiente registrada', '', 6),
('2025-10-30', 'Diseño de banners completado', '', 6),
('2025-08-12', 'Pendiente registrada', '', 7),
('2025-08-20', 'SEO básico implementado', '', 7),
('2025-11-02', 'Pendiente registrada', '', 8),
('2025-11-04', 'Inventario revisado parcialmente', '', 8),
('2025-10-10', 'Pendiente registrada', '', 9),
('2025-10-15', 'Proveedores contactados', '', 9),
('2025-11-09', 'Pendiente registrada', '', 10),
('2025-11-11', 'Políticas internas revisadas', '', 10);

INSERT INTO Etiquetas (nombre) VALUES
('Urgente'),
('Administrativo'),
('Cliente'),
('Interno'),
('Revisión'),
('Desarrollo'),
('Finanzas'),
('Legal'),
('Infraestructura'),
('Marketing');

INSERT INTO Pendientes_Etiquetas (Pendientes_id, Etiquetas_id) VALUES
(1, 8),
(1, 5),

(2, 6),
(2, 9),

(3, 7),
(3, 2),

(4, 4),

(5, 7),
(5, 5),

(6, 10),

(7, 6),
(7, 9),

(8, 2),
(8, 4),

(9, 3),
(9, 10),

(10, 8);

INSERT INTO Adjuntos (Descripcion, ubicacion_archivo, Pendientes_id) VALUES
('Contrato escaneado', '/uploads/contratos/contrato_q3.pdf', 1),
('Manual de instalación del sistema', '/uploads/sistema/manual_v2.docx', 2),
('Informe de ventas consolidado', '/uploads/ventas/informe_q3.xlsx', 3),
('Material de capacitación', '/uploads/capacitacion/ciberseguridad.pdf', 4),
('Presupuesto mensual preliminar', '/uploads/finanzas/presupuesto_nov2025.xlsx', 5);









