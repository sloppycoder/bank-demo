FROM python:3-slim-buster as base

FROM base as builder

RUN apt-get -qq update \
    && apt-get install -y --no-install-recommends \
        file \
        g++ \
        libffi-dev 

COPY requirements.txt .
RUN pip install --root="/install" -r requirements.txt

FROM base
COPY --from=builder /install /

COPY entrypoint.sh *.py ids.txt /

EXPOSE 8089
ENTRYPOINT ./entrypoint.sh

