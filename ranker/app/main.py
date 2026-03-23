from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def root():
    return {"ping": "pong"}

@app.get("/health")
async def healthcheck():
    return {"status": "ok"}