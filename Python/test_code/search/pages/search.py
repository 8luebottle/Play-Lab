from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys


class DuckDuckGoSearchPage:
    URL = 'https://www.duckduckgo.com'

    SEARCH_INPUT = (By.NAME, 'q') # q == search bar named element
    """
    <input id="search_form_input_homepage" class="js-search-input search__input--adv" type="text" 
    autocomplete="off" name="q" tabindex="1" value="" autocapitalize="off" autocorrect="off" 
    placeholder="Search the web without being tracked">
    """

    def __init__(self, browser): # came from the fixture | dependency injection
        self.browser = browser

    def load(self):
        self.browser.get(self.URL) # tell the browser to load a web page of URL

    def search(self, phrase):
        search_input = self.browser.find_element(*self.SEARCH_INPUT) # locator == a query for finding elements
        """call browser find element by name q"""

        search_input.send_keys(phrase, + Keys.RETURN) # hits enter --> submits the search
        