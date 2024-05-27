import os
import sys
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import threading

# Configuration
num_browsers = 200
url = "http://127.0.0.1"
username = "user1"
password = "pass"
chromedriver_path = os.getenv('CHROMEDRIVER_PATH')

# Check for debug mode
debug_mode = '--debug' in sys.argv
if debug_mode:
    num_browsers = 1

def setup_browser():
    # Set up Chrome options
    options = Options()
    if not debug_mode:
        options.headless = True
        options.add_argument("--disable-gpu")
        options.add_argument("--no-sandbox")
        options.add_argument("--disable-dev-shm-usage")
    
    # Initialize the WebDriver instance
    service = Service(chromedriver_path)
    driver = webdriver.Chrome(service=service, options=options)
    return driver

def interact_with_page(driver):
    try:
        driver.get(url)
        
        # Wait until the element with aria-label="openLogin" is present and click it
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.CSS_SELECTOR, '[aria-label="openLogin"]'))
        ).click()
        
        # Wait until the drawer with the text "Login" is present
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.XPATH, '//h2[text()="Login"]'))
        )
        
        # Find the username and password fields and input the data
        username_field = driver.find_element(By.ID, "username")
        password_field = driver.find_element(By.ID, "password")
        
        username_field.send_keys(username)
        password_field.send_keys(password)
        
        # Click the sendLogin button
        driver.find_element(By.ID, "sendLogin").click()
        
    except Exception as e:
        print(f"An error occurred: {e}")
    finally:
        if not debug_mode:
            driver.quit()

def main():
    threads = []

    for _ in range(num_browsers):
        driver = setup_browser()
        if debug_mode:
            interact_with_page(driver)
        else:
            thread = threading.Thread(target=interact_with_page, args=(driver,))
            thread.start()
            threads.append(thread)

    if not debug_mode:
        for thread in threads:
            thread.join()

if __name__ == "__main__":
    main()
