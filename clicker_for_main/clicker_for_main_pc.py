import os
import time
import json
import ctypes
import requests
import threading
import pyautogui


ROOT_PATH = os.path.dirname(__file__)
POSTED = False
STARTED = False
VK_CAPITAL = 0x14


def get_capslock_status() -> bool:
    hllDll = ctypes.WinDLL("User32.dll")
    return True if hllDll.GetKeyState(VK_CAPITAL) == 1 else False


def get_request() -> dict:
    return json.loads(requests.get("http://localhost:7575/posts", timeout=2).text)


def post_request(settings, status):

    settings["status"] = status
    print(f"*********** STATUS - {settings['status']} ***********")
    requests.post("http://localhost:7575/posts", json=settings)


settings = get_request()

while True:
    if get_capslock_status():
        posted = False
        while get_capslock_status():
            if not posted:
                post_request(settings, "on")
                posted = True

            def start_clicking():
                while get_capslock_status():
                    for button in settings["buttons"]:
                        pyautogui.keyDown(button)
                        pyautogui.keyUp(button)
                    time.sleep(int(settings["duration"]))

            if not started:
                started = True
                threading.Thread(target=start_clicking).start()

    elif not get_capslock_status():
        posted = False
        while not get_capslock_status():
            if not posted:
                post_request(settings, "off")
                started = False
                posted = True
