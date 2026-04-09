# GoogleTest
if(WITH_TESTS)
  set(BUILD_GMOCK    ON  CACHE BOOL "" FORCE)
  set(INSTALL_GTEST  OFF CACHE BOOL "" FORCE)
  set(GTEST_HAS_ABSL OFF CACHE BOOL "" FORCE)

  add_subdirectory(googletest)
endif()