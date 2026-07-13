FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt .

RUN --mount=type=secret,id=pip.conf,target=/root/.pip/pip.conf pip install -r requirements.txt 

COPY . .

CMD ["python", "app.py"]
