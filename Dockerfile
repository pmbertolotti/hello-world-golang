FROM public.ecr.aws/bitnami/python:latest

WORKDIR /usr/src/app

COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 80

CMD [ "python", "-u", "./test.py" ]
