Feature: Code Execution
  As a developer
  I want to execute code snippets
  So that I can see their output

  Scenario: Valid Python Code
    Given I have a Python code "print('Hello, World!')"
    When I execute the code
    Then the output should be "Hello, World!\n"
    And there should be no errors

  Scenario: Invalid Python Code
    Given I have a Python code "print('Hello, World!')"
    When I execute the code
    Then the output should be "SyntaxError: EOL while scanning string literal\n"
    And there should be an error

  Scenario: Python Code with Arguments
    Given I have a Python code "import sys\nprint(sys.argv[1])"
    And the code arguments are "test-arg"
    When I execute the code
    Then the output should be "test-arg\n"
    And there should be no errors

  Scenario: Long-running Python Code
    Given I have a Python code "import time\ntime.sleep(2)\nprint('Done')"
    When I execute the code
    Then the output should be "Done\n"
    And there should be no errors

  Scenario: Python Code with Syntax Error
    Given I have a Python code "def func:\nprint('Hello')"
    When I execute the code
    Then the output should be "SyntaxError: invalid syntax\n"
    And there should be an error

  Scenario: Python Code with Runtime Error
    Given I have a Python code "print(1 / 0)"
    When I execute the code
    Then the output should be "ZeroDivisionError: division by zero\n"
    And there should be an error

  Scenario: Empty Python Code
    Given I have a Python code ""
    When I execute the code
    Then the output should be ""
    And there should be no errors

  Scenario: Python Code with Unicode Characters
    Given I have a Python code "print('こんにちは世界')"
    When I execute the code
    Then the output should be "こんにちは世界\n"
    And there should be no errors

  Scenario: Python Code with Multiple Arguments
    Given I have a Python code "import sys\nprint(' '.join(sys.argv[1:]))"
    And the code arguments are "arg1 arg2 arg3"
    When I execute the code
    Then the output should be "arg1 arg2 arg3\n"
    And there should be no errors

  Scenario: Python Code with Environment Variable
    Given I have a Python code "import os\nprint(os.getenv('TEST_ENV'))"
    And the environment variable "TEST_ENV" is set to "test_value"
    When I execute the code
    Then the output should be "test_value\n"
    And there should be no errors
