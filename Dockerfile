FROM python:3.9-buster

ENV PYTHONUNBUFFERED=1

WORKDIR /code

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

EXPOSE 8000

RUN markata build

CMD ["python","-m","http.server","-d","markout"]
