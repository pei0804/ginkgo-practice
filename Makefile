test:
	@ginkgo -r  --failOnPending  --race -skipPackage vendor
