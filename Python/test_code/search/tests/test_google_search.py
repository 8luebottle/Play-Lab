from selenium import webdriver
from selenium.webdriver.common.by import By

def test_google_search():
    driver = webdriver.Chrome()
    driver.get('https://google.com')
    word = 'blue bottle'
    driver.find_element(By.NAME, 'q').send_keys(word)
    driver.find_element(By.NAME, 'btnK').submit()
    assert word in driver.title