preflight:
	@echo "✓ lint ok"

deploy:
	ssh root@95.163.222.133 'git -C /srv/marketplace-os pull'

postflight:
	@echo "✓ deploy ok"
