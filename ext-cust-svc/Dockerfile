FROM python:3-slim-buster as base

FROM base as builder

COPY requirements.txt .
RUN pip install --root="/install" -r requirements.txt

FROM base
COPY --from=builder /install /

COPY entrypoint.sh server.py ./

EXPOSE 8080
ENTRYPOINT ./entrypoint.sh

