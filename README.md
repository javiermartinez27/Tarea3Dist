# Tarea3Dist

Conexión entre admin y broker, luego admin y DNSs LOCAL

- Al ejecutar el admin pide un comando. Si este es `create`, `update` o `delete` lo manda al broker. Este le responde con las IPs en las que estarán los DNS
- Luego se conecta al DNS con esta IP y le manda el mismo comando.
- Los 3 comandos (`create`, `delete` y `update`) están funcionando. 
- El `create` debe ser de la siguiente forma: `create nombre.dominio IP`. Ejemplo: `create google.cl 123.0.0.1`. Los demás comandos son tal cual en el PDF de la tarea.

# DNS

- Crean todo lo que tengan que crear (registros ZF, logs y relojes). PERO la consistencia del admin no está hecha, por lo que no sé cómo reaccionan si se les manda un `update` o un `delete`.

# Cliente

- El cliente se conecta al Broker si hace una petición `get nombre.dominio` y este le responde con la dirección de algún DNS.
