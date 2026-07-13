FROM python:3.11-alpine

RUN apk update && apk upgrade

RUN pip install --no-cache-dir --upgrade pip wheel setuptools

CMD ["python", "-c", "print('Hello')"]
