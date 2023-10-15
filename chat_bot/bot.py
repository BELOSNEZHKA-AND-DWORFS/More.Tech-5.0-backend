#! venv/bin/python3

from fastapi import FastAPI
from pydantic import BaseModel
import openai
import os
import json
from dotenv import load_dotenv

dotenv_path = os.path.join(os.path.dirname(__file__), '.env')
if os.path.exists(dotenv_path):
    load_dotenv(dotenv_path)
openai.api_key = os.environ.get("OPENAI_KEY")

if openai.api_key is None:
    raise EnvironmentError("OPENAI_KEY environment variable is not set")

app = FastAPI()

class ChatRequest(BaseModel):
    request: str

@app.post("/chat")
def chat(request_data: ChatRequest):
    print("1")
    messages = [
        {"role": "system", "content": "You are a helpful assistant."}
    ]

    user_request = request_data.request
    with open('prompt.txt', encoding="utf-8") as prompt_file:
        prompt = prompt_file.read()
        messages.append({"role": "user", "content": f"{prompt}\n{user_request}"})
    print("2")
    completion = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=messages
    )
    print("3")
    result = json.loads(str(completion))
    print(result)
    response = result["choices"][0]["message"]["content"]
    print(response)
    return {"response": response}
