import json
import pytest
import requests
import selenium.webdriver


@pytest.fixture
def browser():
    # For Test
    b = selenium.webdriver.Chrome()
    b.implicitly_wait(10)
    yield b
    b.quit()

# @pytest.fixture
# def cbt_config(scope='session'):
#     with open('cbt_config.json') as config_file:
#         config = json.load(config_file)
#
#     assert 'authentication' in config
#     assert 'username' in config['authentication']
#     assert 'key' in config['authentication']
#
#     assert 'webdriver' in config
#     assert 'name' in config['webdriver']
#     assert 'browserName' in config['webdriver']
#     assert 'platform' in config['webdriver']
#
#     return config


# @pytest.fixture
# def browser(cbt_config, request):
#     """For local browser"""
#     username = cbt_config['authentication']['username'].replace('@', '%40')
#     key = cbt_config['authentication']['key']
#     url = f"http://{username}:{key}@hub.crossbrowsertesting.com:80/wd/hub"
#
#     caps = cbt_config['webdriver']
#     caps['name'] += ' | ' + request.node.name
#     b = selenium.webdriver.Remote(desired_capabilities=caps, command_executor=url)
#
#     b.implicitly_wait(30)
#     yield b # act like return statement
#     b.quit() # quit, close, and kill the webdriver instance


# @pytest.hookimpl(tryfirst=True, hookwrapper=True)
# def pytest_runtest_makereport(item, call):
#     outcome = yield
#     setattr(item, 'test_result', outcome.get_result())


# @pytest.fixture
# def cbt_uploader(cbt_config, browser, request):
#     URL = 'https://crossbrowsertesting.com/api/v3'
#
#     username = cbt_config['authentication']['username']
#     key = cbt_config['authentication']['key']
#
#     uploader = requests.Session()
#     uploader.auth = (username, key)
#
#     response = uploader.post(f'{URL}/selenium/{browser.seesion_id}/videos')
#     video_hash = response.json()['hash']
#
#     yield
#
#     uploader.delete(f'{URL}/selenium/{browser.session_id}/videos/{video_hash}')
#
#     score = 'fail' if request.node.test_resutl.failed else 'pass'
#
#     uploader.put(
#         f'{URL}/selenium/{browser.session_id}',
#         data={'action': 'set_score', 'score': score})
