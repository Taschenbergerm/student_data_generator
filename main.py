import names 
import fastapi
import uvicorn


app = fastapi.FastAPI()

@app.get("/name/full")
def get_name():
    return names.get_full_name()


@app.get("/name/first")
def get_name():
    return names.get_first_name()

@app.get("/name/last")
def get_name():
    return names.get_last_name()


if __name__ == "__main__":
    uvicorn.run(app)