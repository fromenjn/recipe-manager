Feature: List recipes
  In order to see all recipes
  As an API client
  I want to retrieve a list of recipes via GET /recipes

  Scenario: Listing recipes successfully
    Given the server is running
    When I send a GET request to "/recipes"
    Then the response code should be 200
    And the response should not contain "Pancakes"
    And the response should contain "Chocolate"
    And the response should contain "Spaghetti"
