# Tarea3Dist

Conexión entre admin y broker, luego admin y DNS1 LOCAL

- Al ejecutar el admin pide un comando. Si este es "create", "update" o "delete" lo manda al broker. Este le responde con las IPs en las que estarán los DNS
- Luego se conecta al DNS con esta IP y le manda el mismo comando.

*Por ahora, solo se conecta al DNS1, y todo funciona de manera local. El Broker en el puerto 9000 y el DNS en el 9001.
