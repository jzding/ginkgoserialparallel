# How to run
ginkgo -v -p serial parallel
# What is it doing
The test suite does the following:
- ```ginkgo -v -p serial parallel``` runs 2 test suites serial first, then parallel. -p enables parallel mode but if a suite is serial, this will be ignored for that suite
- first run the serial tests as ordered in the file and not in parallel. This is thanks to the Serial and Ordered  decorators
- After all the serial tests are done, run the parallel tests in parallel. In this case no decorators are used
