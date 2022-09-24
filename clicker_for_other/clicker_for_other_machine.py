import os
import time
import json
import configparser
import requests
import threading
import pyautogui

ROOT_PATH = os.path.dirname(__file__)
STARTED = False
config = configparser.ConfigParser()
config.read(
    os.path.join(
        ROOT_PATH,
        "config.ini"
    )
)
address = config["address_server"]["ip"]


def get_request() -> dict:
    return json.loads(requests.get(f"http://{address}/posts", timeout=2).text)


first_print = get_request()
print(f"BUTTONS - {first_print['buttons']}\n"
      f"DURATION - {first_print['duration']}\n"
      f"STATUS - {first_print['status']}")

while True:
    time.sleep(1)
    data = get_request()

    if data["status"] == "on":

        def start_clicking():
            while data["status"] != "off":
                for button in data["buttons"]:
                    pyautogui.keyDown(button)
                    pyautogui.keyUp(button)
                time.sleep(int(data["duration"]))

        if not STARTED:
            STARTED = True
            threading.Thread(target=start_clicking).start()

    elif data["status"] == "off":
        STARTED = False
