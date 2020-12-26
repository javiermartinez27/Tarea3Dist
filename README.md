# Tarea3Dist

Conexión entre admin y broker, luego admin y DNS1 LOCAL

- Al ejecutar el admin pide un comando. Si este es `create`, `update` o `delete` lo manda al broker. Este le responde con las IPs en las que estarán los DNS
- Luego se conecta al DNS con esta IP y le manda el mismo comando.
- Si se manda un `create`, crea un registro ZF (ejemplo: `create google.cl 123.0.0.1`), inicializa el reloj correspondiente y añade la acción al log. Si se manda otro create actualiza el reloj. Ojo con el formato del create.
- Si se manda un `delete` (ejemplo: `delete google.cl`), busca el nombre.dominio en el registro y lo borra. Además suma 1 al reloj y añade la acción al log.
*Por ahora, solo se conecta al DNS1, y todo funciona de manera local. El Broker en el puerto 9000 y el DNS en el 9001.

El cliente se conecta al Broker si hace una petición `get nombre.dominio` y este le responde con la dirección de algún DNS.
