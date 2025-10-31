# Asistente de correo inteligente

La aplicación procesa correos electrónicos (.msg de Outlook) para ayudar a organizar y resumir información de manera automática.  

El programa recibe un email como entrada y realiza tres tareas principales:

- __Resumen automático:__
  
Analiza el contenido del correo y genera un resumen breve que captura los puntos clave del mensaje.

- __Identificación del tema:__
  
Compara el contenido del correo con una base de temas previamente definidos (por ejemplo, proyectos, clientes o categorías de trabajo) y determina a cuál corresponde.  
Si el tema ya existe, lo asocia; si no, puede sugerir uno nuevo.

- __Actualización de la cronología:__

Agrega el resumen del correo a una línea de tiempo (timeline) dentro del tema correspondiente, permitiendo seguir la evolución de conversaciones o proyectos a lo largo del tiempo.

__Objetivo:__ Facilitar la gestión y el seguimiento de comunicaciones por correo electrónico, reduciendo la carga manual de clasificación y registro.


### Tecnologías 

- Lenguaje: Go (Golang)
- Ollama: Ejecución local de LLMs
