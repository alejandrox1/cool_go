FROM python:3.7-rc-slim-stretch

WORKDIR /app
ADD requirements.txt /app/
RUN pip install --trusted-host pypi.python.org -r requirements.txt
ADD app.py

EXPOSE 80

ENV NAME Alberto

CMD ["python", "app.py"]
