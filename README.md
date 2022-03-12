## investment-distributor</br></br>
---

 <h4>⭐️ 1. Introducción</h4>

ㅤ➢ Cloud Hosting: Heroku</br>
ㅤ➢ Host: https://investment-distributor.herokuapp.com</br>
ㅤ➢ Port HTTP: 8000</br>
ㅤ➢ Port HTTPS: {****}</br>
ㅤ➢ DB: Firebase</br>
ㅤ➢ Objetivo: Otorgar la mejor opción de inversión en base al monto peticionado, junto con las estadísticas de las solicitudes realizadas. </br>

---

<h4>⭐️ 2. Lista de Operaciones</h4>
La API resuelve las siguientes consultas:</br></br>
• Operación: /credit-assignment/ </br>
• Funcionalidad: </br>
ㅤㅤ- Recibe un monto y lo distribuye entre tres opciones de créditos, de forma que no quede resto.</br>
• Consideraciones: </br>
ㅤㅤ- El monto ingresado debe ser múltiplo de 100. </br>
ㅤㅤ- Las opciones de créditos son de los montos: $300, $500 y $700. </br>
ㅤㅤ- Si el monto ingresado no puede retornar resto cero, entonces responderá con Status: 400 Bad Request. </br>
• Método: POST</br>
• URI: investment-distributor.herokuapp.com/credit-assignment</br></br>

<h4>Request</h4>

```json
{ "investment":300 }
```
<h4>Response</h4>

```json
{"credit_type_300":1,"credit_type_500":0,"credit_type_700":0}
```
---
</br>
• Operación: /statistics/ </br>
• Funcionalidad: </br>
ㅤㅤ- Retorna las siguientes estádisticas en base a la información almacenada en la base de datos:</br>
ㅤㅤTotal de asignaciones realizadas.</br>
ㅤㅤTotal de asignaciones exitosas.</br>
ㅤㅤTotal de asignaciones no exitosas.</br>
ㅤㅤPromedio de inversión exitosa.</br>
ㅤㅤPromedio de inversión no exitosa.</br>

• Consideraciones: </br>
ㅤㅤ- No es necesario el entregar un request para obtener los datos.</br>
• Método: POST</br>
• URI: investment-distributor.herokuapp.com/statistics</br></br>

<h4>Request</h4>

```json
{"average_successful_investment":300,"average_unsuccessful_investment":0,"total_assignments_made":2,"total_successful_assignments":2,"total_unsuccessful_assignments":0}
```
---
</br>
• Operación: /statistics/ </br>
• Funcionalidad: </br>
ㅤㅤ- Resetea las estadísticas eliminando los datos almacenados en la base de datos.</br>
• Método: DELETE</br>
• URI: investment-distributor.herokuapp.com/statistics</br></br>
