macro(package_add_test TESTNAME)
  add_executable(${TESTNAME} ${ARGN})

  target_link_libraries(${TESTNAME} gtest gmock gtest_main)

  gtest_discover_tests(${TESTNAME}
    WORKING_DIRECTORY
      ${ROOT_PATH}
  )
endmacro()

package_add_test(test1 "${TEST_PATH}/casem/test.cpp")