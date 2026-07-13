FROM ubuntu:22.04

RUN apt-get update && apt-get install -y --no-install-recommends \
        python3=3.10* \
        python3-pip=22.0* \
        curl=7.81* \
        wget=1.21* \
        nginx=1.18* \
    && apt-get clean && rm -rf /var/lib/apt/lists/*


WORKDIR /app

COPY requirements.txt .

RUN --mount=type=cache,target=/root/.cache/pip pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 5000 80 443

CMD ["python3", "app.py"]
