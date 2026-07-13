from flask import Flask, render_template
import datetime
import os

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html', 
                         time=datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S'),
                         hostname=os.uname().nodename)

@app.route('/health')
def health():
    return {'status': 'ok', 'service': 'flask-app'}

@app.route('/info')
def info():
    return {
        'app': 'Simple Flask App',
        'version': '1.0.0',
        'python': '3.10',
        'server': 'nginx + flask'
    }

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=False)
