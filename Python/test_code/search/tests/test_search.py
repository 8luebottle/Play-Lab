"""End-to-End Test"""
# Ref : SmartBear

from ..pages.search import DuckDuckGoSearchPage
from ..pages.result import DuckDuckGoResultPage

def test_basic_duckduckgo_search(browser):
    search_page = DuckDuckGoSearchPage(browser)
    result_page = DuckDuckGoResultPage(browser)

    # Given the DuckDuckGo home page is displayed
    search_page.load()

    # When the user searches for "bluebottle"
    word = "bluebottle"
    search_page.search(word)

    assert word in result_page.title()
    assert word == result_page.search_input_value()
    assert result_page.result_count_for_phrase(word) > 0
