# To run this, you're required to install fastapi and uvicorn
# pip3 install fastapi uvicorn

# After that simply run python3 ./python-server.py and visit localhost:3333/api?name=Petri

from fastapi import FastAPI, responses
import uvicorn
app = FastAPI()

PORT = 3333

@app.get('/api')
def api(name = "World"):
  return {
    "hello": name,
  }

@app.exception_handler(404)
def not_found(_, __):
  return responses.JSONResponse({"oops": "This is not the page you are looking for"})

if __name__ == "__main__":
  print(f'Listening on localhost:{PORT}')
  uvicorn.run(app, host="0.0.0.0", port=PORT)