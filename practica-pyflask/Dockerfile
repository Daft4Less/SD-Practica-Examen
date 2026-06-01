FROM python:3.11

RUN useradd -m usuario

WORKDIR /app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

RUN chown -R usuario:usuario /app

USER usuario

ENV PORT=5000 \
    FLASK_DEBUG=0 \
    PYTHONUNBUFFERED=1

EXPOSE 5000

CMD ["python", "run.py"]