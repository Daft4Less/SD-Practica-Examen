const express = require('express');
const app = express();
app.get('/', (req, res) => res.send('API Node.js ejecutándose en Docker - Practica 2'));
app.listen(3000, () => console.log('Servidor en puerto 3000'));