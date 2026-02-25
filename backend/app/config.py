from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    #db
    database_url: str

    #jwt
    secret_key: str



    class Config:
        env_file = "backend/.env"

settings = Settings()