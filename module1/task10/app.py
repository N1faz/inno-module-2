from flask import Flask, jsonify
import os
import datetime

app = Flask(__name__)

@app.route('/')
def hello():
    env = os.getenv('ENV', 'unknown')
    debug = os.getenv('DEBUG', 'False')
    return jsonify({
        'message': 'Hello from Docker Compose!',
        'environment': env,
        'debug': debug,
        'time': datetime.datetime.now().isoformat(),
        'host': os.uname().nodename
    })

@app.route('/health')
def health():
    return jsonify({'status': 'healthy'})

if __name__ == '__main__':
    app.run(
        host='0.0.0.0',
        port=5000,
        debug=os.getenv('DEBUG', 'False').lower() == 'true'
    )
