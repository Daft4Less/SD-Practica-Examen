from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/')
def home():
    return jsonify({
        "mensaje": "¡Hola desde Python en Docker!",
        "practica": 3
    })