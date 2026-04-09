cmake_minimum_required(VERSION 3.31)

project(casem
  LANGUAGES
    CXX
)

# Global variables
set(ROOT_PATH   "${CMAKE_CURRENT_SOURCE_DIR}")
set(SOURCE_PATH "${ROOT_PATH}/src")
set(TEST_PATH   "${ROOT_PATH}/test")

# Add option definitions
include(CMakeOptions.txt)

# Add third-party dependencies provided as sources
add_subdirectory(3rd)

# Casem target
add_executable(casem)

# Provide source files
add_subdirectory(src)

# Provide test files
if(WITH_TESTS)
  enable_testing()
  include(GoogleTest)
  add_subdirectory(test)
endif()